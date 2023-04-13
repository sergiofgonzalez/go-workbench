package main

import "fmt"

// ByteSlice is an alias to a []byte so that we can define methods on it
type ByteSlice []byte


// Write copies the data at the end of the byteslice and returns the number of
// bytes appended and an error indicating whether an error ocurred.
func (s *ByteSlice) Write(data []byte) (int, error) {
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
	return len(data), nil
}

func main() {
	bs := ByteSlice{0, 1, 2}
	bs.Write([]byte{3, 4, 5})

	fmt.Println(bs)

	s := ByteSlice{}
	name := "Jason"
	last := "Isaacs"
	age := 57
	fmt.Fprintf(&s, "Hello to %v %v who is %v years old\n", name, last, age)
	fmt.Println(string(s))
}