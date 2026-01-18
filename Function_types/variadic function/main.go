package main

import "fmt"

//variadic functions don't have fixed parameter numbers, they can be dynamic based i=on argument
func add(num ...int) int {
	sum := 0

	for _, num := range num {
		sum += num
	}
	return sum
}

//can take any type of value as parameter bcz of interface
func variety(element ...interface{}) {
	for _, value := range element {
		fmt.Println(value)
	}
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1, 2, 3))

	fmt.Println("---------------")

	variety(1, 2, "Riyadh")
}
