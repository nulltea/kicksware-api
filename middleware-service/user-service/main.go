package main

import "user-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}
