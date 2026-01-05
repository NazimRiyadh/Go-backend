// init function runs first, before main also

package main

import "fmt"

var a = 10

func main() {
	fmt.Println("I am main, printing the value of a:", a)
}

//init running before main
func init() {
	fmt.Println("I am init, changing the value a")
	a = 20
}
