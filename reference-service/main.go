package main

import "github.com/timoth-y/kicksware-api/reference-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}