package ztime

import "testing"

func TestGetCurDate(t *testing.T) {
	curDate := GetCurDate(STYLE1)
	if len(curDate) != 8 {
		t.FailNow()
	}
	/*switch iType {
	case STYLE1:
		return time.Now().Format("20060102")
	default:
		return time.Now().Format("20060102")
	}*/
}
