package test

import (
	"bytes"
	"calc_service/internal/orchestrator"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleCalculate(t *testing.T) {
	reqBody := `{"expression": "4 + 4 - 4 * 4"}`
	req, err := http.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(orchestrator.HandleCalculate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)
	if _, exists := response["id"]; !exists {
		t.Errorf("Expected ID in response, got %v", response)
	}
}

func TestHandleGetExpressions(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/api/v1/expressions", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(orchestrator.HandleGetExpressions)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestHandleGetExpressionByID(t *testing.T) {
	reqBody := `{"expression": "4 + 4 - 4 * 4"}`
	req, err := http.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(orchestrator.HandleCalculate)

	handler.ServeHTTP(rr, req)

	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)
	id := response["id"]

	req, err = http.NewRequest(http.MethodGet, "/api/v1/expressions/"+id, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(orchestrator.HandleGetExpressionByID)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var exprResponse map[string]*orchestrator.Expression
	json.NewDecoder(rr.Body).Decode(&exprResponse)

	if expr, exists := exprResponse["expression"]; !exists || expr.ID != id {
		t.Errorf("Expected expression ID %s, got %v", id, expr)
	}
}
