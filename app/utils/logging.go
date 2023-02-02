package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

func LoggingSettings(logfileName string) {
	// Set log file
	logfile, err := os.OpenFile(logfileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer logfile.Close()

	multiLogFile := io.MultiWriter(os.Stdout, logfile)

	// Set log format
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
