package logger

import (
	"fmt"
	"log"
)

func Info(msg string) {
	log.Panic(msg)
}

func FailOnError(err error, msg string) {
	if err != nil {
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
