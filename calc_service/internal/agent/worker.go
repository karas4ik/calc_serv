package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

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

var wg sync.WaitGroup

func StartWorker() {
	computingPower := getComputingPower()

	for i := 0; i < computingPower; i++ {
		wg.Add(1)
		go worker()
	}

	wg.Wait()
}

func worker() {
	defer wg.Done()

	for {
		resp, err := http.Get("http://localhost:5000/internal/task")
		if err != nil {
			fmt.Println("Error getting task:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		if resp.StatusCode == http.StatusNotFound {
			time.Sleep(2 * time.Second)
			continue
		}

		var task Task
		json.NewDecoder(resp.Body).Decode(&task)
		resp.Body.Close()

		result := task.Arg1

		resultRequest := ResultRequest{ID: task.ID, Result: result}
		jsonData, _ := json.Marshal(resultRequest)

		http.Post("http://localhost:5000/internal/result", "application/json", bytes.NewBuffer(jsonData))

		fmt.Printf("Processed task %s: Result = %f\n", task.ID, result)
	}
}

func getComputingPower() int {
	if cp, err := strconv.Atoi(os.Getenv("COMPUTING_POWER")); err == nil {
		return cp
	}
	return 1
}

func PerformOperation(task Task) float64 {
	switch task.Operation {
	case "addition":
		return task.Arg1 + task.Arg2
	case "subtraction":
		return task.Arg1 - task.Arg2
	case "multiplication":
		return task.Arg1 * task.Arg2
	case "division":
		if task.Arg2 == 0 {
			return 0
		}
		return task.Arg1 / task.Arg2
	default:
		return 0
	}
}
