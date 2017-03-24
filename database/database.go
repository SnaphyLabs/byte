package database

import (

)


//All Database Session will inherit this interface
type DbSession interface {
	//Find and return list of model found in db.
	Find(lookup interface{}, offset, limit int) (interface{}, error)
	//Update an model
	Update(item, original  interface{}) error
	//Insert an item or list of items with this query..
	Insert(item  interface{}) error
	//Delete a single item by matched query
	Delete(item  interface{}) error
	//Clear a data with matched query
	Clear(lookup interface{}) (int, error)
}




/*//Generate a new controller
func (ms *DataStorage)NewController(c controllers.ControllerProvider) error   {
	session := ms.Connect()
	c["dbSession"] = session
	return nil
}*/

