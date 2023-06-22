package main

import "fmt"

func main() {
	var x uint8 = 1 << 1 | 1 << 5 // 0000 0010 OR 0010 0000 = 0010 0010
	var y uint8 = 1 << 1 | 1 << 2 // 0000 0010 OR 0000 0100 = 0000 0110

	fmt.Printf("%08b\n", x) // 0010 0010
	fmt.Printf("%08b\n", y) // 0000 0110

	fmt.Printf("%08b\n", x&y) // 0010 0010 AND 0000 0110 = 0000 0010
	fmt.Printf("%08b\n", x|y) // 0010 0010 OR 0000 0110 = 0010 0110
	fmt.Printf("%08b\n", x^y) // 0010 0010 XOR 0000 0110 = 0010 0100
	fmt.Printf("%08b\n", x&^y) // 0010 0010 AND NOT 0000 0110 = 0010 0000

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)		// 1, 5
		}
	}

	fmt.Printf("%08b\n", x<<1) // 0100 0100
	fmt.Printf("%08b\n", x>>1) // 0001 0001
}