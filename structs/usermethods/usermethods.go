package usermethods

import (
	"time"

	"github.com/google/uuid"
)

type account struct {
	Id          uuid.UUID
	Username    string
	Email       string
	CreatetDate time.Time
}

func NewAccount(username, email string) account {
	user := account{
		Id:          uuid.New(),
		Username:    username,
		Email:       email,
		CreatetDate: time.Now(),
	}
	return user
}

// Определение метода updateUser
func (user *account) UpdateUser(newUser account) {
	if newUser.Username != "" {
		user.Username = newUser.Username
	}
	if newUser.Email != "" {
		user.Email = newUser.Email
	}
}
