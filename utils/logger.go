package utils

import (
	"os"
	"log"
)

const logFile = "weblog.log"

func WebLog(s string) {
	// Open log file
	logF, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logF.Close()
	log.SetOutput(logF)

	// var buf bytes.Buffer
	// logger := log.New(&buf, "logger: ", log.Lshortfile)
	log.Println(s)
}
