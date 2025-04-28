# Task Service

This microservice manages the lifecycle of tasks (submission, status, results) as the Task Management bounded context.

- **Tech stack:** Go, Gin, SQLite
- **Architecture:** Clean Architecture, DDD
- **APIs:**
  - POST `/tasks` - Submit a new task
  - GET `/tasks/:id` - Get task status/result
  - GET `/tasks` - List all tasks

## Structure
- `domain/` - Entities (Task), Value Objects
- `usecase/` - Business logic (CreateTask, ListTasks, etc.)
- `infrastructure/` - SQLite repository implementations
- `delivery/http/` - Gin HTTP handlers
- `main.go` - Entry point

## Testing
- Table-driven tests for all usecases and handlers

## Security
- Input validation, parameterized queries, no direct SQL in handlers

## Documentation
- OpenAPI spec and usage examples
