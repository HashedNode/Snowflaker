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

const Base32Alphabet = "ybndrfg8ejkmcpqxot1uwisza345h769" //z-base-32 alphabet alternatives to Crockford's Base32 alphabet 0123456789ABCDEFGHJKMNPQRSTVWXYZ
var DecodedBase32Alphabet [256]byte

const Base58Alphabet = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var DecodeBase58Alphabet [256]byte
