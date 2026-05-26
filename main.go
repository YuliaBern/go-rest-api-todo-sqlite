package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
	"github.com/YuliaBern/todo-api-sqlite/internal/todo"
)

func main() {
	// Підключення до SQLite
	db, err := sql.Open("sqlite", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Створення таблиці
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			completed BOOLEAN DEFAULT FALSE,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	repo := todo.NewRepository(db)
	service := todo.NewService(repo)
	handler := todo.NewHandler(service)

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handler.CreateTask(w, r)
		} else if r.Method == http.MethodGet {
			handler.GetTasks(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("🚀 Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}