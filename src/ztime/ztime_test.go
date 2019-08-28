package ztime

import "testing"

func TestGetCurDate(t *testing.T) {
	curDate := GetCurDate(STYLE1)
	if len(curDate) != 8 {
		t.FailNow()
	}
}
