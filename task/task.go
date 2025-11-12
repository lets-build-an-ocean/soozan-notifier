package task

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

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
	go HandleTask(task)
	return Id
}

func HandleTask(task Task) {
	maxAttempts := 3
	baseDelay := time.Second

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		status := provider.PushSMS(task.To, task.Text)
		fmt.Printf("task : %s - attemp : %d - status : %d\n", task.Id, attempt, status)

		if status == 200 {
			return
		} else {
			sleepDuration := baseDelay * (1 << (attempt - 1))
			time.Sleep(sleepDuration)
		}
	}
}
