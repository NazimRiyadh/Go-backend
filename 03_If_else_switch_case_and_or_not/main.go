package main

import (
	"fmt"
	"time"
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
	b := 10
	switch b {
	case 1:
		fmt.Println("You choose one")
	case 2:
		fmt.Println("You choose two")
	case 3:
		fmt.Println("You choose three")
	default:
		fmt.Println("I am here by default!")
	}

	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend")
	default:
		fmt.Println("It's Workday :{")

	}
	a := 2
	switch a {
	case 'c', 'd':
		fmt.Println("I am actually a letter")
	case 1, 2:
		fmt.Println("I am a number")
	default:
		fmt.Println("I donno what I am!")

	}

	//type Switching
	who_am_I := func(i interface{}) {

		switch i.(type) {
		case int:
			fmt.Println("I am a int")
		case string:
			fmt.Println("I am a string")
		case bool:
			fmt.Println("I am a bool")
		default:
			fmt.Println("I donno the type")
		}
	}
	who_am_I("Hola!")
}

// and=&&, or= ||, not= !
