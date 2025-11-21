package main

import (
	"encoding/json"
	"fmt"
	"gosms/task"
	"net/http"
)

type TaskInputJSON struct {
	To   string `json:"to"`
	Text string `json:"text"`
}

func WriteResponse(w http.ResponseWriter, statusCode int, body string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write([]byte(body))
}

func NewTaskView(w http.ResponseWriter, r *http.Request) {
	var taskInput TaskInputJSON
	err := json.NewDecoder(r.Body).Decode(&taskInput)

	if err != nil {
		WriteResponse(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	TaskID := task.CreateTask(taskInput.To, taskInput.Text)
	WriteResponse(w, http.StatusCreated, TaskID)
}

func HealthCheckView(w http.ResponseWriter, r *http.Request) {
	WriteResponse(w, http.StatusOK, "Hello Parsack !")

}

func main() {
	http.HandleFunc("/new-task", NewTaskView)
	http.HandleFunc("/health", HealthCheckView)
	fmt.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
