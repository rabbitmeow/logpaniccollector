package logpaniccollector

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/robfig/cron/v3"
)

// LogPanic ...
type LogPanic struct {
	LogFile string
}

// New making new instance of this library
func New() *LogPanic {
	lpc := &LogPanic{
		LogFile: "service.log",
	}
	return lpc
}

// AutoRemoveLog for auto removing the log file based on the time.Duration
func (l *LogPanic) AutoRemoveLog(cronFormat string) {
	c := cron.New()
	c.AddFunc(cronFormat, func() {
		os.Remove(l.LogFile)
	})
	c.Start()
}

// RecoverPanic recover your panic, then write it in your log file
func (l *LogPanic) RecoverPanic(uri string, recover interface{}) bool {
	if recover != nil {
		errPanic := fmt.Sprintf("endpoint: %s | panic: %v", uri, recover)
		l.WritePanic(errPanic, debug.Stack())
		return true
	}
	return false
}

// WriteLog is use for write log to defined log file
func (l *LogPanic) WriteLog(logMsg string) {
	l.writer(logMsg, "]")
}

// WritePanic is use for write panic to defined log file
func (l *LogPanic) WritePanic(e interface{}, stack []byte) {
	l.writer(e, "] \n", string(stack))
}

func (l *LogPanic) writer(errFormat ...interface{}) {
	f, err := os.OpenFile(l.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
