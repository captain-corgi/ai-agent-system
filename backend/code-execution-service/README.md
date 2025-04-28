# Code Execution Service

This microservice executes code securely in sandboxed Docker containers as part of the Tool Execution bounded context.

- **Tech stack:** Go, Gin, Docker Go SDK
- **Architecture:** Clean Architecture, DDD
- **API:**
  - POST `/execute/code` - Execute code in a sandbox

## Structure
- `domain/` - Entities (CodeJob)
- `usecase/` - Code execution logic
- `infrastructure/` - Docker client
- `delivery/http/` - Gin HTTP handlers
- `main.go` - Entry point

## Testing
- Table-driven tests for all usecases and handlers

## Security
- Docker sandboxing, FS restriction, input validation

## Documentation
- OpenAPI spec and usage examples
