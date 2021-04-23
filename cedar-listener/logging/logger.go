package logging

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type LogLevel string

const (
	InfoLog  LogLevel = "info"
	DebugLog LogLevel = "debug"
	WarnLog  LogLevel = "warn"
	FatalLog LogLevel = "fatal"
	PanicLog LogLevel = "panic"
	ErrorLog LogLevel = "error"
	TraceLog LogLevel = "trace"
)

func setFieldsInJson() {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
		"time": time.Now(),
	})
}

func SetInfoLogLevel() {
	SetLogLevel(InfoLog)
}

func SetWarningLogLevel() {
	SetLogLevel(WarnLog)
}

func SetDebugLogLevel() {
	SetLogLevel(DebugLog)
}

func SetErrorLogLevel() {
	SetLogLevel(ErrorLog)
}

func SetPanicLogLevel() {
	SetLogLevel(PanicLog)
}

func SetTraceLogLevel() {
	SetLogLevel(TraceLog)
}

func SetFatalLogLevel() {
	SetLogLevel(FatalLog)
}

func SetLogLevel(level LogLevel) {
	switch level {
	case InfoLog:
		log.SetLevel(log.InfoLevel)
	case DebugLog:
		log.SetLevel(log.DebugLevel)
	case TraceLog:
		log.SetLevel(log.TraceLevel)
	case ErrorLog:
		log.SetLevel(log.ErrorLevel)
	case WarnLog:
		log.SetLevel(log.WarnLevel)
	case PanicLog:
		log.SetLevel(log.PanicLevel)
	case FatalLog:
		log.SetLevel(log.FatalLevel)
	}
}

func Debug(message string) {
	setFieldsInJson()
	log.Debug(message)
}

func Warn(message string) {
	setFieldsInJson()
	log.Warn(message)
}

func Error(message string) {
	setFieldsInJson()
	log.Error(message)
}

func Fatal(message string) {
	setFieldsInJson()
	log.Fatal(message)
}

func Info(message string) {
	setFieldsInJson()
	log.Info(message)
}

func Panic(message string) {
	setFieldsInJson()
	log.Panic(message)
}

func Trace(message string) {
	setFieldsInJson()
	log.Trace(message)
}
