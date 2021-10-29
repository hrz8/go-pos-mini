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
