package task

import (
	"fmt"
)

type LogLevel int

const (
	INFO LogLevel = iota
	ERROR
	FATAL
)

var levelNames = map[LogLevel]string{
	INFO:  "INFO",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

func logger(level LogLevel, task Task) {
	fmt.Printf("{'levelname': '%s', task: {'id' : '%s', 'number': '%s', 'scenario': '%s', 'params': '%s', 'status': '%s', }\n",
		levelNames[level], task.id, task.number, task.scenario, task.params, task.status)
}

func Info(task Task) {
	logger(INFO, task)
}

func Error(task Task) {
	logger(ERROR, task)
}

func Fatal(task Task) {
	logger(FATAL, task)
}
