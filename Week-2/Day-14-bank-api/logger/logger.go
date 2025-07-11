package logger

import (
	"encoding/json"
	"log"
	"time"
)

type LogEntry struct {
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Details   any       `json:"details,omitempty"`
}

func logJSON(level, msg string, details any) {
	entry := LogEntry{
		Level:     level,
		Message:   msg,
		Timestamp: time.Now(),
		Details:   details,
	}

	jsonData, err := json.Marshal(entry)
	if err != nil {
		log.Printf("[fallback-log] %s: %+v\n", level, entry)
		return
	}

	log.Println(string(jsonData))
}

func Info(msg string, details any)  { logJSON("info", msg, details) }
func Warn(msg string, details any)  { logJSON("warn", msg, details) }
func Error(msg string, details any) { logJSON("error", msg, details) }
