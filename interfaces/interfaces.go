package interfaces

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

type ResponseAccount struct {
	ID uint
	Name string
	Balance uint
}

type ResponseUser struct{
	ID uint
	Username string
	Email string
	Accounts []ResponseAccount
}