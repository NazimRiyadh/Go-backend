package main

import "fmt"

func main() {
	a := []int{4, 7, 10, 13}

	for idx, num := range a {
		fmt.Println("index: ", idx, "number: ", num)
	}

	name := "riyadh"


	//unicode point rune
	//here "a" kind of reflects the byte 
	//here "b" reflects the unicode of the letter
	for a, b := range name {
		fmt.Println(a, b)
	}
}
