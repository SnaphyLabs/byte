package api_management

import (
	"github.com/graphql-go/graphql"
	"github.com/SnaphyLabs/SnaphyByte/controllers"
	"github.com/SnaphyLabs/SnaphyByte/interfaces"
	"fmt"
)

type(
	Resource struct {
		name         string
		description  string
		fields       map[string]*Field
		hooks        interface{} //TODO: Later implementation..
		interfaces   []interface{}
		schema       *graphql.Object
		validator    interface{}
		resources    subResources
		controller   *controllers.Controller
		storage      database.Storage
	}

	subResources map[string]*Resource


	ResourceConfig struct{
		Name 		string
		Description 	string
		Fields 		[]*Field
		Storage      	database.Storage
	}
)



//Add Field to a Resources at runtime..
func (r *Resource)addField(field *Field) error {
	if r.fields == nil{
		r.fields = make(map[string]*Field)
	}

	if f, ok := r.fields[field.Name]; f == nil || !ok{
		r.fields[field.Name] = field
	}else{
		return fmt.Errorf("Field %s already present in %s Resource", field.Name, r.name)
	}
	return nil
}



//Add a new Resources at runtime....
func (sr subResources)add(rsrc *Resource) error{
	if r, ok := sr[rsrc.name]; r == nil || !ok{
		sr[rsrc.name] = rsrc
	}else{
		return fmt.Errorf("Resource %s already present.", rsrc.name)
	}
	return nil
}



// new creates a new resource with provided spec, handler and config
func newResource(rc ResourceConfig) (*Resource, error) {
	r := &Resource {
		name: 		rc.Name,
		description: 	rc.Description,
		storage: 	rc.Storage,
	}

	//Now add controller..
	//TODO: Remove collection name from controller..if necessary
	if ctrl, err := controllers.NewController(rc.Name, rc.Storage); err != nil{
		return nil, err
	}else{
		r.controller = ctrl
	}


	return r, nil
}



//Load resources to memory
func (r *Resource)load() error  {
	//Load resources to memory
	return nil
}


//Remove resources from memory
func (r *Resource)remove() error  {
	//Load resources to memory
	return nil
}


//Reload resources into memory with changes..applied
func (r *Resource)reload() error  {
	//Load resources to memory
	return nil
}

