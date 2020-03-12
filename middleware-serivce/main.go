package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"middleware-serivce/api"
	"middleware-serivce/model"
	"middleware-serivce/repo/mongo"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	mongoURL = "mongodb://localhost:27017"
	mongoDB = "sneaker-resale-platform"
	mongoCollection = "SneakerProducts"
	mongoTimeout=30
)

func main() {
	repo := getRepo()
	service := model.NewBlogService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/products/sneakers/{code}", handler.Get)
	r.Post("/products/sneakers", handler.Post)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :8420")
		errs <- http.ListenAndServe(httpPort(), r)

	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)

}

func httpPort() string {
	port := "8420"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func getRepo() model.SneakerProductRepository {
	repo, err := mongo.NewMongoRepository(mongoURL, mongoDB, mongoCollection, mongoTimeout)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}

/*
func chooseRepoByEnv() model.SneakerProductRepository {
	switch os.Getenv("URL_DB") {
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
		repo, err := mongo.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
*/
