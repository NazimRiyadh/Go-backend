package main

import "fmt"

func calc() int {
	result := 0
	fmt.Println("First calc: ", result)

	show := func() {
		result = result + 10
		fmt.Println("defer calc: ", result)
	}

	defer show()

	result = 5
	fmt.Println("Last calc: ", result)

	return result
}

func calcilator() (result int) {
	result = 0
	fmt.Println("First calculator: ", result)

	show := func() {
		result = result + 10
		fmt.Println("defer calculator: ", result)
	}

	defer show()

	result = 5
	fmt.Println("Last calculator: ", result)
	return result
}

func main() {

}
