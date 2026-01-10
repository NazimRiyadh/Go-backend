package main

import "fmt"

func a() {
	a := 0
	fmt.Println("First: ", a)

	defer fmt.Println("Defer Second: ", a)

	a = a + 1
	fmt.Println("Third: ", a)

}

func main() {
	a()
}
