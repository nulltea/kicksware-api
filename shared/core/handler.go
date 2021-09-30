package core

type Handler interface {
	Setup()
	Serve() error
}
