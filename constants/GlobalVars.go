package constants

import (
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
