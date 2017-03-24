package models


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


