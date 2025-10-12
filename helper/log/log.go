package log

import (
	"fmt"
	"log"
)

func Info(message any) {
	var logFlag string = "[INFO]"

	var logMsg string = fmt.Sprintf("%s %v", logFlag, message)
	log.Println(logMsg)
}
func Error(message any) {
	var logFlag string = "[ERROR]"

	var logMsg string = fmt.Sprintf("%s %v", logFlag, message)
	log.Println(logMsg)
}
func Fatal(message any) {
	var logFlag string = "[FATAL]"
	var exit bool = true

	var logMsg string = fmt.Sprintf("%s %v", logFlag, message)
	if exit {
		log.Fatalln(logMsg)
	}
}
