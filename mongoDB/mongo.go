package mongoDB

import (
	"time"
	"github.com/SnaphyLabs/rest-layer/resource"
)

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





type FindFn func(lookup interface{}, offset, limit int) (interface{}, error)
type UpdateFn func(item, original  interface{}) error
type InsertFn func(item interface{}) error
type DeleteFn func(item interface{}) error
type ClearFn func(lookup *resource.Lookup) (int, error)
type ConnectFn func() (interface{}, error)
type CloseFn func() (error)





