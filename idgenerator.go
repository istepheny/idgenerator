package idgenerator

import (
	"github.com/istepheny/idgenerator/contract"
	"github.com/istepheny/idgenerator/drivers"
	"github.com/istepheny/idgenerator/drivers/int53IDGenerator"
)

func Driver(name string) contract.IDGenerator {
	return drivers.Get(name)
}

func ID() int64 {
	return Driver(int53IDGenerator.DriverName).ID()
}
