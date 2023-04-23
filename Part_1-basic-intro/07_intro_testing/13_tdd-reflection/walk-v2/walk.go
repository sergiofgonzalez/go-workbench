package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		fld := val.Field(i)
		fn(fld.String())
	}
}