package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	log *slog.Logger
}

func NewLogger() Logger {
	options := slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(os.Stdout, &options)
	return Logger{
		log: slog.New(handler),
	}
}

func (l *Logger) Debug(msg string) {
	l.log.Debug(msg)
}

func (l *Logger) Info(msg string) {
	l.log.Info(msg)
}

func (l *Logger) Warn(msg string) {
	l.log.Warn(msg)
}

func (l *Logger) Error(msg string) {
	l.log.Error(msg)
}

func (l *Logger) Fatal(msg string) {
	l.log.Error(msg)
	fmt.Printf("Message above caused fatal\n")
	os.Exit(1)
}
