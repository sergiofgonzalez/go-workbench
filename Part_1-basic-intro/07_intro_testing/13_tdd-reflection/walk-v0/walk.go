package main

func walk(x interface{}, fn func(string)) {
	fn("Hello!")
}