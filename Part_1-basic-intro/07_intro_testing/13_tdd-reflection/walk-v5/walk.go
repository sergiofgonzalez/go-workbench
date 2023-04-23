package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	for i := 0; i < val.NumField(); i++ {
		fld := val.Field(i)

		switch fld.Kind() {
		case reflect.String:
			fn(fld.String())
		case reflect.Struct:
			walk(fld.Interface(), fn)
		}
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