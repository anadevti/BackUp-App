package backup

import (
	"log"
	"os"
)

var (
	logFile *os.File
	logger  *log.Logger
)

func InitLogger() error {
	var err error
	logFile, err = os.OpenFile("backup.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	logger = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

func LogInfo(msg string) {
	if logger != nil {
		logger.Println("INFO: " + msg)
	}
}

func LogError(msg string) {
	if logger != nil {
		logger.Println("ERROR: " + msg)
	}
}
