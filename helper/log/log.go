package log

import (
	"fmt"
	"log"
)

var (
	Info  int = 1
	Error int = 2
	Fatal int = 3
)

func Log(level int, message string) {
	var logFlag string = ""
	var exit bool = false
	if level >= 4 {
		fmt.Println("The log level you provided could not be found")
		return
	}
	if level == Info {
		logFlag = "[INFO]"
	}
	if level == Error {
		logFlag = "[ERROR]"
	}
	if level == Fatal {
		logFlag = "[FATAL]"
		exit = true
	}

	var logMsg string = fmt.Sprintf("%s %s", logFlag, message)
	if exit {
		log.Fatalln(logMsg)
	}
	log.Println(logMsg)
}
