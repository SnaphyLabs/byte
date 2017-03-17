package schemas



import (
	"time"
	"github.com/graphql-go/graphql"
	"github.com/SnaphyLabs/SnaphyByte/SchemaInterfaces"
)

type ErrorMessages struct {
	Message string
}

func (e *ErrorMessages) Error() string {
	return e.Message
}



var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Interfaces:[] *graphql.Interface{
		SchemaInterfaces.BaseModelInterface,
	},
	Fields: graphql.Fields{
		"Id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"Created": &graphql.Field{
			Type: graphql.String,
			Description: "Get the time of the data created",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source["Created"] != nil {
					return p.Source["Created"], nil
				}else{
					t := time.Now().UTC()
					return t.String(), nil
				}
			},
		},
		"Updated": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Description: "Store the last updated time",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source["Updated"] != nil {
					return p.Source["Updated"], nil
				}else{
					t := time.Now().UTC()
					return t.String(), nil
				}
			},
		},
		"Type": &graphql.Field{
			Type:graphql.NewNonNull(graphql.String),
			Description:"Store the collection type of User Model",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source["Type"], nil;
			},
		},
		"Password": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			//TODO: Password must be stored using bcrypt encryption..
			Description: "Password of user.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if p.Source["Id"] == nil{
					if p.Source["Password"] == nil || p.Source["Password"] == ""{
						return nil, &ErrorMessages{
							Message: "Password is required",
						}
					}
				}

				return p.Source["Password"], nil

			},
		},
	},
})

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
			//curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:"My+new+todo"){id,text,done}}'


		"createTodo": &graphql.Field{
			Type:        todoType, // the return type for this field
			Description: "Create new todo",
			Args: graphql.FieldConfigArgument{
				"text": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (user interface{}, err error) {

				return user, err
			},
		},
			//curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:"a",done:true){id,text,done}}'


		"updateTodo": &graphql.Field{
			Type:        todoType, // the return type for this field
			Description: "Update existing todo, mark it done or not done",
			Args: graphql.FieldConfigArgument{
				"done": &graphql.ArgumentConfig{
					Type: graphql.Boolean,
				},
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (list interface{}, err error) {
				return list, err
			},
		},
	},
})
