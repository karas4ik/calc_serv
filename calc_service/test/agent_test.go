package test

import (
	"bytes"
	"calc_service/internal/agent"
	"calc_service/internal/orchestrator"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPerformOperation(t *testing.T) {
	task := agent.Task{Arg1: 4, Arg2: 4, Operation: "+"}
	result := agent.PerformOperation(task)

	if result != 8 {
		t.Errorf("expected 8, got %f", result)
	}

	task.Operation = "-"
	result = agent.PerformOperation(task)

	if result != 0 {
		t.Errorf("expected 0, got %f", result)
	}

	task.Operation = "*"
	task.Arg2 = 4
	result = agent.PerformOperation(task)

	if result != 16 {
		t.Errorf("expected 16, got %f", result)
	}

	task.Operation = "/"
	task.Arg2 = 4
	result = agent.PerformOperation(task)

	if result != 1 {
		t.Errorf("expected 1, got %f", result)
	}
}

func TestHandleResult(t *testing.T) {
	reqBody := `{"id": "1", "result": 8}`
	req, err := http.NewRequest(http.MethodPost, "/internal/result", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(orchestrator.HandleResult)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
