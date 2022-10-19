package remoteIDGenerator

import (
	"github.com/istepheny/idgenerator/contract"
	"github.com/istepheny/idgenerator/drivers"
)

const DriverName = "RemoteIDGenerator"

func init() {
	drivers.Register(DriverName, func() contract.IDGenerator { return NewRemoteIDGenerator() })
}

type RemoteIDGenerator struct{}

func NewRemoteIDGenerator() *RemoteIDGenerator {
	return &RemoteIDGenerator{}
}

func (RemoteIDGenerator) ID() int64 {
	panic("implement me")
}
