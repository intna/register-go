package middlerware

import (
	"log"
	"os"
)

var logger *log.Logger

func Log() *log.Logger{
    logger = log.New(os.Stdout, "[MyApp] ", log.LstdFlags|log.Lmicroseconds)
	return logger;
}