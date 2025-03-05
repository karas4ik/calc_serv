package orchestrator

import (
	"calc_service/internal/models"
	"calc_service/internal/parser"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var expressions = make(map[string]models.Expression)
var mu sync.Mutex
var tasks = make(chan models.Task, 10)

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Expression string `json:"expression"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Expression == "" {
		http.Error(w, "invalid data", http.StatusUnprocessableEntity)
		return
	}
	id := uuid.New().String()
	mu.Lock()
	expressions[id] = models.Expression{ID: id, Status: "pending"}
	mu.Unlock()

	taskList := parser.ParseExpression(req.Expression)
	for _, task := range taskList {
		tasks <- task
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func getExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	var list []models.Expression
	for _, expr := range expressions {
		list = append(list, expr)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"expressions": list})
}

func getExpressionByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	mu.Lock()
	defer mu.Unlock()
	expression, exists := expressions[id]
	if !exists {
		http.Error(w, "expression not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(expression)
}

func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	select {
	case task := <-tasks:
		json.NewEncoder(w).Encode(task)
	default:
		http.Error(w, "no task available", http.StatusNotFound)
	}
}

func receiveResultHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID     string  `json:"id"`
		Result float64 `json:"result"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid data", http.StatusUnprocessableEntity)
		return
	}
	mu.Lock()
	expression, exists := expressions[req.ID]
	if !exists {
		http.Error(w, "task not found", http.StatusNotFound)
		mu.Unlock()
		return
	}
	expression.Status = "completed"
	expression.Result = req.Result
	expressions[req.ID] = expression
	mu.Unlock()
	w.WriteHeader(http.StatusOK)
}
