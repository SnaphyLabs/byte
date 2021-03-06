package common

import (
	"github.com/graphql-go/graphql"
	"github.com/SnaphyLabs/SnaphyByte/controllers"
	"github.com/SnaphyLabs/SnaphyByte/interfaces"
	"fmt"
	"errors"
	"github.com/SnaphyLabs/SnaphyByte/models"
)



type(
	Model struct {
		name          string
		description   string
		fields        map[string]*Field
		hooks         interface{} //TODO: Later implementation..
		//interfaces  []interface{}
		graphql       *graphql.Object
		validator     interface{}
		relation      interface{} //It should be name of related model and model relation type..
		interfaces    []*graphql.Interface
	}



	//All model defined in one model list..
	ModelConfig struct {
		models     	map[string]*Model
		controller    	*controllers.Controller
		storage       	database.Storage
		interfaces      []*graphql.Interface
	}


	//Used to create a model
	RuleConfig struct{
		Name 		string
		Description 	string
		Fields 		[]*Field
		Interfaces      []*graphql.Interface
	}
)


//--------------------------------------------SUBRESOURCE METHOD--------------------------------------------------
//Generate new ModelConfig
func NewModelConfig(storage database.Storage) (*ModelConfig, error) {
	mc := &ModelConfig{
		storage: storage,

	}

	//Now add controller..
	//TODO: Remove collection name from controller..if necessary
	if ctrl, err := controllers.NewController(mc.storage); err != nil{
		return nil, err
	}else{
		mc.controller = ctrl
	}
	return mc, nil
}



//Add a new Resources at runtime....
func (mc *ModelConfig)add(rsrc *Model) error{
	if mc.models == nil{
		mc.models = make(map[string]*Model)
	}

	if r, ok := mc.models[rsrc.name]; r == nil || !ok{
		mc.models[rsrc.name] = rsrc
	}else{
		return fmt.Errorf("Model %s already present.", rsrc.name)
	}
	return nil
}


func (mc *ModelConfig)Models() map[string]*Model{
	return mc.models
}



// new creates a new model with provided spec, handler and config
//Assosiates the newly created resource with the subresources..
func (sr *ModelConfig)NewModel(rc *RuleConfig) (error) {
	m := &Model {
		name: 		rc.Name,
		description: 	rc.Description,
		interfaces:     rc.Interfaces,
	}




	//Now create a new schema of type graphql.Object
	schema := graphql.NewObject(graphql.ObjectConfig{
		Name: rc.Name,
		Description: rc.Description,
		Fields: graphql.Fields{
			//Blank fields to be added dynamically..
		},
		Interfaces: m.interfaces,
	})


	//Now add schema to Resource..
	m.graphql = schema
	//TODo: Remove..


	if rc.Fields == nil{
		return  errors.New("Fields cannot be empty")
	}else{
		if len(rc.Fields) == 0{
			return errors.New("Fields cannot be empty")
		}
	}

	for _, f := range rc.Fields{
		if err := m.addField(f, sr); err != nil{
			return err
		}
	}


	//TODO: Handle Relation handling..hasOne, belongsTo, hasMany, hasAndBelongsToMany, hasManyThrough



	//TODO: add hooks..etc..

	//Now add model config to model..
	return sr.add(m)
	//return nil
}






//-----------------------------------------------------END MODEL-CONFIG METHOD---------------------------------------------------------------------




//---------------------------------------------------------MODEL METHOD-------------------------------------------------------------------
//Add Field to a Resources at runtime..
func (r *Model)addField(field *Field, mc *ModelConfig) error {
	///TODO: add Validation of field..
	//TODO: Convert field name to lower case.
	//TODO: Field name will always be case insensitive..
	if r.fields == nil{
		r.fields = make(map[string]*Field)
	}

	if field.Name == ""{
		return errors.New("Field name cannot be empty")
	}

	if r.graphql == nil{
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
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if model, ok := p.Source.(*models.BaseModel); ok {
				return model.Payload[field.Name], nil
			}
			return nil, nil
		},
	}

	switch field.Type {
	case "ID":
		if field.AllowNull == false{
			gqlField.Type = graphql.NewNonNull(graphql.ID)
		}else{
			gqlField.Type = graphql.ID
		}
	case "Int":
		if field.AllowNull == false{
			gqlField.Type = graphql.NewNonNull(graphql.Int)
		}else{
			gqlField.Type = graphql.Int
		}
	case "Float":
		if field.AllowNull == false{
			gqlField.Type = graphql.NewNonNull(graphql.Float)
		}else{
			gqlField.Type = graphql.Float
		}
	case "String":
		if field.AllowNull == false{
			gqlField.Type = graphql.NewNonNull(graphql.String)
		}else{
			gqlField.Type = graphql.String
		}
	case "Boolean":
		if field.AllowNull == false{
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
		if customType, ok := mc.models[field.Type]; ok{
			if field.AllowNull == false{
				gqlField.Type = graphql.NewNonNull(customType.graphql)
			}else{
				gqlField.Type = customType.graphql
			}
		}else{
			return fmt.Errorf("Unknown type %s in field definition", field.Type)
		}
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

	r.graphql.AddFieldConfig(field.Name, gqlField)
	return nil
}

//Returned the graphql object..
func (m *Model)GraphQLObject() *graphql.Object  {
	return m.graphql
}


//Load resources to memory
func (r *Model)load() error  {
	//Load resources to memory
	return nil
}


//Remove resources from memory
func (r *Model)remove() error  {
	//Load resources to memory
	return nil
}


//Reload resources into memory with changes..applied
func (r *Model)reload() error  {
	//Load resources to memory
	return nil
}

