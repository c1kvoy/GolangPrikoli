package usermethods

import (
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

// Определение метода updateUser
func (user *Account) UpdateUser(newUser Account) {
	if newUser.Username != "" {
		user.Username = newUser.Username
	}
	if newUser.Email != "" {
		user.Email = newUser.Email
	}
}
