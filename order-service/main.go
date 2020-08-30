package main

import "github.com/timoth-y/kicksware-platform/middleware-service/order-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}