package main

import (
	"fmt"
	"gqlServerType2"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers(), graphql.MaxParallelism(20)}
	schema := graphql.MustParseSchema(gqlServerType2.Schema, &gqlServerType2.Resolver{}, opts...)

	http.Handle("/graphql", &relay.Handler{Schema: schema})
	fmt.Println("Now Server is running on 'http://localhost:8080/graphql'")
	http.ListenAndServe(":8080", nil)
}
