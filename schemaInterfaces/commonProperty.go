package schemaInterfaces

import (
	"github.com/graphql-go/graphql"
)


//Property Having common interface..
var (
	BaseModelInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name: "CommonProperty",
		Description: "An object with an ID, Created, Updated, ETag",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.ID),
				Description: "The id of the object",
			},
			"created": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"updated": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			/*"type": &graphql.Field{
				Type:	graphql.NewNonNull(graphql.String),
				Description: `Stores the name of collection.
				When Creating multi-tenant model each model will have a type which will tell what typpe of collection does it belongs to.`,
			},*/
		},
		//Implement Resolve type...return a simple grapql.Object as its has a mixed type of resolvers.
		ResolveType: func(p graphql.ResolveTypeParams) (*graphql.Object){
			return &graphql.Object{}
		},
	})

)

