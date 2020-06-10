package main

import "cdn-service/startup"

func main()  {
	srv := startup.InitializeServer()
	srv.Start()
}