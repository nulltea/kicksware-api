package main

import "go.kicksware.com/api/reference-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}