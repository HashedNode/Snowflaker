package main

import (
	"crystal_snowflake/utils"
	"fmt"
	"log"
)

func main() {
	node, err := utils.NewSnowflakeNode(2)

	if err != nil {
		log.Fatal("error occurred while generate a node, err is: ", err)
	}

	snowflakeId := node.GenerateSnowflakeId()

	fmt.Printf("Int64  ID: %d\n", snowflakeId)
	fmt.Printf("String ID: %s\n", snowflakeId.ToString())
	fmt.Printf("Base64 ID: %s\n", snowflakeId.ToBase64())

	// Print out the ID's timestamp
	fmt.Printf("ID Time  : %d\n", snowflakeId.Time())

	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", snowflakeId.Node())

	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", snowflakeId.Step())

}
