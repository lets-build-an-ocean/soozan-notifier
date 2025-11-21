package logger

import (
	"fmt"
	"os"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

var levelNames = map[LogLevel]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
	FATAL: "FATAL",
}

var currentLevel = INFO

func SetLevel(level LogLevel) {
	currentLevel = level
}

func log(level LogLevel, message string) {
	if level < currentLevel {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] %s: %s\n", timestamp, levelNames[level], message)

	if level == FATAL {
		os.Exit(1)
	}
}

func Debug(message string) {
	log(DEBUG, message)
}

func Info(message string) {
	log(INFO, message)
}

func Warn(message string) {
	log(WARN, message)
}

func Error(message string) {
	log(ERROR, message)
}

func Fatal(message string) {
	log(FATAL, message)
}