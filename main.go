package main

import (
	"context"
	"log"
	"mongo-graph/graph"
	"mongo-graph/graph/generated"
	"mongo-graph/services/user"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost"))
	if err != nil {
		log.Fatal("Unable to connect mongo db")
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Unable to ping mongo db")
	}

	// User Init
	userService := user.NewUserService(client.Database("bagidu"))

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
			UserService: userService,
		}}),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	// http.Handle("/graphql", handler)
	http.ListenAndServe(":8888", nil)

}
