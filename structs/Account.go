package main

import (
	_ "encoding/json"
	"first/structs/usermethods"
	"fmt"

	_ "github.com/davecgh/go-spew/spew"
)

func main() {
	user1 := usermethods.NewAccount("Vsevolod", "klonsve@mail.ru")
	user2 := usermethods.NewAccount("Roflus", "klonsve@mail.ru")
	fmt.Println("Username before: ", user1)
	user1.UpdateUser(user2)
	fmt.Println("Username after: ", user2)
	// user1.UpdateUser(newUser)
	// json_data, _ := json.MarshalIndent(user1, "", "")
	// fmt.Print(string(json_data), "\n")
	// spew.Dump(user1)
}
