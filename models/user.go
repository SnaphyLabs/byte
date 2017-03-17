package models

type (
	//User implements ModelProvider interface
	//User inherit Model
	User struct {
		Model
		Id string
		Created string
		Updated string
		Type string
		Password string
		Username string
		Name string
	}
)


//Get the User model
func(u *User) getUser() (*User, error){
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

