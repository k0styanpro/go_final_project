package main

import (
	"log"
	"os"

	"go_final_project/pkg/api"
	"go_final_project/pkg/db"
	"go_final_project/pkg/server"
)

func main() {
	dbFile := getDBFile()
	if err := db.Init(dbFile); err != nil {
		log.Fatalf("❌ Database init failed: %v", err)
	}

	api.Init()

	log.Printf("🚀 Starting server on port %d", server.GetPort())
	if err := server.Run(); err != nil {
		log.Fatalf("💥 Server crashed: %v", err)
	}
}
\
func getDBFile() string {
	if dbFile := os.Getenv("TODO_DBFILE"); dbFile != "" {
		return dbFile
	}
	return "scheduler.db"
}
