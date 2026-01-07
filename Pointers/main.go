//what is the main use of poniter?
//they points at some memory address
//we change value directly using the poniter
//pointer is game changer, as they help to pass by reference in functions

package main

import "fmt"

//pass by reference, passing the direct memory address
func change(b *int) {
	*b = 20
}

//pass by valude
func nochange(a int) {
	a = 30
}

func main() {
	a := 10
	b := &a

	fmt.Println("Value of a: ", a)
	fmt.Println("Value of a at address: ", b)
	fmt.Println("Value at addrees: ", b, " is ", *b)

	change(b)
	fmt.Println("after change function, function changed me :(, I passed the reference, a:", a)

	nochange(a)
	fmt.Println("after nochange function, function can't change me :), I passed the value, a:", a)

}
