package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotReader rot13Reader) Read(b []byte) (int, error) {
	n, err := rotReader.r.Read(b)
	for i := 0; i < n; i++ {
		if (b[i] >= 'A' && b[i] <= 'Z') {
			b[i] = ((b[i] - 'A' + 13) % ('Z' - 'A' + 1)) + 'A'
		} else if (b[i] >= 'a' && b[i] <= 'z') {
			b[i] = ((b[i] - 'a' + 13) % ('z' - 'a' + 1)) + 'a'
		}
		// fmt.Printf("%q\n", b[i])
	}
	return n, err
}

func main() {
	r := rot13Reader{strings.NewReader("Lbh penpxrq gur pbqr!")}

	// custom reader
	b := make([]byte, 4)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Printf("b[:n]=%q\n", b[:n])
	}

	// Fine trick to output the stream to stdout
	r = rot13Reader{strings.NewReader("Abjurer!\n")}
	io.Copy(os.Stdout, &r)

	r = rot13Reader{strings.NewReader("Terra!\n")}
	io.Copy(os.Stdout, &r)

	r = rot13Reader{strings.NewReader("Ares!\n")}
	io.Copy(os.Stdout, &r)
}