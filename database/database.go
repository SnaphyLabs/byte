package database

import (
	"time"
	"github.com/rs/rest-layer/resource"
)


//All Database Session will inherit this interface
type DbSession interface {
	Find(lookup , offset, limit int)
	Update(item *resource.Item, original *resource.Item) error
	Insert(item *resource.Item) error
	Delete(item *resource.Item) error
	Clear(lookup *resource.Lookup) (int, error)

}


/*

type Storage interface {

	GetSession() *DbSession
	//Each controller must inherit ConnectionProvider interface.
	NewController(c controllers.ControllerProvider) error
}
*/


//MongoStorage implements interface DbSession
type DataStorage struct {
	//Url Address
	Address []string
	Timeout time.Duration
	Database string
	Username string
	Collection string
	Password string
	Session *interface{}

	//Define abstract methods..
	Find FindFn
	Update UpdateFn
	Insert InsertFn
	Delete DeleteFn
	Clear ClearFn
	Connect ConnectFn
	Close CloseFn
}


type FindFn func(lookup , offset, limit int)
type UpdateFn func(item *resource.Item, original *resource.Item) error
type InsertFn func(item *resource.Item) error
type DeleteFn func(item *resource.Item) error
type ClearFn func(lookup *resource.Lookup) (int, error)
type ConnectFn func() (*interface{}, error)
type CloseFn func() (error)



//Connect to database and return a session with database..
//Create a new Data Storage function with connect function..
func NewDataStorage(storage *DataStorage) (*DataStorage, error){
	session, err := storage.Connect()
	if err != nil{
		return nil, err
	}
	storage.Session = session
	return storage, nil
}




/*//Generate a new controller
func (ms *DataStorage)NewController(c controllers.ControllerProvider) error   {
	session := ms.Connect()
	c["dbSession"] = session
	return nil
}*/

/*
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
*/
