package main

import "go.kicksware.com/api/beta-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}