package drivers

import (
	"sync"

	"git.ucloudadmin.com/monkey/idgenerator/contract"
)

type IDGeneratorFactory func() contract.IDGenerator

var (
	mu        sync.RWMutex
	factories = make(map[string]IDGeneratorFactory)
)

func Register(name string, factory IDGeneratorFactory) {
	mu.Lock()
	defer mu.Unlock()
	if factory == nil {
		panic("register idGenerator is nil")
	}

	if _, dup := factories[name]; dup {
		panic("register called twice for idGenerator " + name)
	}

	factories[name] = factory
}

func Get(name string) contract.IDGenerator {
	mu.RLock()
	defer mu.RUnlock()
	factory, ok := factories[name]
	if !ok {
		panic("idGenerator " + name + " is not available")
	}
	return factory()
}
