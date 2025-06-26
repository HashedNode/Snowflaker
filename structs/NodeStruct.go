package structs

import (
	"sync"
	"time"
)

type Node struct {
	Mutx      sync.Mutex
	Epoch     time.Time
	Time      int64
	Node      int64
	Step      int64
	NodeMax   int64
	NodeMask  int64
	StepMask  int64
	TimeShift uint8
	NodeShift uint8
}

func (node *Node) GenerateSnowflakeId() SnowflakeId {

	defer node.Mutx.Unlock()
	node.Mutx.Lock()

	now := node.timeMillisSinceEpoch()
	if now == node.Time {
		node.Step = (node.Step + 1) & node.StepMask
		if node.Step == 0 {
			for now <= node.Time {
				now = node.timeMillisSinceEpoch()
			}
		}
	} else {
		node.Step = 0
	}

	node.Time = now
	snowflakeId := SnowflakeId((now)<<node.TimeShift | (node.Node << node.NodeShift) | (node.Step))
	return snowflakeId

}

func (node *Node) timeMillisSinceEpoch() int64 {
	return time.Since(node.Epoch).Milliseconds()

}
