package int53IDGenerator

import (
	"sync"

	bwmarrinSnowflake "github.com/bwmarrin/snowflake"
	"github.com/istepheny/idgenerator/contract"
	"github.com/istepheny/idgenerator/drivers"
	"github.com/istepheny/idgenerator/snowflake"
)

const DriverName = "Int53IDGenerator"

func init() {
	drivers.Register(DriverName, func() contract.IDGenerator { return NewInt53IDGenerator() })
}

var (
	once             = sync.Once{}
	int53IDGenerator *Int53IDGenerator
)

type Int53IDGenerator struct {
	node *bwmarrinSnowflake.Node
}

func NewInt53IDGenerator() *Int53IDGenerator {
	once.Do(func() {
		o := snowflake.Options{
			NodeBits: 8,
			StepBits: 4,
		}

		int53IDGenerator = &Int53IDGenerator{
			node: snowflake.New(o),
		}
	})

	return int53IDGenerator
}

func (g *Int53IDGenerator) ID() int64 {
	return g.node.Generate().Int64()
}
