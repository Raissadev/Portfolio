package logger

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

var (
	Log *log.Logger
)

func init() {
	var logpath, _ = filepath.Abs("app/log/log.txt")

	flag.Parse()
	var file, err = os.Open(logpath)

	if err != nil {
		panic(err)
	}

	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logpath)

	defer file.Close()
}
