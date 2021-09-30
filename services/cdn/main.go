package main

import "go.kicksware.com/api/services/cdn/startup"

func main()  {
	srv := startup.InitializeServer()
	srv.Start()
}
