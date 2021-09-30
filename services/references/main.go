package main

import "go.kicksware.com/api/services/references/startup"

func main() {
	srv := startup.InitializeServer()
	srv.Start()
}
