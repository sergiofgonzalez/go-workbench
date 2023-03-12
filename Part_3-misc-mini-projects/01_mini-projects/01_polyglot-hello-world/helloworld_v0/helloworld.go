package main

import "fmt"

func main() {
	greeting := greet("es")
	fmt.Println(greeting)
}

type language string

func greet(l language) string {
	switch l {
	case "en":
		return "Hello, world!"
	case "fr":
		return "Bonjour, le monde!"
	case "es":
		return "Hola, mundo!"
	case "de":
		return "Hallo, Welt!"
	default:
		return ""
	}
}
