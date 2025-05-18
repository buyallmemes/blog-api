package logging

import (
	"context"
	"io"
	"log/slog"
	"os"
)

// LogLevel represents the logging level
type LogLevel string

// Log levels
const (
	DebugLevel LogLevel = "DEBUG"
	InfoLevel  LogLevel = "INFO"
	WarnLevel  LogLevel = "WARN"
	ErrorLevel LogLevel = "ERROR"
)

// Logger is a wrapper around slog.Logger that provides additional functionality
type Logger struct {
	*slog.Logger
}

// Config holds the configuration for the logger
type Config struct {
	// Level is the minimum log level that will be output
	Level LogLevel
	// Output is where the logs will be written to
	Output io.Writer
	// AddSource adds the file and line number to the log output
	AddSource bool
}

// DefaultConfig returns a default configuration for the logger
func DefaultConfig() *Config {
	return &Config{
		Level:     InfoLevel,
		Output:    os.Stdout,
		AddSource: false,
	}
}

// New creates a new logger with the given configuration
func New(cfg *Config) *Logger {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	var level slog.Level
	switch cfg.Level {
	case DebugLevel:
		level = slog.LevelDebug
	case WarnLevel:
		level = slog.LevelWarn
	case ErrorLevel:
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: cfg.AddSource,
	}

	handler := slog.NewJSONHandler(cfg.Output, opts)
	logger := slog.New(handler)

	return &Logger{
		Logger: logger,
	}
}

// WithContext adds context values to the logger
func (l *Logger) WithContext(ctx context.Context) *Logger {
	// Extract any relevant values from context and add them to the logger
	// This is a placeholder - you can add actual context values as needed
	return l
}

// WithField adds a field to the logger
func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{
		Logger: l.Logger.With(key, value),
	}
}

// WithFields adds multiple fields to the logger
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	attrs := make([]any, 0, len(fields)*2)
	for k, v := range fields {
		attrs = append(attrs, k, v)
	}
	return &Logger{
		Logger: l.Logger.With(attrs...),
	}
}

// Default returns a default logger
func Default() *Logger {
	return New(DefaultConfig())
}
