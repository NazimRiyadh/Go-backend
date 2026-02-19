package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	user1 := user{
		name: "habib",
		age:  30,
	}
	fmt.Println("Name: ", user1.name)
	fmt.Println("Age: ", user1.age)

	user2 := user{
		name: "riyadh",
		age:  24,
	}
	fmt.Println("Name: ", user2.name)
	fmt.Println("Age: ", user2.age)
}
