package main

import (
	"encoding/json"
	"first/structs/usermethods"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	user1 := usermethods.NewAccount("Vsevolod", "klonsve@mail.ru")
	fmt.Println("Username before: ", user1)

	newUser := usermethods.Account{
		Username: "Roflus!",
	}
	user1.UpdateUser(newUser)
	json_data, _ := json.MarshalIndent(user1, "", "")
	fmt.Print(string(json_data), "\n")
	spew.Dump(user1)
}
