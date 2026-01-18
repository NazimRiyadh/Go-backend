package main

import "fmt"

var (
	a = 20
	b = 40
)

func add(x int, y int) {
	sum := x + y
	fmt.Println(sum)
}

//used variable shadowing as multiple variables with same name exist in different scopes
func main() {
	p := 2
	q := 4

	//a, b are in Global scope
	//p,q are in local scope
	add(p, q)
	add(a, b)
	add(a, p)

}

//6
//60
//22
