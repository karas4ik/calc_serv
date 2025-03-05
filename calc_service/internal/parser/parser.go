package parser

import (
	"calc_service/internal/models"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func ParseExpression(expr string) []models.Task {
	parts := strings.Fields(expr)
	var tasks []models.Task
	for i := 0; i < len(parts); i++ {
		if len(parts) > i+2 {
			arg1 := parseFloat(parts[i])
			arg2 := parseFloat(parts[i+2])
			operation := parts[i+1]

			tasks = append(tasks, models.Task{
				ID:        uuid.New().String(),
				Arg1:      arg1,
				Arg2:      arg2,
				Operation: operation,
			})
			i += 2
		}
	}
	return tasks
}

func parseFloat(str string) float64 {
	value, _ := strconv.ParseFloat(str, 64)
	return value
}
