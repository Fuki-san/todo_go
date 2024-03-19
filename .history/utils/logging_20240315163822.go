package utils

import "os"

func LoggingSettings(logFile string){
	logfile, err := os.OpenFile(logFile, os.O_RDWR | os.CRE)
}
