# Web Browsing Service

This microservice performs web browsing and scraping via chromedp as part of the Tool Execution bounded context.

- **Tech stack:** Go, Gin, chromedp
- **Architecture:** Clean Architecture, DDD
- **APIs:**
  - `POST /browse` - Fetch and process a web page
    - **Request:** `{ "url": "https://..." }`
    - **Response:** `{ "result": "..." }`
    - **Errors:** `400 Bad Request`, `500 Internal Server Error`
  - `GET /browse/health` - Health check endpoint
    - **Response:** `{ "status": "ok" }`

## Structure Diagram
```mermaid
graph TD
  WebDomain[domain/]
  WebUsecase[usecase/]
  WebInfra[infrastructure/]
  WebDeliv[delivery/http/]
  Main[main.go]
  WebDomain-->WebUsecase
  WebUsecase-->WebInfra
  WebInfra-->WebDeliv
  WebDeliv-->Main
```

## Features
- Automated web browsing and scraping
- Domain restriction and output sanitization
- Integration with chromedp

## Data Flow Diagram (DFD)
```mermaid
graph TD
  FE[Frontend]-->|POST /browse|WebBrowse[Web Browsing Service]
  WebBrowse-->|Headless|Chrome[chromedp]
  Chrome-->|Result|WebBrowse
  WebBrowse-->|Output|FE
```

## Entity Relationship Diagram (ERD)
```mermaid
erDiagram
  BROWSEJOB {
    string id PK
    string task_id FK
    string url
    string result
    string status
  }
```

## Database Table
| Field      | Type   | PK | FK | Description         |
|------------|--------|----|----|---------------------|
| id         | TEXT   | Y  |    | Job ID              |
| task_id    | TEXT   |    | Y  | FK to Task          |
| url        | TEXT   |    |    | URL to fetch        |
| result     | TEXT   |    |    | Output/result       |
| status     | TEXT   |    |    | Job status          |

## Testing
- Table-driven tests for all usecases and handlers

## Security
- Domain restriction, output sanitization, input validation

## Documentation
- OpenAPI spec and usage examples
