package logger

import (
	"log"
	"os"
)

// Global logger instances
var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

// InitLogger initializes loggers
func InitLogger() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	// Info logs (console + file)
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Error logs (file)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}