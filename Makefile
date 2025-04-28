# Makefile for AI Agent System

.PHONY: all ai-service task-service code-execution-service filesystem-service web-browsing-service frontend build clean run-all

all: build

build: 
	cd backend/ai-service && go build -o ai-service
	cd backend/task-service && go build -o task-service
	cd backend/code-execution-service && go build -o code-execution-service
	cd backend/filesystem-service && go build -o filesystem-service
	cd backend/web-browsing-service && go build -o web-browsing-service

ai-service:
	cd backend/ai-service && go run main.go

task-service:
	cd backend/task-service && go run main.go

code-execution-service:
	cd backend/code-execution-service && go run main.go

filesystem-service:
	cd backend/filesystem-service && go run main.go

web-browsing-service:
	cd backend/web-browsing-service && go run main.go

frontend:
	cd frontend/shell && npm run dev & \
	cd ../task-management-app && npm run dev & \
	cd ../result-viewer-app && npm run dev

run-all: 
	$(MAKE) -j5 ai-service task-service code-execution-service filesystem-service web-browsing-service

ai-service-integration-test:
	cd backend/ai-service && go test -v ./integration/...

task-service-integration-test:
	cd backend/task-service && go test -v ./integration/...

code-execution-service-integration-test:
	cd backend/code-execution-service && go test -v ./integration/...

filesystem-service-integration-test:
	cd backend/filesystem-service && go test -v ./integration/...

web-browsing-service-integration-test:
	cd backend/web-browsing-service && go test -v ./integration/...

# Run all integration tests for backend microservices
backend-integration-test: ai-service-integration-test task-service-integration-test code-execution-service-integration-test filesystem-service-integration-test web-browsing-service-integration-test

# Unified test target (can be extended for frontend tests)
test: backend-integration-test

clean:
	rm -f backend/ai-service/ai-service backend/task-service/task-service backend/code-execution-service/code-execution-service backend/filesystem-service/filesystem-service backend/web-browsing-service/web-browsing-service
