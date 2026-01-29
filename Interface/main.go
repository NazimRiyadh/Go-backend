package main

import "fmt"

type people interface {
	printDetails()
}

type government interface {
	corruption(money int)
}

type user struct {
	name   string
	age    int
	class  int
	saving int
}

func (obj user) printDetails() {
	fmt.Println("I am the function from 'people' interface")
}

func (obj user) corruption(a int) {
	fmt.Println("I am corruption as i am from 'government' interface")
}

func main() {

	var first people = user{
		name:  "riyadh",
		age:   18,
		class: 11,
	}

	fmt.Println(first)
	first.printDetails()
	neww, ok := first.(government)
	if ok == false {
		fmt.Println("Interface conversion failed")
	}
	neww.corruption(10)

}
