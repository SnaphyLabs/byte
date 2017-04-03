package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/SnaphyLabs/SnaphyByte/models"
)

//Property Having common interface..
var (

	BaseModelInterface = graphql.NewInterface(graphql.InterfaceConfig{
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

	//Interface for payload
	PayloadInterface = graphql.NewInterface(graphql.InterfaceConfig{
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
					if ModelConfig != nil{
						if ModelConfig.Models() != nil{
							if model, ok := ModelConfig.Models()["User"]; ok{
								return model.GraphQLObject()
							}
						}
					}
				}else{
					return BookType
				}
			}
			return nil
		},
	})



)

