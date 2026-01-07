package main

import "fmt"

func main() {
	fmt.Println("Welocme to Array!")

	//if you don't declare array with values , by default the value is 0s
	a := [5]int{}
	fmt.Println(a)

	b := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(" value of array b: ", b)
	fmt.Println(" value at index no 3: ", b[3])

}
