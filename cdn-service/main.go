package main

import "github.com/timoth-y/kicksware-api/cdn-service/startup"

func main()  {
	srv := startup.InitializeServer()
	srv.Start()
}