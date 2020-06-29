package interfaces

import "github.com/jinzhu/gorm"

// User Schema
type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

// Account Schema
type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

// ResponseAccount Schema
type ResponseAccount struct {
	ID      uint
	Name    string
	Balance uint
}

// ResponseUser Schema
type ResponseUser struct {
	ID       uint
	Username string
	Email    string
	Accounts []ResponseAccount
}
