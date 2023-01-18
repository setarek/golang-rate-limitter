package utils

import "strconv"

func ParseInt64(value interface{}) int64 {
	switch v := value.(type) {
	case string:
		val, _ := strconv.Atoi(v)
		if val >= 0 {
			return int64(val)
		}
		return 0
	case float64:
		if v >= 0 {
			return int64(v)
		}
		return 0
	case uint:
		return int64(v)
	case int:
		if v >= 0 {
			return int64(v)
		}
		return 0
	case int64:
		return v
	default:
		return 0
	}
}
