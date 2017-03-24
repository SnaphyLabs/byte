package controllers

import "github.com/SnaphyLabs/SnaphyByte/database"

type (

	//Interface defining controller connection..
	ControllerProvider interface {

	}

	//Connection implements ConnectionProvider
	Controller struct {
		//Store datastorage..value
		//database.DataStorage
	}
)



//Will generate a new session for concurrent query.
func (ctrl *Controller) NewSession() (*database.DbSession, error) {
	//Get the interface method
	//session := database.DbSession(ctrl.dbSession)
	 //return session.Copy(), nil
}
