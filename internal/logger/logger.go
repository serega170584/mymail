package logger

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const WARNING = "WARNING:"
const ERROR = "ERROR:"
const INFO = "INFO:"

type Logger struct {
	isDebug  bool
	filename string
}

func New(isProd bool, isDebug bool, filename string) *Logger {
	return &Logger{!isProd && isDebug, filename}
}

func (logger *Logger) sendMessageToBuffer(message string, level string) {
	writer := &bytes.Buffer{}

	baseLogger := log.New(writer, level, log.Llongfile)
	baseLogger.Output(2, message)
	fmt.Println(writer)
}

func (logger *Logger) sendMessageToFile(message string, level string) {
	file, err := os.OpenFile(logger.filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Fatal: failed to open log: %s", err.Error())
	}

	log.New(file, level, log.Llongfile).Println(message)

	err = file.Close()
	if err != nil {
		log.Printf("Fatal: failed to close log: %s", err.Error())
	}
}

func (logger *Logger) sendMessage(message string, level string) {
	if logger.isDebug {
		logger.sendMessageToBuffer(message, level)
	} else {
		logger.sendMessageToFile(message, level)
	}
}

func (logger *Logger) Warning(message string) {
	logger.sendMessage(message, WARNING)
}

func (logger *Logger) Error(message string) {
	logger.sendMessage(message, ERROR)
}

func (logger *Logger) Info(message string) {
	logger.sendMessage(message, INFO)
}
