package main

import "fmt"

//you can't write ':=' outside function, it is a statement
//Outside main declaration is allowed not any statement

//mainly here func is a anonymous but we stored in a variable 'add'
var add = func(a, b int) int {
	return a + b
}

var sub = func(a, b int) int {
	return a - b
}(5, 3)

func print() {
	fmt.Println("Printing Outter sub func value: ", sub)
}

var a = 10

func main() {

	fmt.Println(add(2, 3))
	print()
}
