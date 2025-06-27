package utils

import (
	"log"
	"os"
	"snowflaker/constants"
	"strconv"
)

func GetEnvOrDefault(keyEnv string, defaultValue string) string {
	value, exists := os.LookupEnv(keyEnv)
	if !exists {
		return defaultValue
	}
	return value
}

func InitGlobalVars() {
	nodeBits, err := strconv.ParseInt(GetEnvOrDefault("NODE_BITS", "10"), 0, 8)
	stepBits, err := strconv.ParseInt(GetEnvOrDefault("STEP_BITS", "12"), 0, 8)

	if err != nil {
		log.Fatal("Unable to parse NODE_BITS")
	}

	constants.Epoch = 1288834974657
	constants.NodeBits = uint8(nodeBits)
	constants.StepBits = uint8(stepBits)
	constants.NodeMax = -1 ^ (-1 << constants.NodeBits)
	constants.NodeMask = constants.NodeMax << constants.StepBits
	constants.StepMask = -1 ^ (-1 << constants.StepBits)
	constants.TimeShift = constants.NodeBits + constants.StepBits
	constants.NodeShift = constants.StepBits
}
