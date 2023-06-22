package main

import "fmt"

func main() {
	s := "プログラム"
	fmt.Printf("% x\n", s)  // spc is intentional (another formatting trick!)

	// Transform the utf-8 string in an array of runes
	r := []rune(s)
	fmt.Printf("%x\n", r)

	// Transform an array/slice of runes into an utf-8 encoded string
	s = string(r)
	fmt.Println(s)

	// Converting an integer value to a string
	fmt.Println(string(65)) // "A", not "65"

	// A rune is already an uint32, so no conversion needed
	c := 'A'
	fmt.Println(c)  // 65

	// A string with one char do needs conversion, but it is implicit
	s = "A"
	fmt.Println(s[0])  // 65

	// When an invalid rune is used the replacement character is used
	fmt.Println(string(1234567))  // �
	fmt.Println(string(0x4eac))		// 京

}

