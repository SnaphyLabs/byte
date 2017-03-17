package controllers

import (
	"github.com/robinskumar73/byte/models"
	"github.com/SnaphyLabs/SnaphyByte/database"
)

type (
	//Inherit controller interface..
	//Define the user controller
	UserController struct {
		//Each controller will inherit a Connection struct..
		Controller
	}
)


func init(){
	uc := new(UserController)
	ms := database.MongoStorage{}
	ms.NewController(uc)
}


func (uc *UserController)getUserById(u *models.User) (error) {
	//TODO: Fetch user from database..

	return nil
}



