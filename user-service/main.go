package main

import "github.com/timoth-y/kicksware-api/user-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}
