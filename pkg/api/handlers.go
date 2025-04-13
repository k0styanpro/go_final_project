package api

import (
	"encoding/json"
	"github.com/k0styanpro/go_final_project/pkg/db"
	"net/http"
	"time"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTask(w, r)
	case http.MethodPut:
		updateTask(w, r)
	case http.MethodDelete:
		deleteTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	tasks, _ := db.GetTasks(search) // с поддержкой поиска
	jsonResponse(w, tasks)
}

func DoneHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	task, _ := db.GetTask(id)
	nextDate, _ := NextDate(time.Now(), task.Date, task.Repeat)
	db.UpdateDate(id, nextDate)
	jsonResponse(w, map[string]string{})
}
