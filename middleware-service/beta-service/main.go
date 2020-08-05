package main

import "beta-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}