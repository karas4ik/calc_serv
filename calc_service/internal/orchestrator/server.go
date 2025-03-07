package orchestrator

import "net/http"

func StartServer() {
	http.HandleFunc("/api/v1/calculate", HandleCalculate)
	http.HandleFunc("/api/v1/expressions", HandleGetExpressions)
	http.HandleFunc("/api/v1/expressions/", HandleGetExpressionByID)
	http.HandleFunc("/internal/task", HandleTask)
	http.HandleFunc("/internal/result", HandleResult)

	http.ListenAndServe(":5000", nil)
}
