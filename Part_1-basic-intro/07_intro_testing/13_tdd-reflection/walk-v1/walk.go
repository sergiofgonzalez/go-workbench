package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	fld := val.Field(0)
	fn(fld.String())
}