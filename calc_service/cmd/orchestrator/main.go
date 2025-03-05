package main

import (
	"calc_service/internal/server"
	"log"
)

func main() {
	log.Println("Starting orchestrator...")
	server.StartOrchestrator()
}
