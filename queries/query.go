package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/SnaphyLabs/SnaphyByte/models"
	"errors"
	"github.com/SnaphyLabs/SnaphyByte/schemaInterfaces"
	"github.com/SnaphyLabs/SnaphyByte/controllers"
	"gopkg.in/mgo.v2"
	"time"
	"log"
	"github.com/SnaphyLabs/mongoByte"
	"github.com/SnaphyLabs/SnaphyByte/resource"
	"fmt"
)

var AuthorController *controllers.Controller
var BookController *controllers.Controller


func init(){
	const (
		MongoDBHosts = "localhost:27017"
		AuthDatabase = "drugcorner"
		AuthUserName = "robins"
		AuthPassword = "12345"
		Collection = "SnaphyModelDefinition"

		BOOK_TYPE = "book"
		AUTHOR_TYPE = "author"
	)
	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	//Get a handler for handling data..

	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	if AuthorController, err := controllers.NewCollection(AUTHOR_TYPE, mongoByte.NewHandler(mongoSession, AuthDatabase, Collection)); err != nil{
		panic(err)
	}

	if BookController, err := controllers.NewCollection(BOOK_TYPE, mongoByte.NewHandler(mongoSession, AuthDatabase, Collection)); err != nil{
		panic(err)
	}
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
					var lookup *resource.Lookup
					return AuthorController.FindById(p.Context, p.Args["id"].(string), lookup)
					//return GetHero(p.Args["episode"]), nil
				},
			},

			"getBook": &graphql.Field{
				Type: BookType,
				Description:"Returns book by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "Return  book by id",
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var lookup *resource.Lookup
					return BookController.FindById(p.Context, p.Args["id"].(string), lookup)
					//return GetHero(p.Args["episode"]), nil
				},
			},
		},
	})

	testSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,

	})


)


