package accelerometer

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"sync"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/support"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	i2c "github.com/davecheney/i2c"
	logging "github.com/op/go-logging"
)

// log is the default package logger
var log = logging.MustGetLogger("trigger-mellis-accel")

const (
	regDevid         = 0x00
	regThreshTap     = 0x1d
	regOfsX          = 0x1e
	regOfsY          = 0x1f
	regOfsZ          = 0x20
	regDur           = 0x21
	regLatent        = 0x22
	regWindow        = 0x23
	regThreshAct     = 0x24
	regThreshInact   = 0x25
	regTimeInact     = 0x26
	regActInact_Ctl  = 0x27
	regThreshFF      = 0x28
	regTimeFF        = 0x29
	regTapAxes       = 0x2a
	regActTap_Status = 0x2b
	regBWRate        = 0x2c
	regPowerCtl      = 0x2d
	regIntEnable     = 0x2e
	regIntMap        = 0x2f
	regIntSource     = 0x30
	regDataFormat    = 0x31
	regDataX0        = 0x32
	regDataX1        = 0x33
	regDataY0        = 0x34
	regDataY1        = 0x35
	regDataZ0        = 0x36
	regDataZ1        = 0x37
	regFifoCtl       = 0x38
	regFifoStatus    = 0x39
)

const (
	powerCtl8Hz       byte = 0x00
	powerCtl4Hz       byte = 0x01
	powerCtl2Hz       byte = 0x02
	powerCtl1Hz       byte = 0x03
	powerCtlSleep     byte = 0x04
	powerCtlMeasure   byte = 0x08
	powerCtlAutoSleep byte = 0x10
	powerCtlLink      byte = 0x20
)

const (
	dataFormatRange2g   byte = 0x00
	dataFormatRange4g   byte = 0x01
	dataFormatRange8g   byte = 0x02
	dataFormatRange16g  byte = 0x03
	dataFormatJustify   byte = 0x04
	dataFormatFullRes   byte = 0x08
	dataFormatIntInvert byte = 0x20
	dataFormatSpi       byte = 0x40
	dataFormatSelfTest  byte = 0x80
)

const (
	bwRate6_25 byte = 0x06
	bwRate12_5 byte = 0x07
	bwRate25   byte = 0x08
	bwRate50   byte = 0x09
	bwRate100  byte = 0x0a
	bwRate200  byte = 0x0b
	bwRate400  byte = 0x0c
	bwRate800  byte = 0x0d
	bwRate1600 byte = 0x0e
	bwRate3200 byte = 0x0f
)

const deviceID byte = 0xE5
const fullResolutionScaleFactor float64 = 3.9

type Adxl345 struct {
	bus     *i2c.I2C
	device  int
	address uint8
}

type Acceleration struct {
	data []float64 /* mg */
}

type Message struct {
	Walking  float64 `json:"walking"`
	Standing float64 `json:"standing"`
	Jogging  float64 `json:"jogging"`
}

func NewAdxl345(address uint8, device int) (*Adxl345, error) {
	adxl := Adxl345{
		device:  device,
		address: address,
	}

	bus, err := i2c.New(address, device)
	if err != nil {
		return nil, err
	}

	adxl.bus = bus
	return &adxl, nil
}

func (adxl *Adxl345) Init() {
	if err := adxl.checkDevID(); err != nil {
		log.Fatalf(err.Error())
	}

	adxl.setRegister(regDataFormat, dataFormatRange16g|dataFormatFullRes)
	adxl.setRegister(regBWRate, bwRate400)
	adxl.setRegister(regPowerCtl, powerCtlMeasure)
}

func (adxl *Adxl345) Destroy() {
}

func (adxl *Adxl345) Read() *Acceleration {
	data := make([]byte, 6, 6)
	var xReg int16
	var yReg int16
	var zReg int16

	adxl.bus.WriteByte(regDataX0)
	adxl.bus.Read(data)

	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.LittleEndian, &xReg)
	binary.Read(buf, binary.LittleEndian, &yReg)
	binary.Read(buf, binary.LittleEndian, &zReg)

	ret := &Acceleration{
		data: make([]float64, 3),
	}
	ret.data[0] = float64(xReg) * fullResolutionScaleFactor
	ret.data[1] = float64(yReg) * fullResolutionScaleFactor
	ret.data[2] = float64(zReg) * fullResolutionScaleFactor

	return ret
}

func (adxl *Adxl345) checkDevID() error {
	data := []byte{0}

	adxl.bus.WriteByte(regDevid)
	adxl.bus.Read(data)

	if data[0] != deviceID {
		errors.New(fmt.Sprintf("ADXL345 at %x on bus %d returned wrong device id: %x\n", adxl.address, adxl.device, data[0]))
	}

	return nil
}

func (adxl *Adxl345) setRegister(register byte, flags byte) {
	data := []byte{register, flags}

	adxl.bus.Write(data)
}

// MyTriggerFactory My Trigger factory
type MyTriggerFactory struct {
	metadata *trigger.Metadata
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &MyTriggerFactory{metadata: md}
}

//New Creates a new trigger instance for a given id
func (t *MyTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	return &MyTrigger{metadata: t.metadata, config: config}
}

// MyTrigger is a stub for your Trigger implementation
type MyTrigger struct {
	metadata *trigger.Metadata
	runner   action.Runner
	config   *trigger.Config
}

// Init implements trigger.Trigger.Init
func (t *MyTrigger) Init(runner action.Runner) {
	t.runner = runner
}

// Metadata implements trigger.Trigger.Metadata
func (t *MyTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

var timeMutex sync.Mutex

// Start implements trigger.Trigger.Start
func (t *MyTrigger) Start() error {
	go t.readData()

	return nil
}

// Stop implements trigger.Trigger.Stop
func (t *MyTrigger) Stop() error {
	// stop the trigger
	return nil
}

func (t *MyTrigger) readData() {
	adxl, _ := NewAdxl345(0x53, 1)
	adxl.Init()

	for {
		data := adxl.Read()

		// Pass the data to the flow
		handlers := t.config.Handlers

		log.Debug("Processing handlers")
		for _, handler := range handlers {
			act := action.Get(handler.ActionId)

			log.Debugf("Found action: '%+x'", act)
			log.Debugf("ActionID: '%s'", handler.ActionId)

			req := t.constructStartRequest(data.data, 100)
			log.Debugf("+v", req)
			startAttrs, _ := t.metadata.OutputsToAttrs(req.Data, false)
			fmt.Println(startAttrs)
			context := trigger.NewContext(context.Background(), startAttrs)
			_, _, err := t.runner.Run(context, act, handler.ActionId, nil)
			if err != nil {
				log.Critical(err.Error())
			}
		}
	}
}

func (t *MyTrigger) constructStartRequest(message []float64, norm int) *StartRequest {
	//TODO how to handle reply to, reply feature
	req := &StartRequest{}
	data := make(map[string]interface{})
	for i := range message {
		message[i] = message[i] / float64(norm)
	}
	data["accelerometer"] = message
	req.Data = data
	return req
}

// StartRequest describes a request for starting a ProcessInstance
type StartRequest struct {
	ProcessURI  string                 `json:"flowUri"`
	Data        map[string]interface{} `json:"data"`
	Interceptor *support.Interceptor   `json:"interceptor"`
	Patch       *support.Patch         `json:"patch"`
	ReplyTo     string                 `json:"replyTo"`
}
