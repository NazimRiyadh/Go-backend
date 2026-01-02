package main

import (
	"fmt"
)

func main() {
	age := 20
	if age >= 18 {
		fmt.Print("adult")
	} else {
		fmt.Print("Children")
	}
	fmt.Println()
	//switch case
	b := 2
	switch b {
	case 1:
		fmt.Println("You choose one")
	case 2:
		fmt.Println("You choose two")
	case 3:
		fmt.Println("You choose three")
	}
}

// and=&&, or= ||, not= !
