package main

import "go.kicksware.com/api/services/users/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}
