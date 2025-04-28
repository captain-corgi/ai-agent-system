# Comprehensive Prompt for Cursor

You are tasked with developing a local AI agent system, inspired by Manus.ai and agenticSeek, that runs entirely on the user’s hardware for privacy and cloud independence. The system should support autonomous task execution, including coding, web browsing, task planning, and filesystem interaction. The project must strictly follow Domain-Driven Design (DDD), use microservices with Clean Architecture for the backend, and implement a Micro Front End architecture for the frontend. All unit tests should employ Table-Driven Testing. Generate the entire project with the highest quality, adhering to the Software Development Lifecycle (SDLC) phases: planning, design, implementation, testing, deployment, and maintenance.

## Project Requirements and Architecture

### Backend (Golang)
- **Microservices**: Implement the following independent services:
  - **Task Service**: Manages task lifecycle (submission, status, results).
  - **AI Service**: Interfaces with the local LLM (Ollama) for reasoning and planning.
  - **Code Execution Service**: Executes code in sandboxed Docker containers.
  - **Web Browsing Service**: Navigates web pages using chromedp.
  - **Filesystem Service**: Performs secure file operations within a designated directory.
- **Clean Architecture**: Each microservice should follow these layers:
  - **Entities**: Core domain models (e.g., Task).
  - **Use Cases**: Business logic (e.g., CreateTaskUseCase).
  - **Interface Adapters**: Repositories and controllers (e.g., SQLiteTaskRepository).
  - **Frameworks/Drivers**: External systems (e.g., Gin for APIs, SQLite for storage).
- **Domain-Driven Design (DDD)**:
  - **Bounded Contexts**: Task Management, AI Reasoning, Tool Execution (Code Execution, Web Browsing, Filesystem Operations).
  - **Entities and Aggregates**: Task as an aggregate root in Task Management.
  - **Repositories**: Interfaces for data access, implemented in the infrastructure layer.
- **Communication**: Services communicate via HTTP/JSON APIs.
- **Security**: Use Docker for sandboxed code execution and restrict filesystem operations to a safe directory (e.g., `/app/workspace`).

### Frontend (ReactJS with Vite, TypeScript, TailwindCSS)
- **Micro Front End Architecture**:
  - **Shell Application**: Hosts micro frontends using Vite’s Module Federation plugin.
  - **Task Management App**: Handles task submission and listing.
  - **Result Viewer App**: Displays task results.
- **Integration**: Apps communicate with backend services via REST APIs, using polling for task status updates.
- **Styling**: Use TailwindCSS for responsive design.

### AI Integration
- **Local LLM**: Use Ollama with a model like DeepSeek-R1.
- **Integration**: AI Service communicates with Ollama’s REST API to process tasks.

### Testing
- **Unit Tests**: Use Table-Driven Testing in Go for backend services.
- **Frontend Tests**: Use Jest and React Testing Library, applying table-driven principles where possible.
- **Integration Tests**: Validate microservice and frontend-backend interactions.

## Implementation Instructions

1. **Set Up the Project Structure**:
   - Create a root directory `ai-agent-system`.
   - Inside, create subdirectories for each microservice: `task-service`, `ai-service`, `code-execution-service`, `web-browsing-service`, `filesystem-service`.
   - For the frontend, create `shell`, `task-management-app`, and `result-viewer-app`.

2. **Backend Development**:
   - For each microservice, initialize a Go module (e.g., `go mod init github.com/yourusername/ai-agent-system/task-service`).
   - Implement the Clean Architecture layers:
     - Define entities in `domain/`.
     - Implement use cases in `usecase/`.
     - Create repositories in `infrastructure/`.
     - Set up API handlers in `delivery/http/`.
   - Use Gin for API routing and SQLite for data storage in the Task Service.
   - Integrate Ollama in the AI Service using `github.com/jmorganca/ollama`.
   - Use the Docker Go client for the Code Execution Service.
   - Implement chromedp for the Web Browsing Service.
   - Use Golang’s `os` package for the Filesystem Service, ensuring operations are restricted to a safe directory.

3. **Frontend Development**:
   - Set up Vite projects for the shell and apps with React, TypeScript, and TailwindCSS.
   - Configure Module Federation in `vite.config.ts` to load apps dynamically in the shell.
   - Implement components for task submission, listing, and result viewing.
   - Use `fetch` or Axios to communicate with backend APIs.

4. **AI Integration**:
   - Install Ollama and download the DeepSeek-R1 model.
   - Start the Ollama server and ensure the AI Service can communicate with it.

5. **Testing**:
   - Write unit tests for backend services using Go’s `testing` package with table-driven tests.
   - Test frontend components with Jest and React Testing Library.
   - Perform integration tests to validate end-to-end functionality.

6. **Security**:
   - Ensure code execution is sandboxed in Docker containers.
   - Restrict filesystem operations to a designated directory.
   - Implement input validation on APIs and frontend.

7. **Deployment**:
   - Create installation scripts for Golang, Node.js, Docker, and Ollama.
   - Compile backend services into binaries.
   - Build frontend static assets.
   - Document installation and usage instructions in a `README.md`.

## Final Instructions
- Ensure all components are implemented according to the specified architectural patterns (DDD, Clean Architecture, Micro Front Ends).
- Prioritize security, especially in code execution and filesystem operations.
- Follow the SDLC phases to maintain a structured development process.
- Use Table-Driven Testing to ensure comprehensive unit test coverage.
- Document the system thoroughly for ease of installation and use.
- Generate the full project, including all source code, configuration files, and documentation, with the highest quality and attention to detail.
- Don't forget to create all '.windsurfrules' at every folder that has functionality