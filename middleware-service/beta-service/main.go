package main

import "github.com/timoth-y/kicksware-platform/middleware-service/beta-service/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}