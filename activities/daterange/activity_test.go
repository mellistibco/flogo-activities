package daterange

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	startDate := map[string]interface{}{
		"Date":   "04042018215500",
		"Offset": -5,
		"Units":  "minutes",
	}
	endDate := map[string]interface{}{
		"Date":   "",
		"Offset": 0,
		"Units":  "",
	}
	tc.SetInput("format", "01022006150405")
	tc.SetInput("startDate", startDate)
	tc.SetInput("endDate", endDate)

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	//check result attr
	newStart := tc.GetOutput("newStartDate")
	newEnd := tc.GetOutput("newEndDate")
	fmt.Println(newStart)
	fmt.Println(newEnd)
}

func TestMissingFields(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	startDate := map[string]interface{}{
		"Date":   "04052018180000",
		"Offset": -5,
		"Units":  "minutes",
	}
	endDate := map[string]interface{}{
		"Date":   "",
		"Offset": 0,
		"Units":  "",
	}
	tc.SetInput("format", "01022006150405")
	tc.SetInput("startDate", startDate)
	tc.SetInput("endDate", endDate)

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	//check result attr
	newStart := tc.GetOutput("newStartDate")
	newEnd := tc.GetOutput("newEndDate")
	fmt.Println(newStart)
	fmt.Println(newEnd)
}

func TestNoEndDate(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	startDate := map[string]interface{}{
		"Date":   "04052018180000",
		"Offset": -5,
		"Units":  "minutes",
	}
	tc.SetInput("format", "01022006150405")
	tc.SetInput("startDate", startDate)

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	//check result attr
	newStart := tc.GetOutput("newStartDate")
	newEnd := tc.GetOutput("newEndDate")
	fmt.Println(newStart)
	fmt.Println(newEnd)
}
