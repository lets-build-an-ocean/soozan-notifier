package main

import (
	"encoding/json"
	"fmt"
	"gosms/logger"
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
	logger.Info("Received new task request")

	var taskInput TaskInputJSON
	err := json.NewDecoder(r.Body).Decode(&taskInput)

	if err != nil {
		WriteResponse(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	logger.Info("Creating SMS task for: " + taskInput.To)

	TaskID := task.CreateTask(taskInput.To, taskInput.Text)

	logger.Info("Task created successfully: " + TaskID)

	WriteResponse(w, http.StatusCreated, TaskID)
}

func HealthCheckView(w http.ResponseWriter, r *http.Request) {
	logger.Debug("Health check endpoint accessed")
	WriteResponse(w, http.StatusOK, "Hello Parsack !")
}

func main() {
	logger.Info("Starting Soozan SMS Notifier service")

	http.HandleFunc("/new-task", NewTaskView)
	http.HandleFunc("/health", HealthCheckView)

	logger.Info("🚀 Server running on http://localhost:8080")
	fmt.Println("🚀 Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("Failed to start server: " + err.Error())
	}
}
