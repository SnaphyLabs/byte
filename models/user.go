package models

type (
	//User implements ModelProvider interface
	//User inherit Model
	User struct {
		//Inherit functions from base controller
		BaseModel
		Password string
		Username string
		FirstName string
		LastName string
		Name string
	}
)

func init()  {
	//Create an user model and save to database..
	//user := &User{}
}

/*

//Get the User model
func(u *User) get() (*User, error){
	return u, nil
}

//Save user to database
func (u *User) save() (error){
	//TODO: Save data to database..

	return nil
}


//Delete data from database..
func (u *User) destroy() error  {
	//TODO: Delete user from database..

	return nil
}
*/

