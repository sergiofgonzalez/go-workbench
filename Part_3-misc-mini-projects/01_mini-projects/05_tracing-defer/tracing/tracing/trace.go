package tracing

import "fmt"

func Trace(label string) string {
	fmt.Printf("In %q\n", label)
	return label
}

func Un(label string) {
	fmt.Printf("Exiting %q\n", label)
}