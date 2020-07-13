package main

import (
	"Go_Backend/graph"
	"Go_Backend/graph/generated"
	"Go_Backend/postgres"
	"github.com/go-pg/pg/v10"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

const defaultPort = "5001"

func main() {
	//pgDB := pg.Connect(&pg.Options{
	//	Addr:     "ec2-52-202-146-43.compute-1.amazonaws.com:5432",
	//	User:     "frttnzabknmuam",
	//	Password: "b3cfa42b07816e1b686739f27736a48a0cdbf2918e3dee949bdbc4767093e669",
	//	Database: "d2778uadjv4tq2",
	//
	//})

	opt, err := pg.ParseURL("postgres://frttnzabknmuam:b3cfa42b07816e1b686739f27736a48a0cdbf2918e3dee949bdbc4767093e669@ec2-52-202-146-43.compute-1.amazonaws.com:5432/d2778uadjv4tq2?sslmode=require")
	if err != nil {
		panic(err)
	}

	pgDB := pg.Connect(opt)

	pgDB.AddQueryHook(postgres.DBLogger{})

	defer pgDB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: pgDB}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
