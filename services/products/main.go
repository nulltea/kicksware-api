package main

import "go.kicksware.com/api/services/products/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}
