package main

import "fmt"

type Vertex struct {
	Lat float64
	Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{40.68433, -74.39967}
	m["Torre del Oro"] = Vertex{37.382412, -5.996490}
	m["Rockefeller Center"] = Vertex{40.758678,	-73.978798}

	fmt.Println(m["Rockefeller Center"])
	fmt.Println(m["Torre del Oro"])
	fmt.Println(m["Bell Labs"])
}