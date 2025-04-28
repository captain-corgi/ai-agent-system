# Task Service

This microservice manages the lifecycle of tasks (submission, status, results) as the Task Management bounded context.

- **Tech stack:** Go, Gin, SQLite
- **Architecture:** Clean Architecture, DDD
- **APIs:**
  - `POST /tasks` - Submit a new task
    - **Request:** `{ "type": "code|browse|fs", "payload": { ... } }`
    - **Response:** `{ "id": "<task_id>", "status": "pending" }`
    - **Errors:** `400 Bad Request`, `500 Internal Server Error`
  - `GET /tasks/:id` - Get task status/result
    - **Response:** `{ "id": "<task_id>", "status": "pending|running|done|failed", "result": { ... } }`
    - **Errors:** `404 Not Found`, `500 Internal Server Error`
  - `GET /tasks` - List all tasks
    - **Response:** `[ { "id": "<task_id>", "status": "..." }, ... ]`
  - `GET /tasks/health` - Health check endpoint
    - **Response:** `{ "status": "ok" }`

## Structure Diagram
```mermaid
graph TD
  TaskDomain[domain/]
  TaskUsecase[usecase/]
  TaskInfra[infrastructure/]
  TaskDeliv[delivery/http/]
  Main[main.go]
  TaskDomain-->TaskUsecase
  TaskUsecase-->TaskInfra
  TaskInfra-->TaskDeliv
  TaskDeliv-->Main
```

## Features
- Task submission, status, and result retrieval
- Task listing and filtering
- Full lifecycle management

## Data Flow Diagram (DFD)
```mermaid
graph TD
  FE[Frontend]-->|POST /tasks|TaskSvc[Task Service]
  TaskSvc-->|Store|SQLite[SQLite]
  TaskSvc-->|Status|FE
```

## Entity Relationship Diagram (ERD)
```mermaid
erDiagram
  TASK ||--o{ STEP : contains
  TASK {
    string id PK
    string type
    string status
    string payload
    string result
    datetime created_at
    datetime updated_at
  }
  STEP {
    string id PK
    string task_id FK
    string type
    string input
    string output
    int sequence
    string status
  }
```

## Database Table
| Field       | Type      | PK | FK | Description                |
|-------------|-----------|----|----|----------------------------|
| id          | TEXT      | Y  |    | Task unique identifier     |
| type        | TEXT      |    |    | Task type (code, browse)   |
| status      | TEXT      |    |    | Task status                |
| payload     | TEXT      |    |    | Task input payload         |
| result      | TEXT      |    |    | Task result (JSON)         |
| created_at  | DATETIME  |    |    | Creation timestamp         |
| updated_at  | DATETIME  |    |    | Last update timestamp      |

## Testing
- Table-driven tests for all usecases and handlers

## Security
- Input validation, parameterized queries, no direct SQL in handlers

## Documentation
- OpenAPI spec and usage examples
