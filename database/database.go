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
	GetSession() *DbSession
	//Each controller must inherit ConnectionProvider interface.
	NewController(c controllers.ControllerProvider) error
}



//MongoStorage implements interface Storage
//MongoStorage implements interface DbSession
type DataStorage struct {
	Address []string
	Timeout time.Duration
	Database string
	Username string
	Password string
	Session *DbSession
}



func (ms *DataStorage)Copy() *DbSession {
	session := DbSession(ms.Session)
	return session.Copy()
}


//Generate a new controller
func (ms *DataStorage)NewController(c controllers.ControllerProvider) error   {
	session := ms.Connect()
	c["dbSession"] = session
	return nil
}


//Connect to database and return a session
//GetSession method will get the session for database communication
func (ms *DataStorage) Connect() *DbSession {
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
	//Check if it implements session interface..
	sessionInterface := DbSession(session)
	ms.Session = &sessionInterface
	return &sessionInterface
}
