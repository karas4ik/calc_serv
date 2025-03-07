package orchestrator

import (
	"os"
	"strconv"

	"github.com/google/uuid"
)

func GetOperationTime(operation string) int64 {
	switch operation {
	case "addition":
		return getEnvInt("TIME_ADDITION_MS", 100)
	case "subtraction":
		return getEnvInt("TIME_SUBTRACTION_MS", 100)
	case "multiplication":
		return getEnvInt("TIME_MULTIPLICATIONS_MS", 100)
	case "division":
		return getEnvInt("TIME_DIVISIONS_MS", 100)
	default:
		return 0
	}
}

func getEnvInt(key string, defaultValue int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func generateID() string {
	return uuid.New().String()
}
