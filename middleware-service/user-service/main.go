package main

import (
	"log"
	"os"
	"strconv"

	env "github.com/joho/godotenv"

	"user-service/api/rest"
	"user-service/core/repo"
	"user-service/middleware/business"
	"user-service/middleware/storage/mongo"
	"user-service/middleware/storage/postgres"
	"user-service/middleware/storage/redis"
	"user-service/server"
)

func main() {
	if os.Getenv("DEBUG") == "True"{
		loadEnv()
	}
	repo := getRepository()
	if repo == nil {
		return
	}
	service := business.NewUserService(repo)
	expirationDelta, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_DELTA")); if err != nil {
		panic(err)
	}
	auth := business.NewAuthServiceJWT(service, expirationDelta)
	handler := rest.NewHandler(service, auth, os.Getenv("CONTENT_TYPE"))
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

func getRepository() repo.UserRepository {
	switch os.Getenv("USE_DB") {
	case "redis":
		redisURL := os.Getenv("REDIS_URL")
		repo, err := redis.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo
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

