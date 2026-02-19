package main

import "fmt"

//Cannot firectly use Enums in GO
//Mainly used to group a set a value to avoid any input mistake

type Orderstatus string

const (
	Recieved  Orderstatus = "recieved"
	Pending               = "pending"
	Inventory             = "Inventory"
	Stock_out             = "stock_out"
)

func ChangeStatus(status Orderstatus) {
	fmt.Println("Oderstatus Changed to ", status)
}

func main() {
	ChangeStatus(Inventory)
}
