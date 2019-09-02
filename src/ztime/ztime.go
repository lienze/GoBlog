package ztime

import "time"

// DAT means Date And Time
// T mean Time
const (
	DAT = iota + 1 // start from 1
	DAT_MILL
	DAT_NANO
	DAT_FULL
	D2AT
	D2AT_MILL
	D2AT_NANO
	T
	T_MILL
	T_NANO
)
const (
	STYLE1 = iota + 1 // start from 1
	STYLE2
	STYLE3
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
	case D2AT:
		return time.Now().Format("2006-01-02 15:04:05")
	case D2AT_MILL:
		return time.Now().Format("2006-01-02 15:04:05.000")
	case D2AT_NANO:
		return time.Now().Format("2006-01-02 15:04:05.0000000")
	case T:
		return time.Now().Format("15:04:05")
	case T_MILL:
		return time.Now().Format("15:04:05.000")
	case T_NANO:
		return time.Now().Format("15:04:05.0000000")
	default:
		time.Now().Format("2006-01-02 15:04:05")
	}
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetCurDate(iType int8) string {
	switch iType {
	case STYLE1:
		return time.Now().Format("20060102")
	case STYLE2:
		return time.Now().Format("2006-1-2")
	case STYLE3:
		return time.Now().Format("2006-01-02")
	default:
		return time.Now().Format("20060102")
	}
}
