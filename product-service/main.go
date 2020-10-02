package main

import "go.kicksware.com/api/product-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}