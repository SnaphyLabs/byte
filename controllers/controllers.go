package controllers

import "github.com/SnaphyLabs/SnaphyByte/database"

type (

	//Interface defining controller connection..
	ControllerProvider interface {
		//Create a new session for concurrent query
		NewSession() (*database.DbSession, error)
	}

	//Connection implements ConnectionProvider
	Controller struct {
		//Define a Db Session of AnyType
		dbSession *database.DbSession
	}
)


func (ctrl *Controller) NewSession() (*database.DbSession, error) {
	//Get the interface method
	session := database.DbSession(ctrl.dbSession)
	 return session.Copy(), nil
}
