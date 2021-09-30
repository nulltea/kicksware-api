package main

import "go.kicksware.com/api/services/orders/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}
