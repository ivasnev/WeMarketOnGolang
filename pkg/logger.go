package pkg

import (
	"log"
	"os"
)

var logger *log.Logger

// InitLogger инициализирует логгер с заданным уровнем логирования.
func InitLogger() {
	logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info выводит информационное сообщение в лог.
func Info(message string) {
	logger.Println(message)
}

// Error выводит сообщение об ошибке в лог.
func Error(message string) {
	logger.Println("ERROR:", message)
}
