package database

import "context"

//All Database Session will inherit this interface
type Storage interface {
	//Find and return list of model found in db.
	Find(filter interface{}, output interface{}) (error)
	//Update an model
	Update(item, original  interface{}) error
	//Insert an item or list of items with this query..
	Insert(ctx context.Context, item  interface{}, output interface{}) error
	//Delete a single item by matched query
	Delete(item  interface{}) error
	//Clear a data with matched query
	Clear(filter interface{}) (int, error)
	//Count the the data..
	Count(filter interface{}) (int, error)
}


