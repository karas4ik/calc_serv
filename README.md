
|----------------------------------|-------------------------------------------------------|
| Submit Arithmetic Expression      | Users can send expressions for computation.          |
| Get Expression Status             | Users can query the status of their submitted expressions. |
| Retrieve Results                  | Users can obtain the results of their calculations.   |
| Parallel Processing               | Agents perform calculations concurrently for efficiency. |

### Statuses

| Status                           | Description                                         |
|----------------------------------|-----------------------------------------------------|
| pending                      | Expression is received and being processed.        |
| completed                    | Calculation is done and result is available.       |
| not found                    | The requested expression or task does not exist.    |
| invalid                      | The submitted expression is not valid.              |

## Example Requests and Responses

### 1. Submitting a Calculation

Request:

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "5 + 3 * 2"
}'
Response:

{
    "id": "unique-expression-id"
}
### 2. Retrieving All Expressions

Request:

curl --location 'http://localhost:8080/api/v1/expressions'
Response:

{
    "expressions": [
        {
            "id": "unique-expression-id",
            "status": "pending",
            "result": null
        }
    ]
}
### 3. Getting Expression by ID

Request:

curl --location 'http://localhost:8080/api/v1/expressions/unique-expression-id'
Response:

{
    "id": "unique-expression-id",
    "status": "completed",
    "result": 11
}
### 4. Retrieving a Task for Processing

Request:

curl --location 'http://localhost:8080/internal/task'
Response:

{
    "id": "task-id",
    "arg1": 5,
    "arg2": 3,
    "operation": "add"
}
### 5. Sending Back Calculation Result

Request:

curl --location 'http://localhost:8080/internal/result' \
--header 'Content-Type: application/json' \
--data '{
  "id": "task-id",
  "result": 8
}'
Response: (HTTP 200 OK)

### Error Codes and Responses

| Status Code | Description                                       |
|-------------|---------------------------------------------------|
| 200     | Successfully processed the request                |
| 404     | The requested expression or task was not found    |
| 422     | Invalid data submitted                             |
| 500     | An internal server error occurred                  |

## Detailed Flow of Operations

### Step-by-Step Example

1. User Submits Expression:
   - A user submits the expression 5 + 3 * 2.
   - The orchestrator receives this input and assigns it an ID.

2. Task Creation:
   - The expression is parsed into tasks. In this case, it breaks down to:
     - Task 1: 5 + 3
     - Task 2: 3 * 2

3. Task Dispatching:
   - The tasks are queued in the task queue.
   - The orchestrator dispatches these tasks to available agents.

4. Agent Processing:
   - Each agent picks up a task, performs the calculation, and returns the result to the orchestrator.

5. Result Retrieval:
   - The orchestrator updates the status of the expression to completed and stores the results.

6. User Queries Status:
   - The user can query the status and retrieve the final computed result.

### Status Flow Example

| Expression ID          | Status        | Result |
|------------------------|---------------|--------|
| unique-expression-id    | pending       | null   |
| unique-expression-id    | completed     | 11     |

### Concurrency and Performance

- The system utilizes goroutines to enable concurrent processing of tasks. This allows multiple calculations to occur simultaneously, significantly improving performance for large sets of expressions.

### Conclusion

The Distributed Arithmetic Expression Calculator is a powerful tool that efficiently processes arithmetic expressions using a distributed architecture. With its robust design, users can easily submit expressions and receive quick results, making it suitable for various applications requiring mathematical computations.

