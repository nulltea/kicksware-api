package main

import (
	"log"

	"go.kicksware.com/api/search-service/startup"
)

func main() {
	srv, container := startup.InitializeServer()
	if err := startup.PerformDataSync(container); err != nil {
		log.Fatalln(err)
	}
	srv.Start()
}
