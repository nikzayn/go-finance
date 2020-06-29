package migrations

import (
	"github.com/go-bank-backend/helpers"
	"github.com/go-bank-backend/interfaces"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Create Dummy Account
func createAccounts() {
	db := helpers.ConnectDB()
	users := &[2]interfaces.User{
		{Username: "nikhil", Email: "nik@gmail.com"},
		{Username: "justin", Email: "justin@gmail.com"},
	}

	for i := 0; i < len(users); i++ {
		// Generated Password
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		// User Details
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		// Create user in DB
		db.Create(&user)
		// Account Details
		account := &interfaces.Account{Type: "Daily Accounts",
			Name:    string(users[i].Username + "'s" + " account"),
			Balance: uint(10000 * int(i+1)),
			UserID:  user.ID}
		// Create account in DB
		db.Create(&account)
	}
	defer db.Close()
}

// DB Migrate
func Migrate() {
	db := helpers.ConnectDB()
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db.AutoMigrate(User, Account)
	defer db.Close()
	createAccounts()
}
