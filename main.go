package main

import (
	"crystal_snowflake/generators"
	"crystal_snowflake/services"
	"crystal_snowflake/utils"
	"log"
	"net/http"
	"strconv"
)

func main() {

	utils.InitGlobalVars()

	nodeSize, err := strconv.ParseInt(utils.GetEnvOrDefault("NODE_SIZE", "10"), 10, 64)
	if err != nil {
		log.Fatal("error getting NODE_SIZE")
	}

	serverPort := utils.GetEnvOrDefault("SERVER_PORT", "8080")

	generators.InitSnowflakeNode(nodeSize)

	http.HandleFunc("/generate-id", services.ServeSnowflakeId)

	address := utils.GetEnvOrDefault("SERVER_ADDRESS", "0.0.0.0")
	serverAddress := address + ":" + serverPort

	err = http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatal("error starting http server, err is", err)
	}
	log.Println("http server started on address", serverAddress)
}
