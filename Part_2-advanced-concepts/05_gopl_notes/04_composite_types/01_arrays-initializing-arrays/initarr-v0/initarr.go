package main

import "fmt"

func main() {
	a := [...]string{1: "$", 2: "€", 0: "£", 3: "¥"}
	fmt.Println(a)
}