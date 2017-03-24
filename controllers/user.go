package controllers

import "github.com/SnaphyLabs/SnaphyByte/models"


type (
	//Inherit controller interface..
	//Define the user controller
	UserController struct {
		//Each controller will inherit a Connection struct..
		Controller
	}
)


func init(){
	//uc := new(UserController)
	//ms.NewController(uc)
}


func (uc *UserController)getUserById(u *models.User) (error) {
	//TODO: Fetch user from database..
	 _, err := uc.Find(nil, 0, 0)
	if err != nil{
		return err
	}


	return nil
}



