package main

import (
	"fmt"
	"log"
	"runtime"
)

type ContextualError struct {
	err      error
	context  string
	file     string
	line     int
	funcName string
}

func (e ContextualError) Error() string {
	return fmt.Sprintf("%s:%d: %s: %s", e.file, e.line, e.funcName, e.context)
}

func WrapError(err error, context string) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "<unknown file>"
		line = 0
	}
	funcName := runtime.FuncForPC(runtime.Caller(1)).Name()

	return ContextualError{
		err:      err,
		context:  context,
		file:     file,
		line:     line,
		funcName: funcName,
	}
}

func main() {
	err := fmt.Errorf("error opening file")
	logError(WrapError(err, "while reading config file"))
}

func logError(err error) {
	log.Println("Error:", err)
}
