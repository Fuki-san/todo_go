package utils

import (
	"log"
	"os"
)

func LoggingSettings(logFile string){
	logfile, err := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil{
		log.Fatalln(err)
	}
	multiLogFile := io.Mu
}
