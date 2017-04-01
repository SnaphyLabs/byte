package queries

import (
	"github.com/graphql-go/graphql"

	"gopkg.in/mgo.v2"
	"time"
	"log"
	"github.com/SnaphyLabs/mongoByte"
	"github.com/SnaphyLabs/SnaphyByte/controllers"
	"github.com/SnaphyLabs/SnaphyByte/resource"
	"github.com/SnaphyLabs/SnaphyByte/schema"
	"fmt"
	//b64 "encoding/base64"
	"github.com/SnaphyLabs/SnaphyByte/models"
	"github.com/SnaphyLabs/SnaphyUtil"
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

	//CommonPropertyInterface *graphql.Interface
	baseModelInterface *graphql.Interface
	payloadInterface *graphql.Interface
	mongoSession *mgo.Session
	queryType *graphql.Object
	payloadDataType *graphql.Object
	payloadInfoType *graphql.Object
	baseModelType *graphql.Object
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
	baseModelInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name: "BaseModelInterface",
		Description: "Base model interface type",
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
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.ETag, nil
					}
					return nil, nil
				},
			},
			"cursor": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					//TODO: Resolve this method..
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.ETag, nil
					}
					return nil, nil
				},
			},
		},
	})



	//Return a information related to payload interface..
	payloadInfoType = graphql.NewObject(graphql.ObjectConfig{
		Name: "PayloadInfo",
		Description: "Payload info could be of any type of collection",
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
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.ETag, nil
					}
					return nil, nil
				},
			},
			"type": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Type, nil
					}
					return nil, nil
				},
			},
			"cursor": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					//Encode cursor..
					if model, ok := p.Source.(*models.BaseModel); ok {
						return SnaphyUtil.Base64Encode(model.ETag), nil
					}
					return nil, nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			baseModelInterface,
		},
	})


	//Interface for payload
	payloadInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name: "PayloadInterface",
		Description: "Payload model could be of any type of collection",
		Fields: graphql.Fields{
			//Can accept fields of variable type..
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.ID, nil
					}
					return nil, nil
				},
			},
		},
		//Implement Resolve type...return a simple grapql.Object as its has a mixed type of resolvers.
		ResolveType: func(p graphql.ResolveTypeParams) (*graphql.Object){
			//Resolve type of interface here..
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



	baseModelType = graphql.NewObject(graphql.ObjectConfig{
		Name: "BaseModel",
		Description: "Payload model could be of any type of collection",
		Fields: graphql.Fields{
			//Info implements info interface..
			"info":{
				Type: payloadInfoType,
				Description: "Contains information related to model",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model, nil
					}
					return nil, nil
				},

			},
			//Payload implements payload interface..
			"payload":{
				Description: "Contains the model data of any type.",
				Type: payloadInterface,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model, nil
					}
					return nil, nil
				},
			},

		},

	})



	//Define a user type...
	userType = graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
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
			"firstName": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Payload["firstName"], nil
					}
					return nil, nil
				},
			},
			"lastName": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Payload["lastName"], nil
					}
					return nil, nil
				},
			},
			"email": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Payload["email"], nil
					}
					return nil, nil
				},
			},
			"password": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Payload["password"], nil
					}
					return nil, nil
				},
			},
			"userName": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Payload["userName"], nil
					}
					return nil, nil
				},
			},
			"age": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Payload["age"], nil
					}
					return nil, nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			payloadInterface,
		},
	})



	//Define a user type...
	bookType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Book",
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
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Payload["name"], nil
					}
					return nil, nil
				},
			},
			"pages": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Payload["pages"], nil
					}
					return nil, nil
				},
			},
			"price": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Payload["price"], nil
					}
					return nil, nil
				},
			},
			"authorId": &graphql.Field{
				Type: graphql.String,
				Description: "Identity of author who has written this model",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if model, ok := p.Source.(*models.BaseModel); ok {
						return model.Payload["authorId"], nil
					}
					return nil, nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			payloadInterface,
		},
	})


	//QueryType
	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			/*"findById": &graphql.Field{
				Type: baseModelType,
				Description:"Find author of the book",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "Return author of the book by id",
						Type: graphql.NewNonNull(graphql.ID),
					},
					"collection": &graphql.ArgumentConfig{
						Description: "Collection type",
						Type: graphql.NewNonNull(graphql.String),
					},
				},

				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var lookup *resource.Lookup = &resource.Lookup{}
					var q *schema.Query = &schema.Query{}
					q.AppendQuery(p.Args)
					lookup.AddQuery(*q)
					if AuthorController, err := controllers.NewCollection(AUTHOR_TYPE, mongoByte.NewHandler(mongoSession, AuthDatabase, Collection)); err != nil{
						panic(err)
					}else{
						return AuthorController.FindById(p.Context, p.Args["id"].(string), lookup)
					}
				},
			},*/
			"find": &graphql.Field{
				Type: graphql.NewList(baseModelType),
				Description:"Find list of author of the book",
				/*Args: graphql.FieldConfigArgument{
					*//*"collection": &graphql.ArgumentConfig{
						Description: "Collection type",
						Type: graphql.NewNonNull(graphql.String),
					},*//*
					*//*"offset": &graphql.ArgumentConfig{
						Description: "Offset",
						Type: graphql.NewNonNull(graphql.String),
					},*//*

				},*/

				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var lookup *resource.Lookup = &resource.Lookup{}
					var q *schema.Query = &schema.Query{}
					q.AppendQuery(p.Args)
					lookup.AddQuery(*q)
					if AuthorController, err := controllers.NewCollection(AUTHOR_TYPE, mongoByte.NewHandler(mongoSession, AuthDatabase, Collection)); err != nil{
						panic(err)
					}else{
						list, err := AuthorController.Find(p.Context, lookup, 0, 50)
						if err != nil{
							fmt.Println(err)
							return nil, err
						}else {
							return list.Models, nil
						}
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
		Types: []graphql.Type{userType, bookType, baseModelType, payloadInfoType},

	})
}




