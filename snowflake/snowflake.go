package snowflake

import (
	"fmt"

	"git.ucloudadmin.com/monkey/idgenerator/nodeid"
	"github.com/bwmarrin/snowflake"
)

const epoch = 1640966400000 // 2022-01-01 00:00:00

type Options struct {
	Epoch    int64
	NodeBits uint8
	StepBits uint8
	NodeId   int64
}

func New(o Options) *snowflake.Node {
	if o.Epoch != 0 {
		snowflake.Epoch = o.Epoch
	} else {
		snowflake.Epoch = epoch
	}

	if o.NodeBits != 0 {
		snowflake.NodeBits = o.NodeBits
	}

	if o.StepBits != 0 {
		snowflake.StepBits = o.StepBits
	}

	if o.NodeId == 0 {
		nodeId, err := nodeid.Lower8BitIP()
		if err != nil {
			panic(fmt.Errorf("snowflake: failed to create generator, %v", err))
		}
		o.NodeId = int64(nodeId)
	}

	node, err := snowflake.NewNode(o.NodeId)
	if err != nil {
		panic(fmt.Errorf("snowflake: failed to create generator, %v", err))
	}

	return node
}
