package graphql

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/TIBCOSoftware/flogo-contrib/trigger/rest/cors"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/graphql-go/graphql"
	"github.com/julienschmidt/httprouter"

	"net/http"
)

const (
	REST_CORS_PREFIX = "REST_TRIGGER"
)

// log is the default package logger
var log = logger.GetLogger("trigger-flogo-graphql")

var validMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
var gqlObjects map[string]*graphql.Object
var graphQlSchema *graphql.Schema

// GraphQLTrigger REST trigger struct
type GraphQLTrigger struct {
	metadata *trigger.Metadata
	server   *Server
	config   *trigger.Config
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &RestFactory{metadata: md}
}

// RestFactory REST Trigger factory
type RestFactory struct {
	metadata *trigger.Metadata
}

//New Creates a new trigger instance for a given id
func (t *RestFactory) New(config *trigger.Config) trigger.Trigger {
	return &GraphQLTrigger{metadata: t.metadata, config: config}
}

// Metadata implements trigger.Trigger.Metadata
func (t *GraphQLTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

func (t *GraphQLTrigger) Initialize(ctx trigger.InitContext) error {
	router := httprouter.New()

	if t.config.Settings == nil {
		return fmt.Errorf("no Settings found for trigger '%s'", t.config.Id)
	}

	if _, ok := t.config.Settings["port"]; !ok {
		return fmt.Errorf("no Port found for trigger '%s' in settings", t.config.Id)
	}

	addr := ":" + t.config.GetSetting("port")

	pathMap := make(map[string]string)

	// Build the GraphQL Object Types & Schemas
	t.buildGraphQLObjects()
	t.buildGraphQLSchema(ctx.GetHandlers())

	// Init handlers
	for _, handler := range ctx.GetHandlers() {

		err := validateHandler(t.config.Id, handler)
		if err != nil {
			return err
		}

		method := strings.ToUpper(handler.GetStringSetting("method"))
		path := handler.GetStringSetting("path")

		log.Debugf("Registering handler [%s: %s]", method, path)

		if _, ok := pathMap[path]; !ok {
			pathMap[path] = path
			router.OPTIONS(path, handleCorsPreflight) // for CORS
		}

		router.Handle(method, path, newActionHandler(t, handler))
	}

	log.Debugf("Configured on port %s", t.config.Settings["port"])
	t.server = NewServer(addr, router)

	return nil
}

func (t *GraphQLTrigger) buildGraphQLObjects() {
	gqlTypes := t.config.Settings["types"].([]map[string]interface{})

	// Create type objects
	gqlObjects = make(map[string]*graphql.Object)

	// Get the graphql types
	for _, typ := range gqlTypes {
		name := typ["Name"].(string)
		fields := make(graphql.Fields)

		for k, f := range typ["Fields"].(map[string]interface{}) {
			fTyp := f.(map[string]interface{})

			fields[k] = &graphql.Field{
				Type: coerceType(fTyp["Type"].(string)),
			}
		}

		obj := graphql.NewObject(
			graphql.ObjectConfig{
				Name:   name,
				Fields: fields,
			})

		gqlObjects[name] = obj
	}
}

func (t *GraphQLTrigger) buildGraphQLSchema(handlers []*trigger.Handler) *graphql.Schema {
	fSchema := t.config.Settings["schema"].(map[string]interface{})

	// Build the graphql schema
	var schema graphql.Schema
	var queryType *graphql.Object

	if strings.EqualFold(t.config.Settings["operation"].(string), "query") {

		var objName string
		queryFields := make(graphql.Fields)

		// Get the object name
		for k, v := range fSchema["Query"].(map[string]interface{}) {
			if strings.EqualFold(k, "name") {
				objName = v.(string)
			} else if strings.EqualFold(k, "fields") {
				qf := v.(map[string]interface{})

				for k, v := range qf {

					// Grab query args
					argObj := v.(map[string]interface{})
					args := make(graphql.FieldConfigArgument)

					for k, v := range argObj["Args"].(map[string]interface{}) {

						argTyp := v.(map[string]interface{})
						args[k] = &graphql.ArgumentConfig{
							Type: coerceType(argTyp["Type"].(string)), // get typ from v.(string)
						}
					}

					for _, handler := range handlers {
						if strings.EqualFold(handler.GetStringSetting("type"), k) {
							// Build the queryField
							queryFields[k] = &graphql.Field{
								Type:    gqlObjects[k], // need to fix this so its case insensitive
								Args:    args,
								Resolve: fieldResolver(handler),
							}
						}
					}
				}
			}
		}

		queryType = graphql.NewObject(
			graphql.ObjectConfig{
				Name:   objName,
				Fields: queryFields,
			})
	}

	schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		})

	return &schema
}

func (t *GraphQLTrigger) Start() error {
	return t.server.Start()
}

// Stop implements util.Managed.Stop
func (t *GraphQLTrigger) Stop() error {
	return t.server.Stop()
}

func fieldResolver(handler *trigger.Handler) graphql.FieldResolveFn {

	return func(p graphql.ResolveParams) (interface{}, error) {

		triggerData := map[string]interface{}{
			"resolveParams": p,
		}

		results, err := handler.Handle(context.Background(), triggerData)

		return results, err
	}

}

// Handles the cors preflight request
func handleCorsPreflight(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	log.Infof("Received [OPTIONS] request to CorsPreFlight: %+v", r)

	c := cors.New(REST_CORS_PREFIX, log)
	c.HandlePreflight(w, r)
}

func newActionHandler(rt *GraphQLTrigger, handler *trigger.Handler) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		log.Infof("Received request for id '%s'", rt.config.Id)

		c := cors.New(REST_CORS_PREFIX, log)
		c.WriteCorsActualRequestHeaders(w)

		pathParams := make(map[string]string)
		for _, param := range ps {
			pathParams[param.Key] = param.Value
		}

		queryValues := r.URL.Query()
		queryParams := make(map[string]string, len(queryValues))
		header := make(map[string]string, len(r.Header))

		for key, value := range r.Header {
			header[key] = strings.Join(value, ",")
		}

		for key, value := range queryValues {
			queryParams[key] = strings.Join(value, ",")
		}

		triggerData := map[string]interface{}{
			"params":      pathParams,
			"pathParams":  pathParams,
			"queryParams": queryParams,
			"header":      header,
		}

		// Check the HTTP Header Content-Type
		contentType := r.Header.Get("Content-Type")
		switch contentType {
		case "application/x-www-form-urlencoded":
			buf := new(bytes.Buffer)
			buf.ReadFrom(r.Body)
			s := buf.String()
			m, err := url.ParseQuery(s)
			content := make(map[string]interface{}, 0)
			if err != nil {
				log.Errorf("Error while parsing query string: %s", err.Error())
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			for key, val := range m {
				if len(val) == 1 {
					content[key] = val[0]
				} else {
					content[key] = val[0]
				}
			}
			triggerData["content"] = content
		default:
			var content interface{}
			err := json.NewDecoder(r.Body).Decode(&content)
			if err != nil {
				switch {
				case err == io.EOF:
					// empty body
					//todo should handler say if content is expected?
				case err != nil:
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}
			triggerData["content"] = content
		}

		// Fix this
		result := graphql.Do(graphql.Params{
			Schema:        *graphQlSchema,
			RequestString: queryParams["query"],
		})

		var replyData interface{}
		var replyCode int

		resp := result.Data.(map[string]interface{})
		if len(resp) != 0 {
			dataAttr, ok := resp["data"]
			if ok {
				replyData = dataAttr
			}
			codeAttr, ok := resp["code"]
			if ok {
				replyCode, _ = data.CoerceToInteger(codeAttr)
			}
		}

		if len(result.Errors) > 0 {
			log.Debugf("REST Trigger Error: %s", result.Errors)
			http.Error(w, result.Errors[0].Error(), http.StatusBadRequest)
			return
		}

		if replyData != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			if replyCode == 0 {
				replyCode = 200
			}
			w.WriteHeader(replyCode)
			if err := json.NewEncoder(w).Encode(replyData); err != nil {
				log.Error(err)
			}
			return
		}

		if replyCode > 0 {
			w.WriteHeader(replyCode)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////
// Utils

func coerceType(typ string) *graphql.Scalar {
	switch typ {
	case "graphql.String":
		return graphql.String
	}

	return nil
}

func validateHandler(triggerId string, handler *trigger.Handler) error {

	method := handler.GetStringSetting("method")
	if method == "" {
		return fmt.Errorf("no method specified in handler for trigger '%s'", triggerId)
	}

	if !stringInList(strings.ToUpper(method), validMethods) {
		return fmt.Errorf("invalid method '%s' specified in handler for trigger '%s'", method, triggerId)
	}

	//validate path

	return nil
}

func stringInList(str string, list []string) bool {
	for _, value := range list {
		if value == str {
			return true
		}
	}
	return false
}
