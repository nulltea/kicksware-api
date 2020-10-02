package main

import "go.kicksware.com/api/user-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}
