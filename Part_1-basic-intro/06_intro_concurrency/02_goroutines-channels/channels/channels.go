package main

import "fmt"

func sum(s []int, c chan int) {
	var total int
	for _, v := range s {
		total += v
	}
	c <- total
}


func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c := make(chan int)
	go sum(s[0:len(s) / 2], c)
	go sum(s[len(s) / 2:], c)
	s1, s2 := <- c, <- c
	fmt.Println(s1 + s2)
}