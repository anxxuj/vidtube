package jsonlog

import (
	"encoding/json"
	"io"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

// Represents severity levels for log entries.
type Level int8

const (
	LevelInfo Level = iota
	LevelError
	LevelFatal
)

func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

// Logger struct defines a JSON logger with output writer, minimum severity
// level, and mutex for concurrent writes.
type Logger struct {
	out      io.Writer
	minLevel Level
	mu       sync.Mutex
}

// New initializes and returns a new Logger instance.
func New(out io.Writer, minLevel Level) *Logger {
	return &Logger{out: out, minLevel: minLevel}
}

// PrintInfo logs an informational message.
func (l *Logger) PrintInfo(message string, properties map[string]string) {
	l.print(LevelInfo, message, properties)
}

// PrintError logs an error message.
func (l *Logger) PrintError(err error, properties map[string]string) {
	l.print(LevelError, err.Error(), properties)
}

// PrintFatal logs a fatal error message and exits the application.
func (l *Logger) PrintFatal(err error, properties map[string]string) {
	l.print(LevelFatal, err.Error(), properties)
	os.Exit(1)
}

// print writes a log entry in JSON format if its severity meets the minimum
// level.
func (l *Logger) print(level Level, message string, properties map[string]string) (int, error) {
	if level < l.minLevel {
		return 0, nil
	}

	aux := struct {
		Level      string            `json:"level"`
		Time       string            `json:"time"`
		Message    string            `json:"message"`
		Properties map[string]string `json:"properties,omitempty"`
		Trace      string            `json:"trace,omitempty"`
	}{
		Level:      level.String(),
		Time:       time.Now().UTC().Format(time.RFC3339),
		Message:    message,
		Properties: properties,
	}

	// Include stack trace for ERROR and FATAL levels.
	if level >= LevelError {
		aux.Trace = string(debug.Stack())
	}

	line, err := json.MarshalIndent(aux, "", "  ")
	if err != nil {
		line = []byte(LevelError.String() + ": unable to marshal log message: " + err.Error())
	}

	// Ensure thread-safe writes to the output.
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.out.Write(append(line, '\n'))
}
