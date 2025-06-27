package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"snowflaker/generators"
	"snowflaker/services"
	"snowflaker/utils"
	"strconv"
	"syscall"
	"time"
)

func main() {

	utils.InitGlobalVars()

	nodeSize, err := strconv.ParseInt(utils.GetEnvOrDefault("NODE_SIZE", "1"), 10, 64)
	if err != nil {
		log.Fatal("error getting NODE_SIZE")
	}

	serverPort := utils.GetEnvOrDefault("SERVER_PORT", "8080")

	generators.InitSnowflakeNode(nodeSize)
	muxServer := http.NewServeMux()

	muxServer.HandleFunc("/generate-id", services.ServeSnowflakeId)

	address := utils.GetEnvOrDefault("SERVER_ADDRESS", "0.0.0.0")
	serverAddress := address + ":" + serverPort

	srv := &http.Server{
		Addr:    serverAddress,
		Handler: muxServer,
	}

	closeServer := make(chan os.Signal, 1)
	signal.Notify(closeServer, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("HTTP server started on", serverAddress)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-closeServer
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server closed")
}
