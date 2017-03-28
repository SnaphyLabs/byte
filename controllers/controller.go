package controllers

import (
	"github.com/SnaphyLabs/SnaphyByte/interfaces"
	"errors"
	"github.com/SnaphyLabs/SnaphyByte/models"
	"golang.org/x/net/context"
	"github.com/SnaphyLabs/SnaphyByte/resource"
	"github.com/SnaphyLabs/SnaphyByte/collections"
)



type (

	//Interface defining base model..
	ControllerInterface interface {
		//Fetch a model..
		FindById(ctx context.Context, id string) (*models.BaseModel, error)
		//Save data to server..
		Find(ctx context.Context, lookup *resource.Lookup, offset, limit int) (*collections.BaseModelList, error)
		//Find one model from database..
		FindOne(ctx context.Context, lookup *resource.Lookup, offset int) (*models.BaseModel, error)
		//Find and populate the data
		FindOrCreate(ctx context.Context, lookup *resource.Lookup, offset int, model *models.BaseModel) error
		//Create a model in database
		Create(ctx context.Context, model *models.BaseModel) error
		//Update a model if present or create if present in database..
		Upsert(ctx context.Context, model *models.BaseModel) error
		//Update a collection with where query..
		UpdateAll(ctx context.Context, lookup *resource.Lookup, model *models.BaseModel) error
		//Update a model by its property
		Update(ctx context.Context, model *models.BaseModel) error
		//Count the total number of model present in database....
		Count(ctx context.Context, lookup *resource.Lookup) (int, error)
		//Clear all data found in query..
		Clear(ctx context.Context, lookup *resource.Lookup)(int, error)
		//Destroy and Item By Id
		//DestroyById(ctx context.Context, id string) error
		DestroyById(ctx context.Context, id, eTag string) error
		//Check if a item is present in database..
		Exists(ctx context.Context,  id string) (bool, error)
		//Create a new instance of model with IdProperty, Created, Updated, Etag, Type
		NewModel() (*models.BaseModel, error)
	}


	//Controller implements ControllerInterface
	Controller struct {
		//Name of user virtual models collection.
		collection_type string
		//Database interface which contains all the method of find, create, update, delete, clear.
		//Each storage is already assosiated with a database`s database and database collection..
		storage database.Storage
	}
)//type




//creates a new controller based on given collection and Storage..
func NewCollection(collection string, storage database.Storage) (*Controller, error){
	if collection != "" && storage != nil{
		return &Controller{
			collection_type: collection,
			storage: storage,
		}, nil
	}else{
		return nil, errors.New("Collection type and storage is required for a controller init")
	}
}




