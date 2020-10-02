package main

import "go.kicksware.com/api/cdn-service/startup"

func main()  {
	srv := startup.InitializeServer()
	srv.Start()
}