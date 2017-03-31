package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/SnaphyLabs/SnaphyByte/models"
	"errors"
	"github.com/SnaphyLabs/SnaphyByte/schemaInterfaces"
	"github.com/SnaphyLabs/SnaphyByte/controllers"
)



func init(){
	
	controllers.NewCollection()
}



var (
	//Define a user type...
	UserType = graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				Description: "Unique Id of user type.",
				/*Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "Get user by Id",
						Type: graphql.NewNonNull(graphql.ID),
					},
				},*/
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(models.BaseModel); ok {
						return model.ID, nil
					}
					return nil, nil
				},
			},
			"eTag": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				/*Args: graphql.FieldConfigArgument{
					"eTag": &graphql.ArgumentConfig{
						Description: "Etag of model",
						Type:  graphql.NewNonNull(graphql.String),
					},
				},*/
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(models.BaseModel); ok {
						return model.ETag, nil
					}
					return nil, nil
				},
			},
			"created": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(models.BaseModel); ok {
						return model.Created, nil
					}
					return nil, nil
				},
			},
			"updated": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(models.BaseModel); ok {
						return model.Updated, nil
					}
					return nil, nil
				},
			},
			"payload": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Object{}),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(models.BaseModel); ok {
						return model.Payload, nil
					}
					return nil, nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			schemaInterfaces.BaseModelInterface,
		},
	})



	//Define a user type...
	BookType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				Description: "Unique Id of user type.",
				/*Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "Get user by Id",
						Type: graphql.NewNonNull(graphql.ID),
					},
				},*/
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(models.BaseModel); ok {
						return model.ID, nil
					}
					return nil, nil
				},
			},
			"eTag": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				/*Args: graphql.FieldConfigArgument{
					"eTag": &graphql.ArgumentConfig{
						Description: "Etag of model",
						Type:  graphql.NewNonNull(graphql.String),
					},
				},*/
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(models.BaseModel); ok {
						return model.ETag, nil
					}
					return nil, nil
				},
			},
			"created": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(models.BaseModel); ok {
						return model.Created, nil
					}
					return nil, nil
				},
			},
			"updated": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(models.BaseModel); ok {
						return model.Updated, nil
					}
					return nil, nil
				},
			},
			"payload": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Object{}),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(models.BaseModel); ok {
						return model.Payload, nil
					}
					return nil, nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			schemaInterfaces.BaseModelInterface,
		},
	})


	//QueryType
	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getAuthor": &graphql.Field{
				Type: UserType,
				Description:"Returns author of the book",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "Return author of the book by id",
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetHero(p.Args["episode"]), nil
				},
			},
		},
	})




)


