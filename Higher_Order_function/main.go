package main

import "fmt"

var add = func(a int, b int) {
	fmt.Println("a+b: ", a+b)
}

//In Golang functions are First class varibale

//"sum" is Higher Order Function(HOF), First Class Function as it recieves a function
//"add" is the callback function which was passed to HOF

func sum(a int, b int, yoyo func(x int, y int)) {
	yoyo(a, b)
}

func main() {
	add(5, 10)
	(sum(1, 2, add))
}
