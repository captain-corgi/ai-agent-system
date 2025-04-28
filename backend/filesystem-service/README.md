# Filesystem Service

This microservice performs secure file operations (read, write, delete) restricted to `/app/workspace` as part of the Tool Execution bounded context.

- **Tech stack:** Go, Gin, os package
- **Architecture:** Clean Architecture, DDD
- **API:**
  - GET `/fs/read` - Read file
  - POST `/fs/write` - Write file
  - DELETE `/fs/delete` - Delete file

## Structure
- `domain/` - Entities (FsOperation)
- `usecase/` - File operation logic
- `infrastructure/` - Path sanitizer, OS access
- `delivery/http/` - Gin HTTP handlers
- `main.go` - Entry point

## Testing
- Table-driven tests for all usecases and handlers

## Security
- Path whitelisting, input validation, audit logs

## Documentation
- OpenAPI spec and usage examples
