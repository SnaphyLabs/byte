package memoryDB

import (
	"golang.org/x/net/context"
	"github.com/SnaphyLabs/SnaphyByte/database"
	"github.com/SnaphyLabs/SnaphyByte/models"
)

//Creates in Memory database to test SnaphyByte for now. later this to be converted to mongodb.

//MemoeyStorage implements DbSession
type MemoryStorage struct {}


func (m *MemoryStorage) NewHandler(ctx context.Context)(*database.DbSession){

}


func (m *MemoryStorage) Find(lookup interface{}, offset, limit int) ([]interface{}, error){
	//TODO: Implement data find method..
	return models.LocalDatabase[models.USER_COLLECTION], nil
	//return nil, nil
}



func (m *MemoryStorage) Update(item, original  interface{}) error  {
	//TODO: Work in progress...
	//models.LocalDatabase[models.USER_COLLECTION][6] = item;
	return nil
}



func (m *MemoryStorage) Insert(item  interface{}) error  {
	//TODO: Work in progress..

	return nil
}

func (m *MemoryStorage) Delete(item  interface{}) error  {
	//TODO: Work in progress..

	return nil
}



func (m *MemoryStorage) Clear(lookup interface{}) (int, error)  {
	//TODO: Work in progress..

	return 0, nil
}

