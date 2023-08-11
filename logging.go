package main

import (
    "log"
    "os"
)

var (
    iLog, wLog, eLog *log.Logger
)

func init() {
    iLog = log.New(os.Stderr, " [ Shenk INF ] ", log.Ldate|log.Ltime)
    wLog = log.New(os.Stderr, " [ Shenk WRN ] ", log.Ldate|log.Ltime)
    eLog = log.New(os.Stderr, " [ Shenk ERR ] ", log.Ldate|log.Ltime)
}
