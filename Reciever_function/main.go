package main

import "fmt"

type user struct {
	name  string
	age   int
	class int
	email string
}

//value reciever function
func (u user) valueReciever(name string, age int, class int, email string) {
	u.name = name
	u.age = age
	u.class = class
	u.email = email

	fmt.Println("I have changed the first user using Value Reciever, it should be:")
	fmt.Println(u.name, u.age, u.class, u.email)
}

//pointer reciever function
func (u *user) pointerReciever(name string, age int, class int, email string) {
	u.name = name
	u.age = age
	u.class = class
	u.email = email

	fmt.Println("I have changed the first user using Pointer Reciever, it should be:")
	fmt.Println(u.name, u.age, u.class, u.email)
}

//Reciever function mainly used to connect a function with struct

func main() {
	first := user{
		name:  "Saleka",
		age:   20,
		class: 10,
		email: "Saleka@gmail.com",
	}

	second := user{
		name:  "Aleka",
		age:   10,
		class: 5,
		email: "Aleka@gmail.com",
	}

	first.valueReciever("Riyadh", 25, 12, "riyadh@gmail.com")
	fmt.Println("But the values in first users are:")
	fmt.Println(first)

	second.pointerReciever("Riyadh", 25, 12, "riyadh@gmail.com")
	fmt.Println("But the values in first users are:")
	fmt.Println(second)
}
