package main

import (
	"fmt"
	"log"
	"os"
)

// LOGFILE conains the path to our custom log file
var LOGFILE = "/tmp/mGo.log"

func main() {
	// os.O_CREATE creates the file as it doesn't exist,
	// os.O_wRONLY makes the file write only
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "customLogLineNumber ", log.LstdFlags)

	// to print line numbers in log entries
	iLog.SetFlags(log.LstdFlags | log.Lshortfile)

	iLog.Println("Hello there!")
	iLog.Println("Another log entry")
}
