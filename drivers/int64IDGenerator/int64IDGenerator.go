package int64IDGenerator

import (
	"sync"

	"git.ucloudadmin.com/monkey/idgenerator/contract"
	"git.ucloudadmin.com/monkey/idgenerator/drivers"
	"git.ucloudadmin.com/monkey/idgenerator/snowflake"
	bwmarrinSnowflake "github.com/bwmarrin/snowflake"
)

const DriverName = "Int64IDGenerator"

func init() {
	drivers.Register(DriverName, func() contract.IDGenerator { return NewInt64IDGenerator() })
}

var (
	once             = sync.Once{}
	int64IDGenerator *Int64IDGenerator
)

type Int64IDGenerator struct {
	node *bwmarrinSnowflake.Node
}

func NewInt64IDGenerator() *Int64IDGenerator {
	once.Do(func() {
		o := snowflake.Options{
			NodeBits: 10,
			StepBits: 12,
		}

		int64IDGenerator = &Int64IDGenerator{
			node: snowflake.New(o),
		}
	})

	return int64IDGenerator
}

func (g *Int64IDGenerator) ID() int64 {
	return g.node.Generate().Int64()
}
