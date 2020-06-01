package models

import (
	"github.com/AdrianMendez1199/simple-crud/db"
)

type User struct {
	// ID       string `json:"id"`
	Name     string `json:"name" `
	LastName string `json:"lastname" `
	Email    string `json:"email" `
	Password string `json:"password" `
}

func (u *User) CreateUser(name, lastname, email, password string) *User {
	var DB = db.ConnectDB()

	u.Name = name
	u.LastName = lastname
	u.Email = email
	u.Password = password

	err := DB.Insert(u)
	if err != nil {
		panic(err)
	}

	return u
}
