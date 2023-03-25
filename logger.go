package eureka_client

import (
	"log"
	"os"
)

const (
	DEBUG = 4
	INFO  = 3
	WARN  = 2
	ERROR = 1
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
)

var logLevel = func() int {
	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	}
	return INFO
}()

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string, err error)
	Error(msg string, err error)
}

type DefaultLogger struct {
}

func (*DefaultLogger) Debug(msg string) {
	if logLevel >= DEBUG {
		log.Println(Green + "DEBUG: " + Reset + msg)
	}
}

func (*DefaultLogger) Info(msg string) {
	if logLevel >= INFO {
		log.Println(Cyan + "INFO: " + Reset + msg)
	}
}

func (*DefaultLogger) Warn(msg string, err error) {
	if logLevel >= WARN {
		log.Printf("%sWARNING: %s%s, %v\n", Yellow, Reset, msg, err)
	}
}

func (*DefaultLogger) Error(msg string, err error) {
	if logLevel >= ERROR {
		log.Printf("%sERROR: %s%s, %v\n", Red, Reset, msg, err)
	}
}

func NewLogger() Logger {
	return &DefaultLogger{}
}
