package container

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"user-service/env"
)

var (
	ErrFactoryNotFunction = errors.New("serviceContainer: the factory must be a function")
	ErrReceiverNotValid = errors.New("serviceContainer: the receiver must be either a reference or a callback function")
	ErrInstanceNotValid = errors.New("serviceContainer: the service instance must be either a reference or an element")
)

type serviceContainer struct {
	factoryMap map[reflect.Type]*serviceBinding
}

type serviceBinding struct {
	factory ServiceFactory
	instance GenericService
	singleton bool
}

func NewServiceContainer() ServiceContainer {
	return &serviceContainer{
		make(map[reflect.Type]*serviceBinding),
	}
}

func (c *serviceContainer) BindSingleton(factory ServiceFactory) ServiceContainer {
	if err := c.bindFactory(factory, true); err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}

func (c *serviceContainer) BindTransient(factory ServiceFactory) ServiceContainer {
	if err := c.bindFactory(factory, false); err != nil {
		log.Fatal(err)
		return nil
	}
	return c
}

func (c *serviceContainer) BindInstance(instance interface{}) ServiceContainer {
	instanceType := reflect.TypeOf(instance); if instanceType == nil {
		log.Fatal(ErrInstanceNotValid)
		return c
	}
	if instanceType.Kind() == reflect.Ptr {
		instanceType = instanceType.Elem()
	}
	c.factoryMap[instanceType] = &serviceBinding{
		factory: func() interface{} {
			return instance
		},
		instance: instance,
		singleton: true,
	}
	return c
}

func (c *serviceContainer) Resolve(receiver interface{}) error {
	receiverType := reflect.TypeOf(receiver); if receiverType == nil {
		return ErrReceiverNotValid
	}

	if receiverType.Kind() == reflect.Ptr {
		receiverElem := receiverType.Elem()

		binding, ok := c.factoryMap[receiverElem]; if !ok {
			return errFServiceNotConfigured(receiverElem.String())
		}

		if instance, err := c.resolveBinding(binding); err == nil {
			reflect.ValueOf(receiver).Elem().Set(reflect.ValueOf(instance))
		} else {
			return err
		}
		return nil
	}

	if receiverType.Kind() == reflect.Func {
		return c.ResolveFor(receiver.(func()interface{}))
	}

	return ErrReceiverNotValid
}

func (c *serviceContainer) ResolveFor(function ServiceFactory) error {
	arguments, err := c.argumentsOf(function); if err != nil {
		return err
	}
	reflect.ValueOf(function).Call(arguments)
	return nil
}

func (c *serviceContainer) GetRawFactory(receiver interface{}) ServiceFactory {
	receiverType := reflect.TypeOf(receiver); if receiverType == nil {
		return nil
	}
	if receiverType.Kind() == reflect.Ptr {
		receiverType = receiverType.Elem()
	}
	return c.factoryMap[receiverType].factory
}

func (c *serviceContainer) IsConfigured(receiver interface{}) bool {
	receiverType := reflect.TypeOf(receiver); if receiverType == nil {
		return false
	}
	if receiverType.Kind() == reflect.Ptr {
		receiverType = receiverType.Elem()
	}
	_, is := c.factoryMap[receiverType]
	return is
}

func (c *serviceContainer) GetConfig() (config env.ServiceConfig) {
	c.Resolve(&config)
	return
}


func (c *serviceContainer) Reset() {
	c.factoryMap = make(map[reflect.Type]*serviceBinding)
}

func (c *serviceContainer) bindFactory(factory ServiceFactory, singleton bool) error {
	factoryType := reflect.TypeOf(factory)
	if factoryType.Kind() != reflect.Func {
		return ErrFactoryNotFunction
	}
	errorInterface := reflect.TypeOf((*error)(nil)).Elem()
	for i := 0; i < factoryType.NumOut(); i++ {
		resolvedType := factoryType.Out(i)
		if resolvedType.Implements(errorInterface) {
			continue
		}
		c.factoryMap[resolvedType] = &serviceBinding{
			factory:  factory,
			instance: nil,
			singleton: singleton,
		}
	}
	return nil
}

func (c *serviceContainer) resolveBinding(binding *serviceBinding) (interface{}, error) {
	if binding.instance != nil {
		return binding.instance, nil
	}
	instance, err := c.invokeFactory(binding.factory); if err != nil {
		return nil, err
	}
	if binding.singleton {
		binding.instance = instance
	}
	return instance, nil
}

func (c *serviceContainer) invokeFactory(factory ServiceFactory) (interface{}, error) {
	args, err := c.argumentsOf(factory); if err != nil {
		return nil, err
	}
	return reflect.ValueOf(factory).Call(args)[0].Interface(), nil
}

func (c *serviceContainer) argumentsOf(factory ServiceFactory) ([]reflect.Value, error) {
	factoryType := reflect.TypeOf(factory)
	argumentsCount := factoryType.NumIn()
	arguments := make([]reflect.Value, argumentsCount)

	for i := 0; i < argumentsCount; i++ {
		argumentType := factoryType.In(i)

		if argumentType.Kind() == reflect.Interface && reflect.TypeOf(c).Implements(argumentType) {
			arguments[i] = reflect.ValueOf(c)
			continue
		} // inject service container argument

		binding, ok := c.factoryMap[argumentType]; if !ok {
			return nil, errFServiceNotConfigured(argumentType.String())
		}

		if instance, err := c.resolveBinding(binding); err == nil {
			arguments[i] = reflect.ValueOf(instance)
		} else {
			return nil, err
		}
	}

	return arguments, nil
}

func errFServiceNotConfigured(service string) error {
	return fmt.Errorf("serviceContainer: service %q resolve factory does not configured", service)
}