package gtime

import "time"

// DAT means Date And Time
// T mean Time
const (
	DAT      = 1
	DAT_MILL = 2
	DAT_NANO = 3
	DAT_FULL = 4
	T        = 5
	T_MILL   = 6
	T_NANO   = 7
)
const (
	STYLE1 = 1
)

func GetCurTime(iType int8) string {
	switch iType {
	case DAT:
		return time.Now().Format("2006-1-2 15:04:05")
	case DAT_MILL:
		return time.Now().Format("2006-1-2 15:04:05.000")
	case DAT_NANO:
		return time.Now().Format("2006-1-2 15:04:05.0000000")
	case DAT_FULL:
		return time.Now().Local().String()
	case T:
		return time.Now().Format("15:04:05")
	case T_MILL:
		return time.Now().Format("15:04:05.000")
	case T_NANO:
		return time.Now().Format("15:04:05.0000000")
	default:
		time.Now().Format("2006-1-2 15:04:05")
	}
	return time.Now().Format("2006-1-2 15:04:05")
}

func GetCurDate(iType int8) string {
	switch iType {
	case STYLE1:
		return time.Now().Format("20060102")
	default:
		return time.Now().Format("20060102")
	}
}
