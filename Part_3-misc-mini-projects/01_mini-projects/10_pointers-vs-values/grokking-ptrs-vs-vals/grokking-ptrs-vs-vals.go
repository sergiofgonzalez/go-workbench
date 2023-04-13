package main

import (
	"fmt"
	"strings"
)

func primitiveByVal(myInt int, myFloat float64, myBool bool, myString string, myRune rune) {
	fmt.Println("== primitiveByVal")

	myInt *= 2
	myFloat *= 2
	myBool = true
	myString = "Hello to Toby"
	myRune = 'Ñ'
	fmt.Println(myInt)
	fmt.Println(myFloat)
	fmt.Println(myBool)
	fmt.Println(myString)
	fmt.Println(string(myRune))
	fmt.Printf("\n\n")
}

func primitiveByRef(myInt *int, myFloat *float64, myBool *bool, myString *string, myRune *rune) {
	fmt.Println("== primitiveByRef")

	*myInt *= 2
	*myFloat *= 2
	*myBool = true
	*myString = "Hello to Toby"
	*myRune = 'Ñ'
	fmt.Println(*myInt)
	fmt.Println(*myFloat)
	fmt.Println(*myBool)
	fmt.Println(*myString)
	fmt.Println(string(*myRune))
	fmt.Printf("\n\n")
}

func requiresNumberVal(num int) {
	fmt.Println("== requiresNumberVal")
	fmt.Println(num)
	fmt.Printf("\n\n")
}

func arrayByVal(myArray [10]int) {
	fmt.Println("== arrayByVal")
	for i := 0; i < len(myArray); i++ {
		myArray[i] *= 2
	}
	fmt.Println(myArray)
	fmt.Printf("\n\n")
}

func arrayByRef(myArray *[10]int) {
	fmt.Println("== arrayByRef")
	for i := 0; i < len(myArray); i++ {
		myArray[i] *= 2
	}
	fmt.Println(myArray)
	fmt.Printf("\n\n")
}

type someStruct struct {
	name string
	age  int
}

func structsMapsSlicesByVal(myStruct someStruct, myMap map[string]int, mySlice []int) {
	// structs, maps, and slices are data structures that hide
	// some underlying storage.
	// Although I receive those by value, I can change the underlying storage
	// contents
	fmt.Println("== structsMapsSlicesByVal")

	myStruct.name = "Idris"
	myStruct.age = 15

	myMap["France"] = 6

	mySlice[0] = -5

	fmt.Println(myStruct)
	fmt.Println(myMap)
	fmt.Println(mySlice)
	fmt.Printf("\n\n")

}

func structsMapsSlicesByValCannotReassign(myStruct someStruct, myMap map[string]int, mySlice []int) {
	// structs, maps, and slices are data structures that hide
	// some underlying storage.
	// Although I receive those by value, I can change the underlying storage
	// contents
	fmt.Println("== structsMapsSlicesByValCannotReassign")

	myOtherStruct := someStruct{"Idris", 15}
	myStruct = myOtherStruct

	myOtherMap := map[string]int{"France": 6}
	myMap = myOtherMap

	myOtherSlice := []int{-5, -4, -3, -2, -1}
	mySlice = myOtherSlice

	fmt.Println(myStruct)
	fmt.Println(myMap)
	fmt.Println(mySlice)
	fmt.Printf("\n\n")
}

func structsMapsSlicesByRefCannotReassign(myStruct *someStruct, myMap *map[string]int, mySlice *[]int) {
	// structs, maps, and slices are data structures that hide
	// some underlying storage.
	// When you receive them by ref, you cannot reassign either
	fmt.Println("== structsMapsSlicesByRefCannotReassign")

	myOtherStruct := someStruct{"Idris", 15}
	myStruct = &myOtherStruct

	myOtherMap := map[string]int{"France": 6}
	myMap = &myOtherMap

	myOtherSlice := []int{-5, -4, -3, -2, -1}
	mySlice = &myOtherSlice

	fmt.Println(*myStruct)
	fmt.Println(*myMap)
	fmt.Println(*mySlice)
	fmt.Printf("\n\n")
}

func structsMapsSlicesByRef(myStruct *someStruct, myMap *map[string]int, mySlice *[]int) {
	// structs, maps, and slices are data structures that hide
	// some underlying storage.
	// When you receive them by ref, you can reassign
	fmt.Println("== structsMapsSlicesByRef")

	myStruct.name = "Idris"
	myStruct.age = 15

	(*myMap)["France"] = 6

	(*mySlice)[0] = -5

	fmt.Println(*myStruct)
	fmt.Println(*myMap)
	fmt.Println(*mySlice)
	fmt.Printf("\n\n")
}

func (s someStruct) valueReceiver() {
	fmt.Println("== valueReceiver")

	// Can't change or reassign
	s.name = "Margot"
	s.age = 35
	fmt.Println(s)

	// and cannot reassign either
	myOtherStruct := someStruct{"Margot", 35}
	s = myOtherStruct

	fmt.Printf("\n\n")
}

func (s *someStruct) ptrReceiver() {
	fmt.Println("== ptrReceiver")
	fmt.Println(s)

	// Can change
	s.name = "Emily"
	s.age = 45
	fmt.Println(s)

	// but cannot reassign
	myOtherStruct := someStruct{"Margot", 35}
	s = &myOtherStruct

	fmt.Printf("\n\n")
}

func doubleMe(num *int) {
	*num *= 2
}

func doubleElems(arr *[10]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
}

type person struct {
	name string
	age  int
}

func doublePersonAge(p *person) {
	p.age *= 2
}

func capitalizeKeys(m map[string]int) {
	for k, v := range m {
		m[strings.ToUpper(k)] = v
	}
}

func doubleSliceElems(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] *= 2
	}
}

func (p *person) changePerson(other person) {
	*p = other
// 	p.name = other.name
// 	p.age = other.age
}

func main() {
	myInt := 5
	myFloat := 5.4321
	myBool := false
	myString := "Hello to Jason Isaacs!"
	myRune := 'Ó'

	// Passing by value means i'll get my values back when the function returns
	primitiveByVal(myInt, myFloat, myBool, myString, myRune)
	fmt.Println(myInt)
	fmt.Println(myFloat)
	fmt.Println(myBool)
	fmt.Println(myString)
	fmt.Println(string(myRune))
	fmt.Printf("\n\n")

	// Passing by ref means the values might change
	primitiveByRef(&myInt, &myFloat, &myBool, &myString, &myRune)
	fmt.Println(myInt)
	fmt.Println(myFloat)
	fmt.Println(myBool)
	fmt.Println(myString)
	fmt.Println(string(myRune))
	fmt.Printf("\n\n")

	// if a function requires a value, I can call it with a pointer
	requiresNumberVal(myInt)
	myIntPtr := &myInt
	requiresNumberVal(*myIntPtr)
	fmt.Printf("\n\n")

	// Arrays are passed by value
	myArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	arrayByVal(myArray)
	fmt.Println(myArray)
	fmt.Printf("\n\n")

	// If you want to change arrays, you need to pass its address
	arrayByRef(&myArray)
	fmt.Println(myArray)
	fmt.Printf("\n\n")

	// You can pass structs, maps, and slices by value
	// and change its contents
	myStruct := someStruct{"Jason", 55}
	myMap := map[string]int{"Spain": 5}
	mySlice := []int{0, 1, 2, 3, 4}
	structsMapsSlicesByVal(myStruct, myMap, mySlice)
	fmt.Println(myStruct)
	fmt.Println(myMap)
	fmt.Println(mySlice)
	fmt.Printf("\n\n")

	// And you cannot reassign when passing by val
	myStruct = someStruct{"Jason", 55}
	myMap = map[string]int{"Spain": 5}
	mySlice = []int{0, 1, 2, 3, 4}
	structsMapsSlicesByValCannotReassign(myStruct, myMap, mySlice)
	fmt.Println(myStruct)
	fmt.Println(myMap)
	fmt.Println(mySlice)
	fmt.Printf("\n\n")

	// And you cannot reassign either when passing by ref
	structsMapsSlicesByRef(&myStruct, &myMap, &mySlice)
	fmt.Println(myStruct)
	fmt.Println(myMap)
	fmt.Println(mySlice)
	fmt.Printf("\n\n")

	// Method receivers
	myStructPtr := &myStruct

	myStruct.valueReceiver()
	myStructPtr.valueReceiver()

	myStruct.ptrReceiver()
	myStructPtr.ptrReceiver()

	// for the README.md
	num := 5
	doubleMe(&num)
	fmt.Println(num)

	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	doubleElems(&arr)
	fmt.Println(arr)

	p := person{"Sergio", 49}
	doublePersonAge(&p)
	fmt.Println(p.age)

	m := map[string]int{"Spain": 0, "France": 1}
	capitalizeKeys(m)
	fmt.Println(m)

	s := []int{0, 1, 2, 3, 4}
	doubleSliceElems(s)
	fmt.Println(s)

	p1 := person{
		name: "Jason Isaacs",
		age: 57,
	}
	p1.changePerson(person{"Idris Elba", 49})
	fmt.Println(p1)
}

