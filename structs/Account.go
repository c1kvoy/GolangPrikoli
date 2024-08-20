package main

import (
	"first/structs/usermethods"
	"fmt"
)

func main() {
	user1 := usermethods.NewAccount("Vsevolod", "klonsve@mail.ru")
	fmt.Println("Username before: ", user1)

	newUser := usermethods.Account{
		Username: "Roflus!",
	}
	user1.UpdateUser(newUser)
	fmt.Println("Username after: ", user1)
}
