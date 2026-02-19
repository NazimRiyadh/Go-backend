package main

import "fmt"

type order struct {
	id     string
	amount float32
	status string
}

func init_order() *order {
	return &order{}

}

func (o *order) create_order(id string, amount float32, status string) {
	o.id = id
	o.amount = amount
	o.status = status
}

func main() {
	order1 := init_order()

	order1.create_order("1", 100, "recieved")

	fmt.Println(*order1)

}
