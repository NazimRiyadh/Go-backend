package main

import "fmt"

//struct embeding means using a struct inside other struct
type social struct {
	fb        string
	linked_in string
}

type user struct {
	name   string
	age    int
	city   string
	social //struct embeding
}

func main() {
	user1 := user{
		name: "riyadh",
		age:  18,
		city: "Dhaka",
		social: social{
			fb:        "NazimRiyadh",
			linked_in: "Nazim Riyadh",
		},
	}

	fmt.Println(user1)
}
