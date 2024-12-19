package middleware

import (
	"log"
	"os"
)

var logger *log.Logger

func Log() *log.Logger {
	logger = log.New(os.Stdout, "[Server]", log.LstdFlags|log.Lmicroseconds)
	return logger
}
