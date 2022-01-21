package idgenerator

import (
	"git.ucloudadmin.com/monkey/idgenerator/contract"
	"git.ucloudadmin.com/monkey/idgenerator/drivers"
	"git.ucloudadmin.com/monkey/idgenerator/drivers/int53IDGenerator"
)

func Driver(name string) contract.IDGenerator {
	return drivers.Get(name)
}

func ID() int64 {
	return Driver(int53IDGenerator.DriverName).ID()
}
