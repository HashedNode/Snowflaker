package utils

import (
	"encoding/base64"
	"encoding/binary"
	"snowflaker/structs"
	"strconv"
)

func ParseInt64ToSnowflakeId(id int64) structs.SnowflakeId {
	return structs.SnowflakeId(id)
}

func ParseStringToSnowflakeId(id string) (structs.SnowflakeId, error) {
	snowId, err := strconv.ParseInt(id, 10, 64)
	return structs.SnowflakeId(snowId), err
}

func ParseBytesToSnowflakeId(id []byte) (structs.SnowflakeId, error) {
	i, err := strconv.ParseInt(string(id), 10, 64)
	return structs.SnowflakeId(i), err
}

func ParseBase64ToSnowflakeId(id string) (structs.SnowflakeId, error) {
	b, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return -1, err
	}
	return ParseBytesToSnowflakeId(b)

}

// ParseIntBytes converts an array of bytes encoded as big endian integer to SnowflakeId
func ParseIntBytes(id [8]byte) structs.SnowflakeId {
	return structs.SnowflakeId(int64(binary.BigEndian.Uint64(id[:])))
}
