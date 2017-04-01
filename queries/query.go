package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/SnaphyLabs/SnaphyByte/models"

	"gopkg.in/mgo.v2"
	"time"
	"log"
	"github.com/SnaphyLabs/mongoByte"
	"github.com/SnaphyLabs/SnaphyByte/controllers"
	"github.com/SnaphyLabs/SnaphyByte/resource"
	"github.com/SnaphyLabs/SnaphyByte/schema"
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



var (
	userType *graphql.Object
	bookType *graphql.Object
	CollectionTypes map[string]*graphql.Object
	CommonPropertyInterface *graphql.Interface
	mongoSession *mgo.Session
	queryType *graphql.Object
	TestSchema graphql.Schema
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




func init(){
	CollectionTypes  = make(map[string]*graphql.Object)
	CommonPropertyInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name: "CommonProperty",
		Description: "An object with an ID, Created, Updated, ETag",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				Description: "Unique Id of model type.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.ID, nil
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
			"eTag": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Args: graphql.FieldConfigArgument{
					"eTag": &graphql.ArgumentConfig{
						Description: "Etag of model",
						Type:  graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.ETag, nil
					}
					return nil, nil
				},
			},
		},

		//Implement Resolve type...return a simple grapql.Object as its has a mixed type of resolvers.
		ResolveType: func(p graphql.ResolveTypeParams) (*graphql.Object){
			if model, ok := p.Value.(*models.BaseModel); ok {
				if model.Type == "author"{
					return userType
				}else{
					return bookType
				}
			}
			return nil
		},
	})


	//Define a user type...
	userType = graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		IsTypeOf: func(p graphql.IsTypeOfParams) bool {
			if model, ok := p.Value.(*models.BaseModel); ok{
				if model.Type == "author"{
					return true
				}
			}

			return false
		},
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				Description: "Unique Id of user type.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.ID, nil
					}
					return nil, nil
				},
			},
			"eTag": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Args: graphql.FieldConfigArgument{
					"eTag": &graphql.ArgumentConfig{
						Description: "Etag of model",
						Type:  graphql.NewNonNull(graphql.String),
					},
				},
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
			CommonPropertyInterface,
		},
	})



	//Define a user type...
	bookType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Book",
		IsTypeOf: func(p graphql.IsTypeOfParams) bool {
			if model, ok := p.Value.(*models.BaseModel); ok{
				if model.Type == "book"{
					return true
				}
			}
			return false
		},
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				Description: "Unique Id of user type.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.ID, nil
					}
					return nil, nil
				},
			},
			"eTag": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Args: graphql.FieldConfigArgument{
					"eTag": &graphql.ArgumentConfig{
						Description: "Etag of model",
						Type:  graphql.NewNonNull(graphql.String),
					},
				},
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
			CommonPropertyInterface,
		},
	})



	//QueryType
	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"findById": &graphql.Field{
				Type: CommonPropertyInterface,
				Description:"Find author of the book",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "Return author of the book by id",
						Type: graphql.NewNonNull(graphql.ID),
					},
					"$collection": &graphql.ArgumentConfig{
						Description: "Type of collection.",
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var lookup *resource.Lookup = &resource.Lookup{
					}

					var query  schema.Query =  schema.Query{}
					query.AppendQuery(p.Args)
					lookup.AddQuery(query)

					if AuthorController, err := controllers.NewCollection(AUTHOR_TYPE, mongoByte.NewHandler(mongoSession, AuthDatabase, Collection)); err != nil{
						panic(err)
					}else{

						return AuthorController.FindById(p.Context, p.Args["id"].(string), lookup)
					}
				},
			},



			/*"getBook": &graphql.Field{
				Type: bookType,
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
			"getAuthor": &graphql.Field{
				Type: userType,
				Description:"Returns author by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "Return  user by id",
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var lookup *resource.Lookup = &resource.Lookup{}
					if authorController, err := controllers.NewCollection(AUTHOR_TYPE, mongoByte.NewHandler(mongoSession, AuthDatabase, Collection)); err != nil{
						panic(err)
					}else{
						return authorController.FindById(p.Context, p.Args["id"].(string), lookup)
					}
				},
			},*/


		},
	})

	TestSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
		Types: []graphql.Type{userType, bookType},

	})
}




