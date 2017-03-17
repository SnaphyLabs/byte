package database

import (
	"gopkg.in/mgo.v2"
	"time"
	"github.com/SnaphyLabs/SnaphyByte/controllers"
	"github.com/rs/rest-layer/resource"
)

//All Database Session will inherit this interface
type DbSession interface {
	//Will create a new session
	Copy() *DbSession
	Find(lookup , offset, limit int)
	Update(item *resource.Item, original *resource.Item) error
	Delete(item *resource.Item) error
	Clear(lookup *resource.Lookup) (int, error)
}


type Storage interface {
	GetSession() *interface{}
	//Each controller must inherit ConnectionProvider interface.
	NewController() *controllers.ControllerProvider
}

//MongoStorage implements interface Storage
//MongoStorage implements interface DbSession
type MongoStorage struct {
	Address []string
	Timeout time.Duration
	Database string
	Username string
	Password string
	Session *mgo.Session
}


func (ms *MongoStorage)Copy() *mgo.Session {
	return ms.Session.Copy()
}



func (ms *MongoStorage)NewController() *controllers.ControllerProvider   {
	
}

//Connect to database and return a session
//GetSession method will get the session for database communication
func (ms *MongoStorage) GetSession() *mgo.Session {
	info := &mgo.DialInfo{
		Addrs:    ms.Address,
		Database: ms.Database,
		Username: ms.Username,
		Password: ms.Password,
	}

	if ms.Timeout != 0 {
		info.Timeout  = ms.Timeout
	}

	session, err := mgo.DialWithInfo(info)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	ms.Session = session
	return session
}
