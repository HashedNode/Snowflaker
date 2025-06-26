package structs

import (
	"crystal_snowflake/constants"
	"encoding/base64"
	"encoding/binary"
	"strconv"
)

type SnowflakeId int64

func (id SnowflakeId) ToInt64() int64 {
	return int64(id)
}

func (id SnowflakeId) ToString() string {
	return strconv.FormatInt(id.ToInt64(), 10)
}

func (id SnowflakeId) ToBytes() []byte {
	return []byte(id.ToString())
}

func (id SnowflakeId) ToBase64() string {
	return base64.StdEncoding.EncodeToString(id.ToBytes())
}

// ToIntBytes that returns a byte array of big indian integer
func (id SnowflakeId) ToIntBytes() [8]byte {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(id))
	return b
}

func (id SnowflakeId) Time() int64 {
	return (int64(id) >> constants.TimeShift) + constants.Epoch
}

// Node returns an int64 of the snowflake ID node number
func (id SnowflakeId) Node() int64 {
	return int64(id) & constants.NodeMask >> constants.NodeShift
}

// Step returns an int64 of the snowflake step (or sequence) number
func (id SnowflakeId) Step() int64 {
	return int64(id) & constants.StepMask
}
