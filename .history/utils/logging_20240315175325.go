package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string){
	//openFile means open or create file
	logfile, err := os.OpenFile(logFile, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil{
		log.Fatalln(err)
	}
	//multiLogFile means that create multi writer
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//setFlags means set log format
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	//setOutput means set where to write log
	log.SetOutput(multiLogFile)
}
