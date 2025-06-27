package generators

import (
	"log"
	"snowflaker/constants"
	"snowflaker/structs"
	"strconv"
	"time"
)

var SnowflakeNode *structs.Node

func InitSnowflakeNode(nodeSize int64) {

	if constants.NodeBits+constants.StepBits > 22 {
		log.Fatal("you can share only 22 between Node/Step")
	}

	constants.Mutx.Lock()
	constants.NodeMax = -1 ^ (-1 << constants.NodeBits)
	constants.NodeMask = constants.NodeMax << constants.StepBits
	constants.StepMask = -1 ^ (-1 << constants.StepBits)
	constants.TimeShift = constants.NodeBits + constants.StepBits
	constants.NodeShift = constants.StepBits
	constants.Mutx.Unlock()

	node := structs.Node{}
	node.Node = nodeSize
	node.NodeMax = -1 ^ (-1 << constants.NodeBits)
	node.NodeMask = node.NodeMax << constants.StepBits
	node.StepMask = -1 ^ (-1 << constants.StepBits)
	node.TimeShift = constants.NodeBits + constants.StepBits
	node.NodeShift = constants.StepBits

	if node.Node < 0 || node.Node > node.NodeMax {
		log.Fatal("Node number must be between 0 and " + strconv.FormatInt(node.NodeMax, 10))
	}
	var currentTime = time.Now()
	node.Epoch = currentTime.Add(time.Unix(constants.Epoch/1000, (constants.Epoch%1000)*1000000).Sub(currentTime))

	SnowflakeNode = &node
}
