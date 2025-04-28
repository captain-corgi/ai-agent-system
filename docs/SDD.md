# Software Design Document (SDD)

## 1. Overview
- Clean Architecture, DDD, Microservices, Micro Frontends
- Go for backend, React/Vite/TS/Tailwind for frontend

## 2. Component Architecture

```mermaid
graph TD
  FE[Frontend Shell]
  TaskApp[Task Management App]
  ResultApp[Result Viewer App]
  TaskSvc[Task Service]
  AISvc[AI Service]
  CodeExec[Code Execution Service]
  WebBrowse[Web Browsing Service]
  FSSvc[Filesystem Service]

  FE-->|Module Federation|TaskApp
  FE-->|Module Federation|ResultApp
  TaskApp-->|REST|TaskSvc
  ResultApp-->|REST|TaskSvc
  TaskSvc-->|REST|AISvc
  TaskSvc-->|REST|CodeExec
  TaskSvc-->|REST|WebBrowse
  TaskSvc-->|REST|FSSvc
```

---

## 3. Feature Design Breakdown

### 3.1 Task Management
- **Purpose:** Manage task lifecycle as aggregate root.
- **Components:**
  - `domain/Task.go`: Entity, status enum
  - `usecase/CreateTask.go`, `GetTask.go`, `ListTasks.go`: Business logic
  - `infrastructure/TaskRepository.go`: SQLite repo
  - `delivery/http/TaskHandler.go`: Gin handlers
- **Sequence Diagram:**
```mermaid
sequenceDiagram
  participant User
  participant FE
  participant TaskSvc
  participant DB
  User->>FE: Submit Task
  FE->>TaskSvc: POST /tasks
  TaskSvc->>DB: Insert Task
  TaskSvc-->>FE: Task ID
  User->>FE: Poll Status
  FE->>TaskSvc: GET /tasks/:id
  TaskSvc->>DB: Query Task
  TaskSvc-->>FE: Status/Result
```
- **Quality/Security:** Input validation, parameterized queries, audit logs.
- **Extensibility:** Add new task types by extending domain/usecase.

### 3.2 Local AI Reasoning
- **Purpose:** Generate stepwise plans for tasks using local LLM.
- **Components:**
  - `domain/Plan.go`, `Step.go`: Entities
  - `usecase/GeneratePlan.go`: Business logic
  - `infrastructure/OllamaClient.go`: REST integration
  - `delivery/http/AIHandler.go`: Gin handler
- **Sequence Diagram:**
```mermaid
sequenceDiagram
  participant TaskSvc
  participant AISvc
  participant Ollama
  TaskSvc->>AISvc: POST /ai/plan
  AISvc->>Ollama: Generate Plan (prompt)
  Ollama-->>AISvc: Plan (JSON)
  AISvc-->>TaskSvc: Plan
```
- **Quality/Security:** Prompt sanitization, local-only requests, logging.
- **Extensibility:** Support new plan types by updating prompt templates and parser.

### 3.3 Secure Code Execution
- **Purpose:** Run user code securely in Docker sandbox.
- **Components:**
  - `domain/CodeJob.go`: Entity
  - `usecase/RunCode.go`: Business logic
  - `infrastructure/DockerRunner.go`: Docker API
  - `delivery/http/CodeHandler.go`: Gin handler
- **Sequence Diagram:**
```mermaid
sequenceDiagram
  participant TaskSvc
  participant CodeExec
  participant Docker
  TaskSvc->>CodeExec: POST /execute/code
  CodeExec->>Docker: Start Container
  Docker-->>CodeExec: Output/Exit
  CodeExec-->>TaskSvc: Result
```
- **Quality/Security:** Container resource limits, no host mounts, network off.
- **Extensibility:** Add new language runners by extending DockerRunner.

### 3.4 Autonomous Web Browsing
- **Purpose:** Automate web navigation and scraping.
- **Components:**
  - `domain/BrowseJob.go`: Entity
  - `usecase/FetchPage.go`: Business logic
  - `infrastructure/ChromedpClient.go`: chromedp integration
  - `delivery/http/WebHandler.go`: Gin handler
- **Sequence Diagram:**
```mermaid
sequenceDiagram
  participant TaskSvc
  participant WebBrowse
  participant Chrome
  TaskSvc->>WebBrowse: POST /browse
  WebBrowse->>Chrome: Fetch Page
  Chrome-->>WebBrowse: HTML/Text
  WebBrowse-->>TaskSvc: Result
```
- **Quality/Security:** Domain allowlist, output sanitization.
- **Extensibility:** Add new scraping strategies in usecase.

### 3.5 Secure Filesystem Operations
- **Purpose:** Perform safe file operations in workspace dir only.
- **Components:**
  - `domain/FsJob.go`: Entity
  - `usecase/ReadFile.go`, `WriteFile.go`, `DeleteFile.go`: Logic
  - `infrastructure/PathSanitizer.go`: Path validation
  - `delivery/http/FsHandler.go`: Gin handler
- **Sequence Diagram:**
```mermaid
sequenceDiagram
  participant TaskSvc
  participant FSSvc
  participant OS
  TaskSvc->>FSSvc: POST /fs/write/read/delete
  FSSvc->>OS: File Op
  OS-->>FSSvc: Result
  FSSvc-->>TaskSvc: Result
```
- **Quality/Security:** Path whitelisting, audit logs.
- **Extensibility:** Support new file ops by adding usecases.

### 3.6 Micro Frontend UI
- **Purpose:** Modular, responsive UI for all user interactions.
- **Components:**
  - `shell/`: Hosts and loads remotes
  - `task-management-app/`: Task submission, list
  - `result-viewer-app/`: Result display
- **Sequence Diagram:**
```mermaid
sequenceDiagram
  participant User
  participant Shell
  participant TaskApp
  participant ResultApp
  participant TaskSvc
  User->>Shell: Open UI
  Shell->>TaskApp: Mount
  User->>TaskApp: Submit Task
  TaskApp->>TaskSvc: POST /tasks
  User->>ResultApp: View Result
  ResultApp->>TaskSvc: GET /tasks/:id
```
- **Quality/Security:** Responsive, accessible, input validation.
- **Extensibility:** Add new micro frontends via Module Federation.

---

## 4. Backend Microservice Structure
- Each service: `domain/`, `usecase/`, `infrastructure/`, `delivery/http/`, `main.go`
- Communication: HTTP/JSON
- Storage: SQLite (Task), Docker (Code), chromedp (Web), OS (FS)

## 5. Frontend Micro Frontend Structure
- Shell: Loads remotes via Vite Module Federation
- Task Management App: Task submission, list, status
- Result Viewer App: Task result display
- Styling: TailwindCSS, responsive design

## 5. Sequence Diagram: Task Lifecycle

```mermaid
sequenceDiagram
  participant User
  participant FE
  participant TaskSvc
  participant AISvc
  participant CodeExec
  participant WebBrowse
  participant FSSvc

  User->>FE: Submit Task
  FE->>TaskSvc: POST /tasks
  TaskSvc->>AISvc: POST /ai/plan
  AISvc-->>TaskSvc: Plan
  loop steps
    alt code step
      TaskSvc->>CodeExec: POST /execute/code
      CodeExec-->>TaskSvc: Output
    else browse step
      TaskSvc->>WebBrowse: POST /browse
      WebBrowse-->>TaskSvc: Output
    else fs step
      TaskSvc->>FSSvc: POST /fs/write/read
      FSSvc-->>TaskSvc: Output
    end
  end
  TaskSvc-->>FE: Task Result
  FE-->>User: Display Result
```

## 6. Deployment Diagram

```mermaid
graph LR
  subgraph User Machine
    FE[Frontend Shell]
    TaskApp
    ResultApp
    TaskSvc
    AISvc
    CodeExec
    WebBrowse
    FSSvc
    Ollama
    Docker
    Chrome
    SQLite
  end
  FE-->|REST|TaskSvc
  FE-->|REST|AISvc
  FE-->|REST|CodeExec
  FE-->|REST|WebBrowse
  FE-->|REST|FSSvc
  TaskSvc-->|DB|SQLite
  AISvc-->|LLM|Ollama
  CodeExec-->|Sandbox|Docker
  WebBrowse-->|Headless|Chrome
```

## 7. Quality Attributes
- Security: Docker sandbox, FS restriction, input validation
- Testability: Table-driven, E2E, coverage
- Maintainability: Modular, documented, DDD
- Performance: Efficient APIs, optimized queries
