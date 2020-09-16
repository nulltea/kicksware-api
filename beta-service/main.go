package main

import "github.com/timoth-y/kicksware-api/beta-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}