package inits

import (
	"log"
	"os"
)

type Logging struct {
	Logger *log.Logger
}

func NewLogger() *Logging {
	logger := log.New(os.Stdout, "GO-WEATHER-APP... ", log.Ldate|log.Ltime|log.Lshortfile)
	return &Logging{
		Logger: logger,
	}
}
