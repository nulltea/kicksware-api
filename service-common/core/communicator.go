package core

type InnerCommunicator interface {
	PostMessage(endpoint string, message interface{}, response interface{}) error
	GetMessage(endpoint string, response interface{}) error
}