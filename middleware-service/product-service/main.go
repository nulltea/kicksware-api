package main

import "product-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}