package domain

import (
	"log"
	"os"
)

type CustomLogger struct{}

func (l *CustomLogger) Warning(message string) {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 777)
	if err != nil {
		log.Printf("Fatal: failed to open log: %s", err.Error())
	}

	customWarning := log.New(f, "WARNING:", log.Lshortfile)
	customWarning.Println(message)

	defer f.Close()
}

func (l *CustomLogger) Error(message string) {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 777)
	if err != nil {
		log.Printf("Fatal: failed to open log: %s", err.Error())
	}

	customError := log.New(f, "ERROR:", log.Lshortfile)
	customError.Println(message)

	defer f.Close()
}

func (l *CustomLogger) Info(message string) {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 777)
	if err != nil {
		log.Printf("Fatal: failed to open log: %s", err.Error())
	}

	customInfo := log.New(f, "INFO:", log.Lshortfile)
	customInfo.Println(message)

	defer f.Close()
}
