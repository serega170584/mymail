package logger

import (
	"bytes"
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

func New(isProd bool, isDebug bool, filename string, prefix string) *Logger {
	return &Logger{!isProd && isDebug, filename}
}

func (logger Logger) sendMessageToBuffer(message string, level string) {
	writer := &bytes.Buffer{}

	log.New(writer, level, log.Ltime).Println(message)
}

func (logger Logger) sendMessageToFile(message string, level string) error {
	file, err := os.OpenFile(logger.filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Fatal: failed to open log: %s", err.Error())
		return err
	}

	log.New(file, level, log.Ltime).Println(message)

	defer file.Close()

	return nil
}

func (logger *Logger) sendMessage(message string, level string) error {
	if logger.isDebug {
		logger.sendMessageToBuffer(message, level)
	} else {
		return logger.sendMessageToFile(message, level)
	}

	return nil
}

func (logger *Logger) Warning(message string) error {
	return logger.sendMessage(message, WARNING)
}

func (logger *Logger) Error(message string) error {
	return logger.sendMessage(message, ERROR)
}

func (logger *Logger) Info(message string) error {
	return logger.sendMessage(message, INFO)
}
