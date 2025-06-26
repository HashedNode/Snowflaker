package main

import (
	"crystal_snowflake/generators"
	"crystal_snowflake/services"
	"crystal_snowflake/utils"
	"errors"
	"log"
	"net/http"
	"strconv"
)

func main() {

	utils.InitGlobalVars()

	nodeSize, err := strconv.ParseInt(utils.GetEnvOrDefault("NODE_SIZE", "1"), 10, 64)
	if err != nil {
		log.Fatal("error getting NODE_SIZE")
	}

	serverPort := utils.GetEnvOrDefault("SERVER_PORT", "8080")

	generators.InitSnowflakeNode(nodeSize)

	http.HandleFunc("/generate-id", services.ServeSnowflakeId)

	address := utils.GetEnvOrDefault("SERVER_ADDRESS", "0.0.0.0")
	serverAddress := address + ":" + serverPort

	srv := &http.Server{
		Addr:    serverAddress,
		Handler: nil,
	}

	go func() {
		log.Println("HTTP server started on", serverAddress)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v", err)
		}
	}()
}
