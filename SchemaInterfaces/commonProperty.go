package SchemaInterfaces

import (
	"github.com/graphql-go/graphql"
)


//Property Having common interface..
var (
	CommonPropertyInterface = graphql.NewInterface(graphql.InterfaceConfig{
		Name: "CommonProperty",
		Description: "An object with an ID, Created, Updated, ETag",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.ID),
				Description: "The id of the object",
			},
			"Created": &graphql.Field{
				Type: graphql.String,
			},
			"Updated": &graphql.Field{
				Type: graphql.Boolean,
			},
		},
		//Implement Resolve type...
		/*ResolveType: func(p graphql.ResolveTypeParams) (*graphql.Object){

		},*/
	})
)

