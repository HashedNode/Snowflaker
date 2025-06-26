package utils

import (
	"crystal_snowflake/constants"
	"crystal_snowflake/structs"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"strconv"
)

// InitDecodingAlphabet init constants decoding alphabet, invoke before using functions: ParseBase32ToSnowflakeId, ParseBase58ToSnowflakeId
func InitDecodingAlphabet() {
	for i := 0; i < len(constants.DecodeBase58Alphabet); i++ {
		constants.DecodeBase58Alphabet[i] = 0xFF
	}

	for i := 0; i < len(constants.Base58Alphabet); i++ {
		constants.DecodeBase58Alphabet[constants.Base58Alphabet[i]] = byte(i)
	}

	for i := 0; i < len(constants.DecodedBase32Alphabet); i++ {
		constants.DecodedBase32Alphabet[i] = 0xFF
	}

	for i := 0; i < len(constants.Base32Alphabet); i++ {
		constants.DecodedBase32Alphabet[constants.Base32Alphabet[i]] = byte(i)
	}
}

func ParseInt64ToSnowflakeId(id int64) structs.SnowflakeId {
	return structs.SnowflakeId(id)
}

func ParseStringToSnowflakeId(id string) (structs.SnowflakeId, error) {
	snowId, err := strconv.ParseInt(id, 10, 64)
	return structs.SnowflakeId(snowId), err
}

func ParseStringToSnowflakeIdFromBase2(id string) (structs.SnowflakeId, error) {
	snowId, err := strconv.ParseInt(id, 2, 64)
	return structs.SnowflakeId(snowId), err
}

func ParseBase32ToSnowflakeId(b []byte) (structs.SnowflakeId, error) {

	var id int64

	for i := range b {
		if constants.DecodedBase32Alphabet[b[i]] == 0xFF {
			return -1, errors.New("invalid Base32 character")
		}
		id = id*32 + int64(constants.DecodedBase32Alphabet[b[i]])
	}

	return structs.SnowflakeId(id), nil
}

func ParseBase36ToSnowflakeId(id string) (structs.SnowflakeId, error) {
	snowId, err := strconv.ParseInt(id, 36, 64)
	return structs.SnowflakeId(snowId), err
}

func ParseBase58ToSnowflakeId(b []byte) (structs.SnowflakeId, error) {

	var snowId int64

	for i := range b {
		if constants.DecodeBase58Alphabet[b[i]] == 0xFF {
			return -1, errors.New("invalid Base58 characters")
		}
		snowId = snowId*58 + int64(constants.DecodeBase58Alphabet[b[i]])
	}

	return structs.SnowflakeId(snowId), nil
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
