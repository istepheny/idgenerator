package fakeIDGenerator

import (
	"github.com/istepheny/idgenerator/contract"
	"github.com/istepheny/idgenerator/drivers"
)

const DriverName = "FakeIDGenerator"

func init() {
	drivers.Register(DriverName, func() contract.IDGenerator { return NewFakeIDGenerator() })
}

type FakeIDGenerator struct{}

func NewFakeIDGenerator() *FakeIDGenerator {
	return &FakeIDGenerator{}
}

func (FakeIDGenerator) ID() int64 {
	return 0
}
