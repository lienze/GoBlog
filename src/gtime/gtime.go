package gtime

import "time"

const (
	BASIC      = 1
	BASIC_MILL = 2
	BASIC_NANO = 3
	BASIC_FULL = 4
)

func GetCurTime(iType int8) string {
	switch iType {
	case BASIC:
		return time.Now().Format("2006-1-2 15:04:05")
	case BASIC_MILL:
		return time.Now().Format("2006-1-2 15:04:05.000")
	case BASIC_NANO:
		return time.Now().Format("2006-1-2 15:04:05.0000000")
	case BASIC_FULL:
		return time.Now().Local().String()
	default:
		time.Now().Format("2006-1-2 15:04:05")
	}
	return time.Now().Format("2006-1-2 15:04:05")
}
