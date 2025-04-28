# Web Browsing Service

This microservice performs web browsing and scraping via chromedp as part of the Tool Execution bounded context.

- **Tech stack:** Go, Gin, chromedp
- **Architecture:** Clean Architecture, DDD
- **API:**
  - POST `/browse` - Fetch and process a web page

## Structure
- `domain/` - Entities (BrowseJob)
- `usecase/` - Browsing logic
- `infrastructure/` - chromedp client
- `delivery/http/` - Gin HTTP handlers
- `main.go` - Entry point

## Testing
- Table-driven tests for all usecases and handlers

## Security
- Domain restriction, output sanitization, input validation

## Documentation
- OpenAPI spec and usage examples
