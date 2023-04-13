package main

import "fmt"

// ByteSlice is an alias to a []byte so that we can define methods on it
type ByteSlice []byte


// Append copies the data at the end of the byteslice and returns it as a new []byte
func (s *ByteSlice) Append(data []byte) {
	bs := *s	// bs has dereferenced the method receiver

	l := len(bs) // Initial length

	if l + len(data) > cap(bs) {
		aux := make([]byte, 2 * (l + len(data)))
		copy(aux, bs)
		bs = aux
	}
	bs = bs[0:l + len(data)]
	copy(bs[l:], data)

	// We cannot reassign the method receiver, but we can change its contents
	// by copying the contents of the resulting ByteSlice
	// The address of s will be the same, but the memory it points to
	// will contain something different
	*s = bs
}

func main() {
	bs := ByteSlice{0, 1, 2}
	bs.Append([]byte{3, 4, 5})

	fmt.Println(bs)
}