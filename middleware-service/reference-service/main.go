package main

import "reference-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}