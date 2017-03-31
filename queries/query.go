package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/SnaphyLabs/SnaphyByte/models"
	"github.com/SnaphyLabs/SnaphyByte/schemaInterfaces"
	"github.com/SnaphyLabs/SnaphyByte/controllers"
	"gopkg.in/mgo.v2"
	"time"
	"log"
	"github.com/SnaphyLabs/mongoByte"
	"github.com/SnaphyLabs/SnaphyByte/resource"
)

const (
	MongoDBHosts = "localhost:27017"
	AuthDatabase = "drugcorner"
	AuthUserName = "robins"
	AuthPassword = "12345"
	Collection = "SnaphyModelDefinition"

	BOOK_TYPE = "book"
	AUTHOR_TYPE = "author"
)

func init(){

	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}
	var err error
	mongoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
	//Get a handler for handling data..

	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}




}



var (
	//AuthorController *controllers.Controller
	//BookController *controllers.Controller
	mongoSession *mgo.Session
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
					if model, ok := p.Source.(*models.BaseModel); ok {
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
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.ETag, nil
					}
					return nil, nil
				},
			},
			"created": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Created, nil
					}
					return nil, nil
				},
			},
			"updated": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Updated, nil
					}
					return nil, nil
				},
			},
			"payload": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
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
					if model, ok := p.Source.(*models.BaseModel); ok {
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
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.ETag, nil
					}
					return nil, nil
				},
			},
			"created": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Created, nil
					}
					return nil, nil
				},
			},
			"updated": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Updated, nil
					}
					return nil, nil
				},
			},
			"payload": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
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
					var lookup *resource.Lookup = &resource.Lookup{}
					if AuthorController, err := controllers.NewCollection(AUTHOR_TYPE, mongoByte.NewHandler(mongoSession, AuthDatabase, Collection)); err != nil{
						panic(err)
					}else{
						return AuthorController.FindById(p.Context, p.Args["id"].(string), lookup)
						//return  AuthorController.FindById(p.Context, "b3cdrv8j1n6jakdgbn60", lookup)
					}

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
					var lookup *resource.Lookup = &resource.Lookup{}
					if BookController, err := controllers.NewCollection(BOOK_TYPE, mongoByte.NewHandler(mongoSession, AuthDatabase, Collection)); err != nil{
						panic(err)
					}else{
						return BookController.FindById(p.Context, p.Args["id"].(string), lookup)
					}
				},
			},
		},
	})

	TestSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,

	})


)


