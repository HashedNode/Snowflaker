package constants

import "sync"

var (
	Epoch     int64 = 1288834974657
	NodeBits  uint8 = 10
	StepBits  uint8 = 12
	Mutx      sync.Mutex
	NodeMax   int64 = -1 ^ (-1 << NodeBits)
	NodeMask        = NodeMax << StepBits
	StepMask  int64 = -1 ^ (-1 << StepBits)
	TimeShift       = NodeBits + StepBits
	NodeShift       = StepBits
)
