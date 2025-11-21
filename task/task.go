package task

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"gosms/logger"
	"gosms/provider"
)

type Task struct {
	Id   string
	To   string
	Text string
}

func generateTaskID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func CreateTask(to string, text string) (TaskID string) {
	Id := generateTaskID()
	task := Task{Id: Id, To: to, Text: text}
	
	logger.Info("Task created and queued: " + Id)
	
	go HandleTask(task)
	return Id
}

func HandleTask(task Task) {
	maxAttempts := 3
	baseDelay := time.Second

	logger.Info("Starting task processing: " + task.Id)

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		logger.Info(fmt.Sprintf("Attempt %d for task %s", attempt, task.Id))

		status := provider.PushSMS(task.To, task.Text)

		if status == 200 {
			logger.Info(fmt.Sprintf("SMS delivered successfully: %s", task.Id))
			return
		} else {
			sleepDuration := baseDelay * (1 << (attempt - 1))
			
			if attempt < maxAttempts {
				logger.Warn(fmt.Sprintf("SMS delivery failed (status %d), retrying in %s: %s", status, sleepDuration, task.Id))
				time.Sleep(sleepDuration)
			} else {
				logger.Error(fmt.Sprintf("SMS delivery failed after %d attempts: %s", maxAttempts, task.Id))
			}
		}
	}
}
