package main

import "go.kicksware.com/api/beta/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}