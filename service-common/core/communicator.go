package core

type InnerCommunicator interface {
	PostMessage(service string, message interface{}, response interface{}) error
	GetMessage(service string, response interface{}) error
}