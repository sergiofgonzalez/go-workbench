package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	contents, err := Contents("../README.md")
	if err != nil {
		fmt.Printf("Couldn't get the file contents: %v\n", err)
	}
	fmt.Println(contents)
}


// Contents reads the given filename and returns its contents as a string
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var contents []byte
	buf := make([]byte, 100)	// We'll read in chunks of 100 bytes
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break	// no more iterations needed
			}
			return "", err
		}
		contents = append(contents, buf[0:n]...)
	}
	return string(contents), nil
}