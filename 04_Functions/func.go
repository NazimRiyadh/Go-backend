package main

import "fmt"

//function structue//--func(arg argtype)return_type{}

func add(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println(sum)
}

//multiple return values
func dual(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul

}

func main() {
	a, b := 10, 12
	add(a, b)
	p, q := dual(a, b)
	fmt.Println("p:", p, "Q:", q)

}
