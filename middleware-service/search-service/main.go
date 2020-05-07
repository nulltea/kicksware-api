package main

import "search-service/startup"

func main() {
	srv, container := startup.InitializeServer()
	startup.PerformDataSync(container)
	srv.Start()
}
