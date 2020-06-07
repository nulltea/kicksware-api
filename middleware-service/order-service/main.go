package main

import "order-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}