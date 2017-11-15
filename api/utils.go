package api

import (
	"strconv"
)

func parseIntParam(value int, min int, max int) string {
	if value < min {
		return intToString(min)
	}
	if value > max {
		return intToString(max)
	}
	return intToString(value)
}

func parseInt64Param(value int64, min int64, max int64) string {
	if value < min {
		return int64ToString(min)
	}
	if value > max {
		return int64ToString(max)
	}
	return int64ToString(value)
}

func parseStringParam(value string, allows []string, def string) string {
	if allows != nil {
		for _, s := range allows {
			if value == s {
				return value
			}
		}
	}
	return def
}

func parseBoolParam(value bool) string {
	if value {
		return "true"
	}
	return "false"
}

func intToString(val int) string {
	return strconv.FormatUint(uint64(val), 10)
}

func int64ToString(val int64) string {
	return strconv.FormatUint(uint64(val), 10)
}
