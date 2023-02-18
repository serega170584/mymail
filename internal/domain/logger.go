package domain

import (
	"awesomeProject/internal/config"
	"bytes"
	"fmt"
	"log"
	"os"
)

const WARNING = "WARNING:"
const ERROR = "ERROR:"
const INFO = "INFO:"

type CustomLogger struct{}

func (l *CustomLogger) Warning(message string) {
	l.message(WARNING, message)
}

func (l *CustomLogger) Error(message string) {
	l.message(ERROR, message)
}

func (l *CustomLogger) Info(message string) {
	l.message(INFO, message)
}

func (l *CustomLogger) message(level string, message string) {
	mainConfig := config.NewConfig()

	isDebug := mainConfig.GetString("app.env") != "prod" && mainConfig.GetBool("app.debug")

	if isDebug {
		buf := bytes.Buffer{}
		customInfo := log.New(&buf, level, log.Lshortfile)
		customInfo.Output(2, message)
		fmt.Println(&buf)
	} else {
		f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Printf("Fatal: failed to open log: %s", err.Error())
		}

		customInfo := log.New(f, level, log.Lshortfile)
		customInfo.Println(message)

		defer f.Close()
	}
}
