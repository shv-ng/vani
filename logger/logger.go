package logger

import (
	"log"
	"os"
)

var Log *log.Logger

func Init(filename string) {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o666)
	if err != nil {
		panic("hey, you didn't give me good file")
	}
	Log = log.New(logFile, "[VANI] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(msg string) {
	_ = Log.Output(2, "[INFO] "+msg)
}

func Warn(msg string) {
	_ = Log.Output(2, "[WARN] "+msg)
}

func Error(msg string) {
	_ = Log.Output(2, "[ERROR] "+msg)
}
