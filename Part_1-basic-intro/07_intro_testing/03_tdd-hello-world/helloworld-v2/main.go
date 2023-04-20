package main

import "fmt"

// Hello returns a greeting for the name given or "Hello, World!" if given an empty string.
func Hello(name string) string {
	if name == "" {
		return "Hello, World!"
	}
	return fmt.Sprintf("Hello to %s!", name)
}

func main() {
	fmt.Println(Hello("Jason Isaacs"))
}
