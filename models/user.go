package models

import (
	"encoding/json"
	"fmt"
)

type (

	UserInterface interface {
		login() *error
		register() *error
		isLogin() *error
		//Add more login user data here
		//TODO: Add more user and define it.
	}

	//User implements ModelProvider interface
	//User implements UserInterface
	//User inherit Model
	User struct {
		//Inherit functions from base controller
		BaseModel
		Username string
		Name string
		Email string
		Password string
	}
)

//Test collection type..
const COLLECTION_TYPE  = "TestSnaphyByte"

func init()  {
	//Create an user model and save to database..
	user := &User{
		Name: "Robins",
		Username: "robinskumar73",
		Email: "robinskumar73@gmail.com",
		Password: "12345",
	}

	user.NewModel(COLLECTION_TYPE)
	//Now print the data..
	 u, _ := json.Marshal(user)
	fmt.Println(string(u))

	//Create an user model and save to database..
	ravi := &User{
		Name: "Ravi",
		Username: "ravigupta",
		Email: "ravi@gmail.com",
		Password: "12345",
	}

	ravi.NewModel(COLLECTION_TYPE)
	//Now print the data..
	 r, _ := json.Marshal(ravi)
	fmt.Println(string(r))

}

