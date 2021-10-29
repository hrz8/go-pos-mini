package helpers

import "reflect"

// NilToEmptyMap is a helper function to convert nil value to {}
func NilToEmptyMap(d *interface{}) interface{} {
	data := *d
	if *d == nil {
		data = make(map[string]interface{})
	}
	return data
}

// Ternary is a helper function to get non-nil value
func Ternary(d interface{}, s interface{}) interface{} {
	data := d
	if d == nil || (reflect.ValueOf(d).Kind() == reflect.Ptr && reflect.ValueOf(d).IsNil()) {
		data = s
	}
	return data
}

// GetOffset is a helper function to get sql offset value from page and limit args
func GetOffset(page int, limit int) int {
	offset := (page - 1) * limit
	if offset < 0 {
		return 0
	}
	return offset
}
