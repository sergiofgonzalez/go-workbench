package main

import "fmt"

type MyType struct {
	field1 string
	field2 int
}

func (X *MyType) method1() {
	fmt.Println("in method1")
}

type MyOtherType MyType


func main() {
	myVar := MyType{"field1", 1}
	myVar.method1()

	myOtherVar := MyOtherType{"Fld1", 2}	// Fields are there
	// myOtherVar.method1()	// but methods are not
	fmt.Println(myOtherVar)
}