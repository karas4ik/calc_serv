package server

import (
	"calc_service/internal/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartOrchestrator() {
	router := mux.NewRouter()
	setupRoutes(router)
	log.Println("Orchestrator is starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func setupRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/calculate", controllers.CalculateHandler).Methods("POST")
	router.HandleFunc("/api/v1/expressions", controllers.GetExpressionsHandler).Methods("GET")
	router.HandleFunc("/api/v1/expressions/{id}", controllers.GetExpressionByIDHandler).Methods("GET")
	router.HandleFunc("/internal/task", controllers.GetTaskHandler).Methods("GET")
	router.HandleFunc("/internal/result", controllers.ReceiveResultHandler).Methods("POST")
}
