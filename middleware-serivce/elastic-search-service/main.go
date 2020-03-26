package main

import (
	"elastic-search-service/api/rest"
	"elastic-search-service/core/repo"
	"elastic-search-service/middleware/business"
	"elastic-search-service/middleware/storage/mongo"
	"elastic-search-service/middleware/storage/postgres"
	"elastic-search-service/server"
	"fmt"
	env "github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func main() {
	loadEnv()
	repo := getRepository()
	if repo == nil {
		return
	}
	service := business.NewSneakerReferenceService(repo)
	handler := rest.NewHandler(service, os.Getenv("CONTENT_TYPE"))
	routes := rest.ProvideRoutes(handler)
	srv := server.NewInstance(os.Getenv("HOST"))
	srv.SetupRouter(routes)
	srv.Start()
}

func loadEnv() {
	err := env.Load("env/.env")
	if err != nil {
		log.Fatal(err)
	}
}

func httpPort() string {
	port := "8420"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func getRepository() repo.SneakerReferenceRepository {
	switch os.Getenv("USE_DB") {
	case "mongo":
		mongoURL := os.Getenv("MONGO_URL")
		mongodb := os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		mongoCollection := os.Getenv("MONGO_COLLECTION")
		repo, err := mongo.NewMongoRepository(
			mongoURL,
			mongodb,
			mongoCollection,
			mongoTimeout,
		)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "postgres":
		postgresUrl := os.Getenv("POSTGRES_URL")
		postgresTable := os.Getenv("POSTGRES_TABLE")
		repo, err := postgres.NewPostgresRepository(postgresUrl, postgresTable)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}

