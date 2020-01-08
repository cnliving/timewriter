package main

import (
	"time"
	"log"
	. "github.com/cnliving/timewriter"
)
const(
	TimeLayout string ="2006-01-02 15:04:05"
)

func main() {
	log.SetFlags(log.Lshortfile|log.Ldate|log.Ltime)
	timeWriter := &TimeWriter{
		Dir:        "./log",
		Compress:   true,
		ReserveDay: 30,
	}
	timeWriter.Init()
	logDebug := log.New(timeWriter, " [Debug] ", log.Lshortfile|log.Ldate|log.Ltime)

	for {
		tmCur:=time.Now()
		logDebug.Printf("this is debug %s \n",tmCur.Format(TimeLayout))
		log.Printf("this is debug %s\n",tmCur.Format(TimeLayout))
		time.Sleep(time.Duration(1)*time.Second)
	}
}