# AI Service

This microservice handles AI Reasoning (plan generation) using a local LLM (Ollama, DeepSeek-R1).

- **Tech stack:** Go, Gin, REST client for Ollama
- **Architecture:** Clean Architecture, DDD
- **API:**
  - POST `/ai/plan` - Generate a plan for a task

## Structure
- `domain/` - Entities (Plan, Step)
- `usecase/` - Plan generation logic
- `infrastructure/` - Ollama REST client
- `delivery/http/` - Gin HTTP handlers
- `main.go` - Entry point

## Testing
- Table-driven tests for plan generation and handlers

## Security
- Input validation, prompt sanitization, logging

## Documentation
- OpenAPI spec and usage examples
