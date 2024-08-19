package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id          uuid.UUID
	Username    string
	Email       string
	CreatetDate time.Time
}

func NewAccount(username, email string) Account {
	user := Account{
		Id:          uuid.New(),
		Username:    username,
		Email:       email,
		CreatetDate: time.Now(),
	}
	return user
}

func (user *Account) updateUser(newUser Account) {
	if newUser.Username != "" {
		user.Username = newUser.Username
	}
	if newUser.Email != "" {
		user.Email = newUser.Email
	}
}

func main() {
	user1 := NewAccount("Vsevolod", "klonsve@mail.ru")
	fmt.Println("Username before: ", user1)
	newUser := Account{
		Username: "Roflus!",
	}
	user1.updateUser(newUser)
	fmt.Println("Username after: ", user1)
}
