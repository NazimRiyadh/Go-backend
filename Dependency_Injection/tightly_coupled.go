package main

import (
	"fmt"
)

type Bakery struct {
	oventype    string
	ingridients []string
}

func (b *Bakery) Bake() {
	switch b.oventype {
	case "gas oven":
		fmt.Println("Started Gas Oven")
	case "electric oven":
		fmt.Println("Satrted Electric Oven")
	}

	for _, ingridient := range b.ingridients {
		switch ingridient {
		case "flour":
			fmt.Println("Mixing Flour")
		case "sugar":
			fmt.Println("Mixing Sugar")
		case "strawberry":
			fmt.Println("Mixing Starwberry")
		case "vanilla":
			fmt.Println("Mixing Vanilla")
		}
	}
	fmt.Println("Pastry is ready!!")
}

func main() {
	Strawberry_Pastry := Bakery{
		oventype:    "gas oven",
		ingridients: []string{"flour", "sugar", "strawberry"},
	}
	Strawberry_Pastry.Bake()

	Vanilla_Pastry := Bakery{
		oventype:    "electric oven",
		ingridients: []string{"flour", "sugar", "vanilla"},
	}
	Vanilla_Pastry.Bake()
}
