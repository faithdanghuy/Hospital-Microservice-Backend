package pointer

import "time"

func String(value string) *string {
	return &value
}

func Int(value int) *int {
	return &value
}

func Float(value float64) *float64 {
	return &value
}

func Bool(value bool) *bool {
	return &value
}
func Time(t *time.Time) time.Time {
	if t != nil {
		return *t
	}
	return time.Time{}
}

func SetString(source *string, value string) {
	if source == nil {
		source = new(string)
	}
	*source = value
}

func SetBool(source *bool, value bool) {
	if source == nil {
		source = new(bool)
	}
	*source = value
}

func SetInt(source *int, value int) {
	if source == nil {
		source = new(int)
	}
	*source = value
}

func SetFloat64(source *float64, value float64) {
	if source == nil {
		source = new(float64)
	}
	*source = value
}
