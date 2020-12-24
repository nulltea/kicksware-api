package events

type Handler interface {
	SubscribeHandlers()
	Listen()
}
