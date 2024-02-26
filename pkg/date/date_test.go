package date

import (
	"reflect"
	"testing"
)

func TestFormatDate(t *testing.T) {
	var t1 int64 = 1708909339
	t.Log(reflect.TypeOf(t1).String())
	var t2 = FormatDate(t1)
	t.Log(t2)
}
