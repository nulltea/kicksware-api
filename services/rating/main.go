package main

import "go.kicksware.com/api/services/rating/startup"

func main() {
	srv := startup.InitializeEventBus()
	srv.SubscribeHandlers()
}
