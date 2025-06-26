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

func (id SnowflakeId) ToBase2() string {
	return strconv.FormatInt(id.ToInt64(), 2)
}

func (id SnowflakeId) ToBase32() string {

	if id < 32 {
		return string(constants.Base32Alphabet[id])
	}

	b := make([]byte, 0, 12)
	for id >= 32 {
		b = append(b, constants.Base32Alphabet[id%32])
		id /= 32
	}
	b = append(b, constants.Base32Alphabet[id])

	for x, y := 0, len(b)-1; x < y; x, y = x+1, y-1 {
		b[x], b[y] = b[y], b[x]
	}

	return string(b)
}

func (id SnowflakeId) ToBase36() string {
	return strconv.FormatInt(int64(id), 36)
}

func (id SnowflakeId) ToBase58() string {

	if id < 58 {
		return string(constants.Base58Alphabet[id])
	}

	b := make([]byte, 0, 11)
	for id >= 58 {
		b = append(b, constants.Base58Alphabet[id%58])
		id /= 58
	}
	b = append(b, constants.Base58Alphabet[id])

	for x, y := 0, len(b)-1; x < y; x, y = x+1, y-1 {
		b[x], b[y] = b[y], b[x]
	}

	return string(b)
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

// MarshalJSON returns a JSON byte array string of the snowflake ID.
func (id SnowflakeId) MarshalJSON() ([]byte, error) {
	buff := make([]byte, 0, 22)
	buff = append(buff, '"')
	buff = strconv.AppendInt(buff, int64(id), 10)
	buff = append(buff, '"')
	return buff, nil
}
