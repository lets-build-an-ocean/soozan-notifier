package main

import (
	"log"
	"net/http"
	"notifier/config"
	"notifier/views"
)

func main() {
	err := config.LoadConfig("../config.json")
	if err != nil {
		log.Fatalf("Critical error loading config: %v", err)
	}
	http.HandleFunc("/new-task", views.NewTaskView)
	http.HandleFunc("/health", views.HealthCheckView)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}
