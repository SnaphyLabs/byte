package controllers

import (
	"github.com/SnaphyLabs/SnaphyByte/interfaces"
	"errors"
	"github.com/SnaphyLabs/SnaphyByte/models"
	"golang.org/x/net/context"
	"github.com/SnaphyLabs/SnaphyByte/resource"
	"github.com/SnaphyLabs/SnaphyByte/collections"
	"strings"
	"fmt"
)



type (

	//Interface defining base model..
	ControllerInterface interface {
		//Fetch a model..
		FindById(ctx context.Context, id interface{}, lookup *resource.Lookup) (*models.BaseModel, error)
		//Find data from database....
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
		//UpdateAll(ctx context.Context, lookup *resource.Lookup, model *models.BaseModel) error
		//Update a model by its property
		Update(ctx context.Context, model *models.BaseModel) error
		//Count the total number of model present in database....
		Count(ctx context.Context, lookup *resource.Lookup) (int, error)
		//Clear all data found in query..
		Clear(ctx context.Context, lookup *resource.Lookup)(int, error)
		//Destroy and Item By Id
		//DestroyById(ctx context.Context, id string) error
		DestroyById(ctx context.Context, id interface{}, eTag string) error
		//Check if a item is present in database..
		Exists(ctx context.Context,  id interface{}) (bool, error)
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
	collection = strings.TrimSpace(collection)
	if collection != "" && storage != nil{
		return &Controller{
			collection_type: strings.ToLower(collection),
			storage: storage,
		}, nil
	}else{
		return nil, errors.New("Collection type and storage is required for a controller init")
	}
}




//Find a model by id
func (c *Controller) FindById(ctx context.Context, id interface{}, lookup *resource.Lookup) (*models.BaseModel, error) {
	if id == nil{
		return nil, errors.New("Id property is required")
	}else{
		lookup.Filter().AppendQuery(map[string]interface{}{
			"id": id,
		})

		if l, err := c.storage.Find(ctx, lookup, 0, 1); err == nil{
			if l == nil{
				return nil, nil
			}else{
				if l.Total > 0{
					return l.Models[0], nil
				}
			}
		}
	}

	return nil, nil
}



//Find list of items from database..
func (c *Controller) Find(ctx context.Context, lookup *resource.Lookup, offset, limit int) (*collections.BaseModelList, error){
	//Simply return the base methods..
	return c.storage.Find(ctx, lookup, offset, limit)
}


//Find one data per query
func (c *Controller) FindOne(ctx context.Context, lookup *resource.Lookup, offset int) (*models.BaseModel, error){
	if l, err := c.storage.Find(ctx, lookup, offset, 1); err != nil{
		return nil, err
	}else{
		if l.Total > 0{
			return l.Models[0], nil
		}else{
			return nil, nil
		}
	}
}



//If data found the
//Find or the value or populate the data
//TODO: TEST it chance of error..
func (c *Controller) FindOrCreate(ctx context.Context, lookup *resource.Lookup, offset int, model *models.BaseModel) error {
	if b, err := c.FindOne(ctx, lookup, offset);  err == nil{
		if b == nil{
			if err = c.Create(ctx, model); err != nil{
				return err
			}
		}else{
			//Swap the value in the given pointer..
			*model = *b
		}
	}else{
		return err
	}
	//All is well. Now return the value..
	return nil
}




//Create a model in database
func (c *Controller) Create(ctx context.Context, model *models.BaseModel) error{
	if model.ID == nil{
		model.NewModel(c.collection_type)
	}

	//Assigns collection type..
	if model.Type != c.collection_type {
		return fmt.Errorf("Mismatched collection type. Expected of type %s", c.collection_type)
	}

	if model.Type == "" {
		return errors.New("Model Type cannot be empty")
	}

	//Now insert model to database..
	return c.storage.Insert(ctx, []*models.BaseModel{model})
}



//Create a list of models in database...
func (c *Controller) CreateAll(ctx context.Context, models []*models.BaseModel) error {
	for _, model := range models{
		if model.ID == nil{
			model.NewModel(c.collection_type)
		}

		//Assigns collection type..
		if model.Type != c.collection_type {
			return fmt.Errorf("Mismatched collection type. Expected of type %s", c.collection_type)
		}

		if model.Type == "" {
			return errors.New("Model Type cannot be empty")
		}
	}

	//Now insert model to database..
	return c.storage.Insert(ctx, models)
}



//Update a model if present or create if present in database..
func (c *Controller)Upsert(ctx context.Context, model *models.BaseModel) error{
	l := &resource.Lookup{}
	//First find model by id
	if ori, err := c.FindById(ctx, model.ID, l); err != nil{
		return err
	}else{
		if ori == nil {
			//Insert...
			return c.Create(ctx, model)
		}else{
			//Update the model..
			return c.Update(ctx, model, ori)
		}
	}

}
/*



//Update a collection with where query..
func (c *Controller)UpdateAll(ctx context.Context, lookup *resource.Lookup, model *models.BaseModel) error{
	//TODO: Work in progress...
	return c.Update(ctx, model, ori)

}
*/


//Update a model by its property
func (c *Controller)Update(ctx context.Context, model *models.BaseModel, ori *models.BaseModel) error{
	//Generate new Etag for new model
	model.Update()
	return c.storage.Update(ctx, model, ori)

}



//Count the total number of model present in database....
func (c *Controller)Count(ctx context.Context, lookup *resource.Lookup) (int, error){
	return c.storage.Count(ctx, lookup)
}



//Clear all data found in query..
func (c *Controller) Clear(ctx context.Context, lookup *resource.Lookup)(int, error){
	return c.storage.Clear(ctx, lookup)
}



//Destroy and Item By Id
func (c *Controller)DestroyById(ctx context.Context, id interface{}, eTag string) error{
	if id == nil || eTag == ""{
		return errors.New("Id or Etag cannot be empty")
	}

	l := &resource.Lookup{}
	//Insert query for id and etag.
	l.Filter().AppendQuery(map[string]interface{}{
		"id": id,
		"_etag": eTag,
	})

	if n, err := c.Clear(ctx, l); err != nil{
		return err
	}else{
		if n == 0{
			//Check if error conflict or not found..
			l := &resource.Lookup{}
			//Insert query for id and etag.
			l.Filter().AppendQuery(map[string]interface{}{
				"id": id,
			})

			if count, err := c.Count(ctx, l); err != nil{
				return err
			}else{
				if count == 0{
					return resource.ErrNotFound
				}else{
					//Data conflicting with
					return resource.ErrConflict
				}
			}
		}
	}
	return nil
}




//Check if a item is present in database..
func (c *Controller)Exists(ctx context.Context,  id interface{}) (bool, error){
	if id == nil{
		return false, errors.New("Id is required")
	}

	//Check if error conflict or not found..
	l := &resource.Lookup{}
	//Set fields with only id property..
	l.AddField([]string{"_id", })

	//Insert query for id and etag.
	l.Filter().AppendQuery(map[string]interface{}{
		"id": id,
	})

	if count, err := c.Count(ctx, l); err != nil{
		return false, err
	}else{
		if count > 0 {
			return true, nil
		}
	}
	return false, nil
}



//Create a new instance of model with IdProperty, Created, Updated, Etag, Type
func (c *Controller)NewModel() (*models.BaseModel, error){
	b := &models.BaseModel{}
	if c.collection_type != "" {
		b.NewModel(c.collection_type)
	}else{
		return nil, errors.New("Collection Type is required")
	}
	return b, nil
}








