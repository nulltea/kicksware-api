package container

type ServiceContainer interface {
	BindSingleton(factory ServiceFactory) ServiceContainer
	BindTransient(factory ServiceFactory) ServiceContainer
	BindInstance(instance interface{}) ServiceContainer
	Resolve(receiver interface{}) error
	ResolveFor(function ServiceFactory) error
	GetRawFactory(instance interface{}) ServiceFactory
	IsConfigured(instance interface{}) bool
	Reset()
}

type ServiceFactory interface {}

type GenericService interface {}