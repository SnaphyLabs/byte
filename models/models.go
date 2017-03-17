package models

import "github.com/SnaphyLabs/SnaphyByte/database"

type (

	//Interface defining controller connection..
	ModelProvider interface {

	}

	//Connection implements ModelProvider
	Model struct {
		//Define a Db Session of AnyType
		dbSession *database.DbSession
	}
)
