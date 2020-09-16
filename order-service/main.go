package main

import "github.com/timoth-y/kicksware-api/order-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}