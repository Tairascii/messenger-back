package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	Log *Logger
)

func init() {
	Log = New()
}

type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func New() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "", log.LstdFlags),
		errorLogger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

func (l *Logger) Info(v ...any) {
	l.infoLogger.Println(formatMessage("INFO", v...))
}

func (l *Logger) Infof(format string, v ...any) {
	l.infoLogger.Printf(formatMessage("INFO", format), v...)
}

func (l *Logger) Warn(v ...any) {
	l.infoLogger.Println(formatMessage("WARN", v...))
}

func (l *Logger) Warnf(format string, v ...any) {
	l.infoLogger.Printf(formatMessage("WARN", format), v...)
}

func (l *Logger) Error(v ...any) {
	l.errorLogger.Println(formatMessage("ERROR", v...))
}

func (l *Logger) Errorf(format string, v ...any) {
	l.errorLogger.Printf(formatMessage("ERROR", format), v...)
}

func (l *Logger) Fatal(v ...any) {
	l.errorLogger.Println(formatMessage("FATAL", v...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...any) {
	l.errorLogger.Printf(formatMessage("FATAL", format), v...)
	os.Exit(1)
}

func formatMessage(level string, v ...any) string {
	return fmt.Sprintf("[%s] %s", level, fmt.Sprint(v...))
}
