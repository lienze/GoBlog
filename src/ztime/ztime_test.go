package ztime

import "testing"

func TestGetCurTime(t *testing.T) {
	curTime := GetCurTime(DAT)
	if len(curTime) > 25 {
		t.FailNow()
	}
}
func TestGetCurDate(t *testing.T) {
	curDate := GetCurDate(STYLE1)
	if len(curDate) != 8 {
		t.FailNow()
	}
}
