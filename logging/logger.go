package logging

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func setFieldsInJson(fields map[string]fmt.Stringer) *log.Entry {
	if len(fields) == 0 {
		log.Fatal("No Fields provided to logger")
		return nil
	}
	requiredFields := make(map[string]interface{}, len(fields))
	for key, value := range fields {
		requiredFields[key] = value
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(requiredFields)
	return &log.Entry{}
}

func Debug(message string) {
	log.Debug(message)
}

func logDebugWithFields(fields map[string]fmt.Stringer, message string) {
	logger := setFieldsInJson(fields)
	if logger != nil {
		logger.Debug(message)
	}
}

func Error(message string) {
	log.Error(message)
}

func logErrorWithFields(fields map[string]fmt.Stringer, message string) {
	logger := setFieldsInJson(fields)
	if logger != nil {
		logger.Error(message)
	}
}

func Fatal(message string) {
	log.Fatal(message)
}

func logFatalWithFields(fields map[string]fmt.Stringer, message string) {
	logger := setFieldsInJson(fields)
	if logger != nil {
		logger.Fatal(message)
	}
}

func Info(message string) {
	log.Info(message)
}

func logInfoWithFields(fields map[string]fmt.Stringer, message string) {
	logger := setFieldsInJson(fields)
	if logger != nil {
		logger.Info(message)
	}
}

func Panic(message string) {
	log.Panic(message)
}

func Trace(message string) {
	log.Trace(message)
}

func logTraceWithFields(fields map[string]fmt.Stringer, message string) {
	logger := setFieldsInJson(fields)
	if logger != nil {
		logger.Trace(message)
	}
}
