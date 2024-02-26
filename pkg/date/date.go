package date

import (
	"reflect"
	"time"
)

const (
	TimeLayout = "2006-01-02 15:04:05"
	MaxTimeStr = "9999-12-31 23:59:59"
)

func NowTime() *time.Time {
	var now = time.Now()
	return &now
}

func FormatDate(date interface{}) string {
	if date == nil {
		return "-"
	}

	switch reflect.TypeOf(date).String() {
	case "time.Time":
		dateT := date.(time.Time)
		if dateT.IsZero() {
			return "-"
		}
		return dateT.Format(TimeLayout)
	case "*time.Time":
		if reflect.ValueOf(date).IsNil() {
			return "-"
		}
		return date.(*time.Time).Format(TimeLayout)
	case "int64":
		t, ok := date.(int64)
		if !ok || t <= 0 {
			return "-"
		}
		return time.Unix(t, 0).Format(TimeLayout)
	default:
		return "-"
	}
}
