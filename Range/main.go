package main

import "fmt"

func main() {
	a := []int{4, 7, 10, 13}

	for idx, num := range a {
		fmt.Println("index: ", idx, "number: ", num)
	}

	name := "riyadh"

	for a, b := range name {
		fmt.Println(a, b)
	}
}
