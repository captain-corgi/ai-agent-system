// main.go
package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/captain-corgi/ai-agent-system/task-service/infrastructure"
	"github.com/captain-corgi/ai-agent-system/task-service/delivery/http"
)

func main() {
	// Restrict workspace (security)
	workspace := "/app/workspace"
	if _, err := os.Stat(workspace); os.IsNotExist(err) {
		log.Fatalf("Workspace directory %s does not exist", workspace)
	}

	dbPath := workspace + "/task.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}
	defer db.Close()

	if err := infrastructure.Migrate(db); err != nil {
		log.Fatalf("failed to migrate db: %v", err)
	}

	repo := infrastructure.NewSQLiteTaskRepository(db)

	handler := http.NewHandler(repo)

	r := gin.Default()
	handler.RegisterRoutes(r)

	addr := ":8080"
	log.Printf("Task service running on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
