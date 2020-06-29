package helpers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// Error handler
func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// Create Hash and Salt
func HashAndSalt(pass []byte) string {
	//Bcrypt will generate password
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)

	// Check for errors
	HandleErr(err)

	// Return the hashed string
	return string(hashed)
}

// Connect Database
func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=nikhil123 sslmode=disable")
	HandleErr(err)
	return db
}
