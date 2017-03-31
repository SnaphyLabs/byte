package models

import (
	"time"
	"errors"
	"strings"
	"github.com/rs/xid"
	"github.com/SnaphyLabs/SnaphyUtil"
)


type (
	//Interface defining base model..
	ModelInterface interface {
		//Fetch a model..
		Get() (ModelInterface, error)
		//Save data to server..
		Save() (error)
		//Destroy model to server..
		Destroy() error
		//Perform initialization on model. removed.
		//init() error
		//Create a new instance of model with IdProperty, Created, Updated, Etag, Type
		NewModel(collectionType string) (error)
		//Generate etag for data.. and assosiate it..
		GenEtag() (error)
		//Update model with new Etag and Updated time property.
		Update() error
		//Generate a new Id for the model..
		NewId()
		//Copy and create new model from it..
		Copy() (*BaseModel, error)
	}


	//BaseModel implements ModelProvider
	BaseModel struct {
		// ID is used to uniquely identify the item in the resource collection.
		ID interface{}
		// ETag is an opaque identifier assigned by Snaphy Byte to a specific version of the item.
		//
		// This ETag is used perform conditional requests and to ensure storage handler doesn't
		// update an outdated version of the resource.
		ETag string
		// Updated stores the last time the item was updated. This field is used to populate the
		// Last-Modified header and to handle conditional requests.
		Updated time.Time
		// Updated stores the last time the item was updated. This field is used to populate the
		// Last-Modified header and to handle conditional requests.
		Created time.Time
		//Type signifies the type of the collection..
		Type string
		Payload map[string]interface{}
	}
)//type





func init() {

}


//Create a new or re-initializes a model with Id, Created, Updated and Type Property
func (b *BaseModel) NewModel(collectionType string)  error{
	collectionType = strings.TrimSpace(collectionType)
	if collectionType != "" {
		b.Type = collectionType
	}else{
		return errors.New("Collection Type is required")
	}
	//Generate a new Id..
	b.newId()

	b.Created = time.Now()
	if err := b.Update(); err != nil{
		return err
	}
	return nil
}



//Update model with updated and new Etag property..
//This method just update the Etag and Updated property it doesn't update model to server.
func (b *BaseModel)Update() error  {
	b.Updated = time.Now()
	if err := b.GenEtag(); err != nil{
		return err
	}
	return nil
}


//Generate a new Id for the model
//PRIVATE METHOD for internal use only
func (b *BaseModel) newId()  {
	//Generate an Id for model
	guid := xid.New()
	b.ID = guid.String()
}


//Copy the current model and create a new model with changed Id, Etag and Updated, created property...
func (b *BaseModel) Copy() (*BaseModel, error)   {
	copyb := &BaseModel{}
	copyb.Payload = b.Payload
	if err := copyb.NewModel(b.Type); err != nil{
		return nil, err
	}else{
		return copyb, nil
	}

}


//Generate Etag for the model
func (b *BaseModel) GenEtag() error{
	//Generate the unique etag for current Data
	if eTag, err := SnaphyUtil.GenEtag(b); err != nil{
		return err
	}else{
		b.ETag = eTag
	}

	return nil
}




//Reload the model from the server. Id value must be present on the model..
func (b *BaseModel) Get() (ModelInterface, error){
	//TODO: implement later
	//mp := ModelInterface(b)
	return nil, nil
}



//Save the model to the server..
func (b *BaseModel) Save() (error){
	//TODO: Save data to database..

	return nil
}


//Delete data from database..
func (b *BaseModel) destroy() error  {
	//TODO: Delete user from database..

	return nil
}

