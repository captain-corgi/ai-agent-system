# Backend Services

This directory contains the microservice backends for the AI Agent System.

Services:
- **task-service**: Manages task lifecycle.
- **ai-service**: Generates plans using the local LLM (Ollama).
- **code-execution-service**: Executes code in sandboxed Docker containers.
- **web-browsing-service**: Performs web browsing via chromedp.
- **filesystem-service**: Provides secure filesystem operations.

Each service follows **Clean Architecture** (domain, usecase, infrastructure, delivery/http) and has its own Go module (`go.mod`).

To build & run a service:
```bash
cd <service-name>
go build
./<service-name>
```

Run unit tests:
```bash
go test ./...
```
