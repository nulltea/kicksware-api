package main

import "go.kicksware.com/api/order-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}