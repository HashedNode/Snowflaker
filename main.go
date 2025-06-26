package main

import (
	"crystal_snowflake/generators"
	"crystal_snowflake/utils"
	"log"
	"strconv"
)

func main() {

	utils.InitGlobalVars()

	nodeSize, err := strconv.ParseInt(utils.GetEnvOrDefault("NODE_SIZE", "10"), 10, 64)
	if err != nil {
		log.Fatal("error getting NODE_SIZE")
	}

	generators.InitSnowflakeNode(nodeSize)

	snowflakeId := generators.SnowflakeNode.GenerateSnowflakeId()

	log.Printf("Int64  ID: %d\n", snowflakeId)
	log.Printf("String ID: %s\n", snowflakeId.ToString())
	log.Printf("Base64 ID: %s\n", snowflakeId.ToBase64())

	// Print out the ID's timestamp
	log.Printf("ID Time  : %d\n", snowflakeId.Time())

	// Print out the ID's node number
	log.Printf("ID Node  : %d\n", snowflakeId.Node())

	// Print out the ID's sequence number
	log.Printf("ID Step  : %d\n", snowflakeId.Step())

}
