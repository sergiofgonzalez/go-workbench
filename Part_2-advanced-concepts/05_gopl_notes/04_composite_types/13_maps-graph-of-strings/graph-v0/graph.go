package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool) // set up value of the map
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	addEdge("mariaje", "sergio")
	addEdge("mariaje", "javi")
	addEdge("sergio", "adri")

	fmt.Println(graph)
}
