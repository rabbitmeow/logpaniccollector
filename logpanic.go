package logpaniccollector

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

var fullPathToFile string

// SetFile for setting your full path to file. It will create the file that contains your log or panic. E.g /var/log/service.log
func SetFile(fullPath string) {
	fullPathToFile = fullPath
}

// WriteLog is use for write log to defined log file
func WriteLog(logMsg string) {
	writer(logMsg, "]")
}

// WritePanic is use for write panic to defined log file
func WritePanic(e interface{}, stack []byte) {
	writer(e, "] \n", string(stack))
}

func writer(errFormat ...interface{}) {
	if fullPathToFile == "" {
		fullPathToFile = "service.log"
	}
	f, err := os.OpenFile(fullPathToFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer f.Close()
	var buf bytes.Buffer
	logger := log.New(&buf, "[", log.Ldate|log.Ltime)
	logger.SetOutput(f)
	logger.Println(errFormat...)
	fmt.Print(&buf)
}
