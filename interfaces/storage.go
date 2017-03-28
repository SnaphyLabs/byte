package database

import (
	"context"
	"github.com/SnaphyLabs/SnaphyByte/resource"
	"github.com/SnaphyLabs/SnaphyByte/collections"
	"github.com/SnaphyLabs/SnaphyByte/models"
)

//All Database Session will inherit this interface
type Storage interface {
	//Find and return list of model found in db.
	Find(ctx context.Context, lookup *resource.Lookup, offset, limit int) (*collections.BaseModelList, error)
	//Update an model..
	Update(ctx context.Context, item *models.BaseModel, original *models.BaseModel) error
	//Insert an item or list of items with this query..
	Insert(ctx context.Context, items []*models.BaseModel) error
	//Delete a single item by matched query
	Delete(ctx context.Context, item *models.BaseModel) error
	//Clear a data with matched query
	Clear(ctx context.Context, lookup *resource.Lookup) (int, error)
	//Count the the data..
	Count(ctx context.Context, lookup *resource.Lookup) (int, error)
}


