package daterange

import (
	"time"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-parsecsv")

const (
	ivStartDate = "startDate"
	ivEndDate   = "endDate"

	ovStartDate = "newStartDate"
	ovEndDate   = "newEndDate"
)

// MyActivity is a stub for your Activity implementation
type DateRangeActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &DateRangeActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *DateRangeActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func calcOffset(date *time.Time, offset int, units string) {
	if offset == 0 {
		return
	}

	switch units {
	case "second", "seconds":
		*date = date.Add(time.Second * time.Duration(offset))
	case "minute", "minutes":
		*date = date.Add(time.Minute * time.Duration(offset))
	case "days":
		*date = date.AddDate(0, 0, offset)
	case "months":
		*date = date.AddDate(0, offset, 0)
	case "years":
		*date = date.AddDate(offset, 0, 0)
	}

}

// Eval implements activity.Activity.Eval
func (a *DateRangeActivity) Eval(ctx activity.Context) (done bool, err error) {
	inputStart := ctx.GetInput(ivStartDate).(map[string]interface{})
	inputEnd := ctx.GetInput(ivEndDate).(map[string]interface{})

	start := time.Now()
	end := time.Now()

	if str, ok := inputStart["Date"].(string); ok && str != "" {
		start, _ = time.Parse("01022006150405", inputStart["Date"].(string))
	}
	if str, ok := inputEnd["Date"].(string); ok && str != "" {
		end, _ = time.Parse("01022006150405", inputEnd["Date"].(string))
	}

	if offset, ok := inputStart["Offset"].(int); ok {
		if units, ok := inputStart["Units"].(string); ok {
			calcOffset(&start, offset, units)
		} else {
			activityLog.Error("Offset specified, but no unit specified.")
		}
	}
	if offset, ok := inputEnd["Offset"].(int); ok {
		if units, ok := inputEnd["Units"].(string); ok {
			calcOffset(&end, offset, units)
		} else {
			activityLog.Error("Offset specified, but no unit specified.")
		}
	}

	ctx.SetOutput(ovStartDate, start.Format("01022006150405"))
	ctx.SetOutput(ovEndDate, end.Format("01022006150405"))

	return true, nil
}
