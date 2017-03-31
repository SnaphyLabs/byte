package main

import (
	"fmt"
	_ "github.com/SnaphyLabs/SnaphyByte/models"
	"net/http"
	"github.com/SnaphyLabs/SnaphyByte/queries"
	"github.com/graphql-go/handler"
)



//Run server here..
func main(){
	fmt.Println("Running server")

	// simplest relay-compliant graphql server HTTP handler
	// using Starwars schema from `graphql-relay-go` examples
	h := handler.New(&handler.Config{
		Schema: &queries.TestSchema,
		Pretty: true,
	})


	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("SnaphyByte/static"))

	http.Handle("/graphql", h)
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)

	/*

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	// and serve!
	http.ListenAndServe(":8080", nil)

	*/

/*
	http.Handle("/graphql", h)
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)*/
}
