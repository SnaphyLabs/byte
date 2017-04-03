package api_management

import (
	"github.com/graphql-go/graphql"
	"github.com/SnaphyLabs/SnaphyByte/controllers"
	"github.com/SnaphyLabs/SnaphyByte/interfaces"
	"fmt"
	"errors"
	"github.com/SnaphyLabs/SnaphyByte/models"
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
func (r *Resource)addField(field *Field, rs subResources) error {
	//TODO: Convert field name to lower case.
	//TODO: Field name will always be case insensitive..
	if r.fields == nil{
		r.fields = make(map[string]*Field)
	}

	if field.Name == ""{
		return errors.New("Field name cannot be empty")
	}

	if r.schema == nil{
		return errors.New("Resource's Schema is required to add field.")
	}

	if f, ok := r.fields[field.Name]; f == nil || !ok{
		r.fields[field.Name] = field
	}else{
		return fmt.Errorf("Field %s already present in %s Resource", field.Name, r.name)
	}

	//Initializing a graphql field..
	gqlField := &graphql.Field{
		Name: field.Name,
		Description: field.Description,
	}

	switch field.Type {
	case "ID":
		if field.Null == false{
			gqlField.Type = graphql.NewNonNull(graphql.ID)
		}else{
			gqlField.Type = graphql.ID
		}
	case "Int":
		if field.Null == false{
			gqlField.Type = graphql.NewNonNull(graphql.Int)
		}else{
			gqlField.Type = graphql.Int
		}
	case "Float":
		if field.Null == false{
			gqlField.Type = graphql.NewNonNull(graphql.Float)
		}else{
			gqlField.Type = graphql.Float
		}
	case "String":
		if field.Null == false{
			gqlField.Type = graphql.NewNonNull(graphql.String)
		}else{
			gqlField.Type = graphql.String
		}
	case "Boolean":
		if field.Null == false{
			gqlField.Type = graphql.NewNonNull(graphql.Boolean)
		}else{
			gqlField.Type = graphql.Boolean
		}
	case "Enum":
		//TODO: handle for enum type
	case "Union":
		//TODO: handle for union type
	case "Interface":
		//Handle for Interface type.
	default:
		if customType, ok := rs[field.Type]; ok{
			if field.Null == false{
				gqlField.Type = graphql.NewNonNull(customType.schema)
			}else{
				gqlField.Type = customType.schema
			}
		}else{
			return fmt.Errorf("Unknown type %s in field definition", field.Type)
		}
	}


	gqlField.Resolve = func(p graphql.ResolveParams) (interface{}, error) {
		if model, ok := p.Source.(*models.BaseModel); ok {
			return model.Payload[field.Name], nil
		}
		return nil, nil
	}

	//TODO: Future handling for
	/**
	Null bool
	Validation interface{}
	ReadOnly bool
	Default string
	Unique bool
	Hidden bool
	Resolve interface{}
	Args interface{}
	 */
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

	//Now create a new schema of type graphql.Object
	schema := graphql.NewObject(graphql.ObjectConfig{
		Name: rc.Name,
		Description: rc.Description,
		Fields: graphql.Fields{
			//Blank fields to be added dynamically..
		},
	})

	//Now add schema to Resource..
	r.schema = schema


	//Now add controller..
	//TODO: Remove collection name from controller..if necessary
	if ctrl, err := controllers.NewController(rc.Name, rc.Storage); err != nil{
		return nil, err
	}else{
		r.controller = ctrl
	}

	if rc.Fields == nil{
		return  nil, errors.New("Fields cannot be empty")
	}else{
		if len(rc.Fields) == 0{
			return nil, errors.New("Fields cannot be empty")
		}
	}

	for _, f := range rc.Fields{
		if err := r.addField(f); err != nil{
			return nil, err
		}
	}



	//TODO: add hooks..etc..
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

