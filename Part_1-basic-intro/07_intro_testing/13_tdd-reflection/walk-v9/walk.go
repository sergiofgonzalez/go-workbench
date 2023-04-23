package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	numValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		numValues = val.Len()
		getField = val.Index
	case reflect.Struct:
		numValues = val.NumField()
		getField = val.Field
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}

	case reflect.String:
		fn(val.String())
	}

	for i := 0; i < numValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// if val is a pointer dereference it
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}