package env

import (
	"os"
	"strconv"
	"time"
)

func GetOr(key string, value ...string) string {
	if key == "" {
		return ``
	}
	var v = os.Getenv(key)
	if v == "" && len(value) > 0 {
		return value[0]
	}
	return v
}

func GetBoolOr(key string, value ...bool) bool {
	var v = GetOr(key)
	if v == "" {
		if len(value) > 0 {
			return value[0]
		}
		return false
	}
	if b, err := strconv.ParseBool(v); err == nil {
		return b
	}
	switch v {
	case `Yes`, `On`, `1`, `True`, `TRUE`, `OK`:
		return true
	case `NO`, `Off`, `0`, `False`, `FALSE`:
		return false
	default:
		return false
	}
}

func GetIntOr(key string, value ...int) int {
	var v = GetOr(key)
	if v == "" {
		if len(value) > 0 {
			return value[0]
		}
		return 0
	}
	if b, err := strconv.Atoi(v); err == nil {
		return b
	}
	if len(value) > 0 {
		return value[0]
	}
	return 0
}

func GetFloatOr(key string, value ...float64) float64 {
		var v = GetOr(key)
		if v == "" {
				if len(value) > 0 {
						return value[0]
				}
				return 0
		}
		if b, err := strconv.ParseFloat(v,10); err == nil {
				return b
		}
		if len(value) > 0 {
				return value[0]
		}
		return 0
}

func GetDuration(key string, value ...time.Duration) time.Duration {
	var v = GetOr(key)
	if v == "" {
		if len(value) > 0 {
			return value[0]
		}
		return 0
	}
	if d, err := time.ParseDuration(v); err == nil {
		return d
	}
	if len(value) > 0 {
		return value[0]
	}
	return 0
}
