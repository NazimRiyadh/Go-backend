package main

import "fmt"

var a = 1
var b = 1

func addNumbers(x int, y int) {
	sum := x + y
	fmt.Println(sum)
}

func main() {
	a := 2 //local scope
	b := 2
	fmt.Println("Outer a:", a, "Outer b:", b) // Use them here
	{
		a := 3 //block scope
		b := 3
		//c =4
		addNumbers(a, b) // prints 6
	}
	//u cannot access c here
}
