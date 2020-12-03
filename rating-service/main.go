package main

import "go.kicksware.com/api/rating-service/startup"

func main() {
	srv := startup.InitializeEventBus()
	srv.SubscribeHandlers()
}