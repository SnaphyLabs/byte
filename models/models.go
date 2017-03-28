package models

import (
	"time"
	"errors"
	"strings"
	"github.com/rs/xid"
	"encoding/json"
	"fmt"
	"crypto/md5"
)


type (

	//Interface defining base model..
	ModelInterface interface {
		//Fetch a model..
		get() (ModelInterface, error)
		//Save data to server..
		save() (error)
		//Destroy model to server..
		destroy() error
		//Perform initialization on model. removed.
		//init() error
		//Create a new instance of model with IdProperty, Created, Updated, Etag, Type
		NewModel(collectionType string) (error)
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
	//Generate an Id for model
	guid := xid.New()
	b.ID = guid.String()

	b.Created = time.Now()
	b.Updated = time.Now()
	//Generate the unique etag for current Data
	if eTag, err := GenEtag(b); err != nil{
		return err
	}else{
		b.ETag = eTag
	}

	return nil
}



// Etag computes an etag based on containt of the payload
func GenEtag(modelInterface ModelInterface) (string, error) {
	b, err := json.Marshal(modelInterface)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5.Sum(b)), nil
}



//Reload the model from the server. Id value must be present on the model..
func (b *BaseModel) get() (ModelInterface, error){
	mp := ModelInterface(b)
	return mp, nil
}



//Save the model to the server..
func (b *BaseModel) save() (error){
	//TODO: Save data to database..

	return nil
}


//Delete data from database..
func (b *BaseModel) destroy() error  {
	//TODO: Delete user from database..

	return nil
}

