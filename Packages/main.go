package main

import (
	"fmt"

	user "github.com/NazimRiyadh/Go/User"
	login "github.com/NazimRiyadh/Go/auth"
)

func main() {
	login.LoginCredential("riyadh", "8754321")
	user := user.User{
		Email: "riyadh@gmail.com",
		Age:   12,
	}

	fmt.Println(user)
}
