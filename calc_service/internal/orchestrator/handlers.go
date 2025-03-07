package orchestrator

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	expressions = make(map[string]*Expression)
	tasks       = make(chan *Task, 10)
	mu          sync.Mutex
)

func HandleCalculate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ExpressionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Expression == "" {
		http.Error(w, "Invalid data", http.StatusUnprocessableEntity)
		return
	}

	id := generateID()
	expr := parseExpression(req.Expression)
	if expr == nil {
		http.Error(w, "Invalid expression", http.StatusUnprocessableEntity)
		return
	}

	result, err := expr.Calculate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	expression := &Expression{
		ID:     id,
		Status: "completed",
		Result: result,
	}

	mu.Lock()
	expressions[id] = expression
	mu.Unlock()
	tasks <- &Task{
		ID:            id,
		Arg1:          result,
		Arg2:          0,
		Operation:     "",
		OperationTime: 0,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func HandleGetExpressions(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	response := struct {
		Expressions []*Expression `json:"expressions"`
	}{Expressions: make([]*Expression, 0, len(expressions))}

	for _, expr := range expressions {
		response.Expressions = append(response.Expressions, expr)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func HandleGetExpressionByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/api/v1/expressions/"):]

	mu.Lock()
	expr, exists := expressions[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]*Expression{"expression": expr})
}

func HandleTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		select {
		case task := <-tasks:
			json.NewEncoder(w).Encode(task)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func HandleResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ResultRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid data", http.StatusUnprocessableEntity)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	if expr, exists := expressions[req.ID]; exists {
		expr.Result = req.Result
		expr.Status = "completed"
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}
