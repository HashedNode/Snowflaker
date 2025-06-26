package main

import (
	"crystal_snowflake/constants"
	"crystal_snowflake/utils"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	constants.InitGlobalVars()

	nodeSize, err := strconv.ParseInt(os.Getenv("NODE_SIZE"), 10, 64)
	if err != nil {
		log.Fatal("error getting NODE_SIZE")
	}

	node, err := utils.NewSnowflakeNode(nodeSize)

	if err != nil {
		log.Fatal("error occurred while generate a node, err is: ", err)
	}

	snowflakeId := node.GenerateSnowflakeId()

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
