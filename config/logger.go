package config

import (
	"io"
	"log"
	"os"
)

func InitAppLogger(filename string) *log.Logger {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Unable to init logger. Err %s", err.Error())
	}

	w := io.MultiWriter(os.Stdout, f)

	return log.New(w, "GOTOS\t", log.LstdFlags)
}
