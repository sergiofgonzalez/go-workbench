package main

import "fmt"

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = byte('A')
	}
	return len(b), nil
}

func main() {
	r := MyReader{}
	b := make([]byte, 5)

	for i := 0; i < 3; i++ {
		n, err := r.Read(b)
		fmt.Printf("\tn=%v, err=%v, b=%v\n", n, err, b)
		fmt.Printf("b[:n]=%q\n", b[:n])
	}
}