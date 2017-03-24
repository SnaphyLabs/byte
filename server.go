package main

import (
	/*"github.com/SnaphyLabs/SnaphyByte/database"
	"time"
	"github.com/rs/rest-layer/resource"
	"encoding/json"*/
	"fmt"
	_ "github.com/SnaphyLabs/SnaphyByte/models"
	/*"errors"
	"reflect"*/
)

const (
	MongoDBHosts = "localhost:27017"
	Database = "drugcorner"
	UserName = "robins"
	Password = "12345"
	Collection =  "Demo"
)


func init(){

	/*//Connect to mongodb database..
	ds := &database.DataStorage{
		Address:[]string{MongoDBHosts},
		Database:Database,
		Username:UserName,
		Password:Password,
		Collection:Collection,
		Timeout:60 * time.Second,
	}


	//Now define all the related methods
	ds.Find = func(lookup interface{}, offset, limit int) (interface{}, error)  {
		//Write the logic to connect to the server..
		//Return the complete userList
		return UserList, nil
	}

	ds.Clear = func(lookup *resource.Lookup) (int, error) {
		return 0, nil
	}

	ds.Connect = func() (interface{}, error) {
		//Write the logic to connect to the server..
		i := new(interface{})
		return i, nil
	}

	//Useful links
	//https://github.com/swhite24/go-rest-tutorial
	//Get type name..
	//https://play.golang.org/

	//Using reflection..
	//http://stackoverflow.com/questions/35657362/how-to-return-dynamic-type-struct-in-golang
	//http://stackoverflow.com/questions/14116840/dynamically-call-method-on-interface-regardless-of-receiver-type
	ds.Delete = func(item interface{}) error {

		u := models.User{}
		UserType := reflect.TypeOf(u)

		if user, ok := item.(UserType); ok{
			for _, value := range UserList{
				if value.ID == user.ID{
					fmt.Println("Item Deleted")
					break
				}
			}
		}

		return errors.New("Data Invalid")
	}

	ds.Insert = func(item interface{}) error {
		//Create an user model and save to database..
		user := &models.User{
			Name: "Robins",
			Username: "robinskumar73",
			Email: "robinskumar73@gmail.com",
			Password: "12345",
		}

		user.NewModel(Collection)
		//Now print the data..
		u, _ := json.Marshal(user)
		fmt.Println(string(u))

		//Create an user model and save to database..
		ravi := &models.User{
			Name: "Ravi",
			Username: "ravigupta",
			Email: "ravi@gmail.com",
			Password: "12345",
		}

		ravi.NewModel(Collection)
		//Now print the data..
		r, _ := json.Marshal(ravi)
		fmt.Println(string(r))

		UserList[0] = user
		UserList[1] = ravi

		return nil
	}

	ds.Update = func(item interface{}, original interface{}) error {
		return nil
	}

	_, err := database.NewDataStorage(ds)
	if err != nil{
		panic(err)
	}
*/
	//Now create controllers here..and other items..


}

//Run server here..
func main(){
	fmt.Println("Running server")

}
