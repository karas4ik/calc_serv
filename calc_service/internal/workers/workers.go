package workers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Task struct {
	ID        string  `json:"id"`
	Arg1      float64 `json:"arg1"`
	Arg2      float64 `json:"arg2"`
	Operation string  `json:"operation"`
}

func StartWorkers() {
	computingPower := getComputingPower()
	for i := 0; i < computingPower; i++ {
		go worker(i + 1)
	}
}

func getComputingPower() int {
	value := os.Getenv("COMPUTING_POWER")
	if value == "" {
		return 4
	}
	power, err := strconv.Atoi(value)
	if err != nil {
		return 4
	}
	return power
}

func worker(workerID int) {
	log.Printf("Worker %d is starting...", workerID)
	for {
		task := getTask()
		if task != nil {
			result := performCalculation(task)
			submitResult(task.ID, result)
		}
		time.Sleep(1 * time.Second)
	}
}

func getTask() *Task {
	resp, err := http.Get("http://localhost:8080/internal/task")
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil
	}
	var task Task
	json.NewDecoder(resp.Body).Decode(&task)
	return &task
}

func performCalculation(task *Task) float64 {
	sleepDuration := getOperationSleepTime(task.Operation)
	time.Sleep(sleepDuration)
	switch task.Operation {
	case "add":
		return task.Arg1 + task.Arg2
	case "subtract":
		return task.Arg1 - task.Arg2
	case "multiply":
		return task.Arg1 * task.Arg2
	case "divide":
		return task.Arg1 / task.Arg2
	default:
		return 0
	}
}

func getOperationSleepTime(operation string) time.Duration {
	switch operation {
	case "add":
		return time.Duration(getEnvAsInt("TIME_ADDITION_MS", 1000)) * time.Millisecond
	case "subtract":
		return time.Duration(getEnvAsInt("TIME_SUBTRACTION_MS", 1000)) * time.Millisecond
	case "multiply":
		return time.Duration(getEnvAsInt("TIME_MULTIPLICATIONS_MS", 1000)) * time.Millisecond
	case "divide":
		return time.Duration(getEnvAsInt("TIME_DIVISIONS_MS", 1000)) * time.Millisecond
	default:
		return 0
	}
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func submitResult(id string, result float64) {
	data := map[string]interface{}{"id": id, "result": result}
	body, _ := json.Marshal(data)
	http.Post("http://localhost:8080/internal/result", "application/json", bytes.NewBuffer(body))
}
