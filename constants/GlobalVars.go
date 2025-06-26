package constants

import (
	"log"
	"os"
	"strconv"
	"sync"
)

var (
	Epoch     int64
	NodeBits  uint8
	StepBits  uint8
	Mutx      sync.Mutex
	NodeMax   int64
	NodeMask  int64
	StepMask  int64
	TimeShift uint8
	NodeShift uint8
)

func InitGlobalVars() {
	//nodebits := os.Getenv("NODE_BITS")
	nodeBits, err := strconv.ParseInt(os.Getenv("NODE_BITS"), 0, 8)
	stepBits, err := strconv.ParseInt(os.Getenv("STEP_BITS"), 0, 8)

	if err != nil {
		log.Fatal("Unable to parse NODE_BITS")
	}

	Epoch = 1288834974657
	NodeBits = uint8(nodeBits)
	StepBits = uint8(stepBits)
	NodeMax = -1 ^ (-1 << NodeBits)
	NodeMask = NodeMax << StepBits
	StepMask = -1 ^ (-1 << StepBits)
	TimeShift = NodeBits + StepBits
	NodeShift = StepBits
}
