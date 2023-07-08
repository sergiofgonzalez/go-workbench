package main

import (
	"fmt"
	"strings"

	"example.com/arraylist/list"
)

func main() {
	fmt.Println("Hello!")
	l := list.New[string]()
	l.Insert("Hello", 0)
	l.Insert("to", 1)
	l.Insert("Jason", 2)
	l.Insert("Isaacs", 3)

	fmt.Println(l)

	// Using traverse to print each element in its own line
	fmt.Println("=== Traverse1")
	l.Traverse(printLnElem)

	// Using traverse to print each element along with its index
	fmt.Println("=== Traverse2")
	l.Traverse(printLnIndexElem())

	// Using traverse to get a copy of the list on which every element has been made uppercase
	fmt.Println("=== Traverse3")
	fmt.Println(copyUppercase(l))

	// Using traverse to mutate the list making each element uppercase
	fmt.Println("=== Traverse4")
	mutateUppercase(l)
	fmt.Println(l)

}

func printLnElem(elem string) {
	fmt.Println(elem)
}

func printLnIndexElem() func(elem string) {
	index := 0
	return func(elem string) {
		fmt.Printf("%d: %v\n", index, elem)
		index++
	}
}

func copyUppercase(src *list.List[string]) *list.List[string] {
	c := list.New[string]()
	i := 0
	src.Traverse(func(e string) {
		err := c.Insert(strings.ToUpper(string(e)), i)
		if err != nil {
			panic("Could not complete copyUppercase")
		}
		i++
	})
	return c
}

func mutateUppercase(l *list.List[string]) {
	i := 0
	l.Traverse(func(e string) {
		err := l.Remove(i)
		if err != nil {
			fmt.Printf("mutateUppercase: %v\n", err)
			panic("Could not remove in mutateUppercase")
		}

		err = l.Insert(strings.ToUpper(string(e)), i)
		if err != nil {
			panic("Could not insert in mutateUppercase")
		}
		i++
	})
}
