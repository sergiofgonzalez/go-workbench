package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"sergio": 49,
		"inma": 53,
		"adri": 15,
		"maria": 10,
		"alexx": 6,
	}

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
 }