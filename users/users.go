package users

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-bank-backend/helpers"
	"github.com/go-bank-backend/interfaces"
	"golang.org/x/crypto/bcrypt"
)

// Login the user
func Login(username string, pass string) map[string]interface{} {
	// Connect DB
	db := helpers.ConnectDB()
	// Get user from User Schema
	user := &interfaces.User{}
	//Check if the user has the status of "RecordNotFound"
	if db.Where("username = ? ", username).First(&user).RecordNotFound() {
		//Return the error message
		return map[string]interface{}{"message": "User Not Found"}
	}

	//Compare Password with the existing password
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	//Check if our password is not mismatched and if there is no error
	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "Wrong Password"}
	}

	//To store all the data of ResponseAccount in an array
	accounts := []interfaces.ResponseAccount{}
	//To check in a database that table named "accounts" where user_id is equal to the user.ID of another table "users"
	db.Table("accounts").Select("id", "name", "balance").Where("user_id = ? ", user.ID).Scan(&accounts)

	//Response User Struct
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
	}

	defer db.Close()

	//Create jwt token for the user
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}

	//Create claim with signing method and jwt token
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	//Signing the token password
	token, err := jwtToken.SignedString([]byte("Token Passworf"))
	//Check for the errors
	helpers.HandleErr(err)

	//Response message
	var response = map[string]interface{}{"message": "All Okay"}
	//Set jwt token as token
	response["jwt"] = token
	//Set response data as responseUser
	response["data"] = responseUser

	return response
}
