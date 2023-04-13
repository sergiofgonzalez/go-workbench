package main

import "fmt"

// Person is a data structure representing a person identified by a name and age
type Person struct {
	name string
	age  int
}

// IncAge increments the age of a person by mutating its age
func IncAge(p *Person) {
	p.age++
	fmt.Printf("%q has turned %d\n", p.name, p.age)
}


func main() {
	actors := []Person{
		{name: "Jason Isaacs", age: 60},
		{name: "Margot Robbie", age: 33},
		{name: "Idris Elba", age: 51},
		{name: "Florence Pugh",age: 30 },
	}

	fmt.Println("Before:", actors)

	// using for i := 0
	// for i := 0; i < len(actors); i++ {
	// 	fmt.Printf("About to call IncAge with %v (%p)", actors[i], &actors[i])
	// 	IncAge(&actors[i])
	// }

	// using range
	for _, actor := range actors {
		fmt.Printf("About to call IncAge with %v (%p)", actor, &actor)
		IncAge(&actor)
	}

	fmt.Println("After :", actors)

	// One possible fix is to map the initial slice to a slice of pointers
	actorPtrs := []*Person{
		&Person{name: "Jason Isaacs", age: 60},
		&Person{name: "Margot Robbie", age: 33},
		&Person{name: "Idris Elba", age: 51},
		&Person{name: "Florence Pugh",age: 30 },
	}

	for _, actorPtr := range actorPtrs {
		fmt.Printf("About to call IncAge with %v (%p)", actorPtr, &actorPtr)
		IncAge(actorPtr)
	}

	fmt.Println("After :", actorPtrs)
	for _, actor := range actorPtrs {
		fmt.Println(*actor)
	}
}

