package orchestrator

type Expression struct {
	ID     string  `json:"id"`
	Status string  `json:"status"`
	Result float64 `json:"result,omitempty"`
}

type ExpressionRequest struct {
	Expression string `json:"expression"`
}

type Task struct {
	ID            string  `json:"id"`
	Arg1          float64 `json:"arg1"`
	Arg2          float64 `json:"arg2"`
	Operation     string  `json:"operation"`
	OperationTime int64   `json:"operation_time"`
}

type ResultRequest struct {
	ID     string  `json:"id"`
	Result float64 `json:"result"`
}
