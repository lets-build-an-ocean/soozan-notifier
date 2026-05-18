package views

import (
	"encoding/json"
	"net/http"
	"notifier/task"
)

type TaskInputJSON struct {
	Number   string   `json:"number"`
	Scenario string   `json:"scenario"`
	Params   []string `json:"params"`
}

func NewTaskView(w http.ResponseWriter, r *http.Request) {
	var taskInput TaskInputJSON
	err := json.NewDecoder(r.Body).Decode(&taskInput)

	if err != nil {
		WriteResponse(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	TaskID := task.CreateTask(taskInput.Number, taskInput.Scenario, taskInput.Params)
	WriteResponse(w, http.StatusCreated, TaskID)
}

func HealthCheckView(w http.ResponseWriter, r *http.Request) {
	WriteResponse(w, http.StatusOK, "OK")
}
