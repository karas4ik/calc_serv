package main

import (
	"calc_service/internal/workers"
	"log"
)

func main() {
	log.Println("Starting agent...")
	workers.StartWorkers()
}
