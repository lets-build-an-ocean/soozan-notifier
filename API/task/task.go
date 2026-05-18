package task

import (
	"notifier/config"
	"notifier/providers"
	"time"
)

type Task struct {
	id       string
	number   string
	scenario string
	params   []string
	status   string
}

func CreateTask(number string, scenario string, params []string) (TaskID string) {
	Id := generateTaskID()
	task := Task{
		id:       Id,
		number:   number,
		scenario: scenario,
		params:   params,
		status:   "pending",
	}

	if redirects := config.GetTestNumberRedirects(task.number); len(redirects) > 0 {
		for _, redirectNum := range redirects {
			task.number = redirectNum
			go HandleTask(task)
		}
	} else {
		go HandleTask(task)
	}

	return Id
}

func HandleTask(task Task) {
	status := providers.Main(task.number, task.scenario, task.params)
	task.status = "main-1"
	if status == 200 {
		Info(task)
		return
	}

	time.Sleep(time.Second)
	status = providers.Main(task.number, task.scenario, task.params)
	task.status = "main-2"
	if status == 200 {
		Info(task)
		return
	}

	time.Sleep(time.Second)
	status = providers.Fallback(task.number, task.scenario, task.params)
	task.status = "fallback-1"
	if status == 200 {
		Info(task)
		return
	}

	time.Sleep(time.Second)
	status = providers.Fallback(task.number, task.scenario, task.params)
	task.status = "fallback-2"
	if status == 200 {
		Info(task)
		return
	}

	task.status = "failed"
	Error(task)
}
