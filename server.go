package SnaphyByte

import (
	"github.com/SnaphyLabs/SnaphyByte/database"
	"time"
	"github.com/rs/rest-layer/resource"
	"github.com/SnaphyLabs/rest-layer/resource"
)

const (
	MongoDBHosts = "localhost:27017"
	Database = "drugcorner"
	UserName = "robins"
	Password = "12345"
	Collection =  "Demo"
)


func init(){
	//Connect to mongodb database..
	ds := &database.DataStorage{
		Address:[]string{MongoDBHosts},
		Database:Database,
		Username:UserName,
		Password:Password,
		Collection:Collection,
		Timeout:60 * time.Second,
	}

	//Now define all the related methods
	ds.Find = func(lookup, offset, limit int)  {
		
	}
	
	ds.Clear = func(lookup *resource.Lookup) (int, error) {
		return 0, nil
	}
	
	ds.Connect = func() (*interface{}, error) {
		i := new(interface{})
		return i, nil
	}
	
	ds.Delete = func(item *resource.Item) error {
		return nil
	}
	
	ds.Insert = func(item *resource.Item) error {
		return nil
	}
	
	ds.Update = func(item *resource.Item, original *resource.Item) error {
		return nil
	}

	_, err := database.NewDataStorage(ds)
	if err != nil{
		panic(err)
	}

	//Now create controllers here..and other items..

}
