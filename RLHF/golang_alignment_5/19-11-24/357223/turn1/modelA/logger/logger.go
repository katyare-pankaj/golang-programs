package logger

import (
	"fmt"
	"log"
)

func LogMessage(msg string, args ...interface{}) {
	log.Println(fmt.Sprintf(msg, args...))
}
