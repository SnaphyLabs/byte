package controllers

import "gopkg.in/mgo.v2"

type (


	//Inherit controller interface..
	//Define the user controller
	UserController struct {
		session *mgo.Session
		//Each controller will inherit a Connection struct..
		Controller
	}


)


func (u *UserController)getUserById(id string) ()



