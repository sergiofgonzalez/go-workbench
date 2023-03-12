package main

import "fmt"

func main() {
	greeting := greet("es")
	fmt.Println(greeting)
}

type language string

var greetings = map[language]string{
	"en": "Hello, world!",
	"fr": "Bonjour, le monde!",
	"es": "Hola, mundo!",
	"de": "Hallo, Welt!",
}

func greet(l language) string {
	v, ok := greetings[l]
	if ok {
		return v
	}
	return fmt.Sprintf("unsupported language: %q", l)
}
