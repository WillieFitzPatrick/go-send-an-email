package logger

import (
  "path/filepath"
  "fmt"
  "os"
  "log"
)

var (
  // Log pointer
  Log *log.Logger
)


func init() {

  logFileName := filepath.Base(os.Args[0])
  logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

  if err != nil {
    fmt.Printf("error opening log file (%s) : %s ", logFileName, err.Error())
  }
  Log = log.New(logFile, "", log.LstdFlags)

}