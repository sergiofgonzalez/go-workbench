package main

import "fmt"

// ByteSlice is an alias to a []byte so that we can define methods on it
type ByteSlice []byte


// Append copies the data at the end of the byteslice and returns it as a new []byte
func (bs ByteSlice) Append(data []byte) []byte {
	l := len(bs) // Initial length

	if l + len(data) > cap(bs) {
		aux := make([]byte, 2 * (l + len(data)))
		copy(aux, bs)
		bs = aux
	}
	bs = bs[0:l + len(data)]
	copy(bs[l:], data)
	return bs
}

func main() {
	bs := ByteSlice{0, 1, 2}
	newBs := bs.Append([]byte{3, 4, 5})

	fmt.Println(bs)
	fmt.Println(newBs)
}