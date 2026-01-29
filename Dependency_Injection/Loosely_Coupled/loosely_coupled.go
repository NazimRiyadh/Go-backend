package main

import "fmt"

type Oven interface {
	Heat()
}

type ingridients interface {
	Mix()
}

//gas oven
type gasoven struct {
}

func (obj gasoven) Heat() {
	fmt.Println("Started Gas oven")
}

//electric oven
type electricoven struct {
}

func (obj electricoven) Heat() {
	fmt.Println("Started electtric Oven")
}

//Ingridients----------

//flour
type flour struct {
}

func (obj flour) Mix() {
	fmt.Println("Mixing Flour")
}

//sugar
type sugar struct {
}

func (obj sugar) Mix() {
	fmt.Println("Mixing Sugar")
}

//strawberry
type strawberry struct {
}

func (obj strawberry) Mix() {
	fmt.Println("Mixing Strawberry")
}

type vanilla struct {
}

func (obj vanilla) Mix() {
	fmt.Println("Mixing Vanilla")
}

type Bakery struct {
	oven        Oven
	ingridients []ingridients
}

func (b *Bakery) Bake() {
	b.oven.Heat()
	for _, ingridient := range b.ingridients {
		ingridient.Mix()
	}
	fmt.Println("Pastry is ready!!")
}

func main() {
	gasoven := gasoven{}
	electricoven := electricoven{}
	flour := flour{}
	sugar := sugar{}
	strawberry := strawberry{}
	vanilla := vanilla{}

	strawberryPastry := Bakery{
		oven:        gasoven,
		ingridients: []ingridients{flour, sugar, strawberry},
	}
	strawberryPastry.Bake()

	vanillaPastry := Bakery{
		oven:        electricoven,
		ingridients: []ingridients{flour, sugar, vanilla},
	}
	vanillaPastry.Bake()
}
