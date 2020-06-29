package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-bank-backend/helpers"
	"github.com/go-bank-backend/users"

	"github.com/gorilla/mux"
)

//Login struct
type Login struct {
	Username string
	Password string
}

//ErrResponse struct
type ErrResponse struct {
	Message string
}

//Login api function
func login(w http.ResponseWriter, r *http.Request) {
	//Read the body from api request
	body, err := ioutil.ReadAll(r.Body)
	//Check for the errors
	helpers.HandleErr(err)

	//Assign variable to Login struct
	var formattedBody Login
	//unmarshal the body while body and formattedBody as arguments
	err = json.Unmarshal(body, &formattedBody)
	//Check for the errors
	helpers.HandleErr(err)
	//Assign username and password of user
	login := users.Login(formattedBody.Username, formattedBody.Password)

	//Check if the message is equal to "All Okay" else show the following errors
	if login["message"] == "All Okay" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		resp := ErrResponse{Message: "Wrong username and password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func StartAPI() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	fmt.Println("App is working on port: 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
