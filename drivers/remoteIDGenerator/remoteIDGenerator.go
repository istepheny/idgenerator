package remoteIDGenerator

import (
	"git.ucloudadmin.com/monkey/idgenerator/contract"
	"git.ucloudadmin.com/monkey/idgenerator/drivers"
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
