package logger

import (
	"fmt"
	"log"
)

type LogInterface interface {
	Debug(i ...interface{})
	Error(i ...interface{})
	Info(i ...interface{})
}

type Logger struct{}

func (logger *Logger) Debug(i ...interface{}) {
	log.Printf("[DEBUG] %s", fmt.Sprintln(i...))
}

func (logger *Logger) Error(i ...interface{}) {
	log.Printf("[ERROR] %s", fmt.Sprintln(i...))
}

func (logger *Logger) Info(i ...interface{}) {
	log.Printf("[INFO] %s", fmt.Sprintln(i...))
}
