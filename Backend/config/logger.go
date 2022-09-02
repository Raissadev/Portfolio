package config

import (
	"io"
	"log/syslog"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	lSyslog "github.com/sirupsen/logrus/hooks/syslog"
)

type LoggerInterface interface {
	Sys() *Logger
	New() *Logger
}

type Logger struct {
	//
}

func (l *Logger) Sys() *Logger {
	log := logrus.New()
	hook, err := lSyslog.NewSyslogHook("udp", "localhost:"+os.Getenv("APP_PORT"), syslog.LOG_INFO, "")

	if err == nil {
		log.Hooks.Add(hook)
	}

	return l
}

func (l *Logger) New() *Logger {
	log.SetReportCaller(true)

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.SetLevel(log.WarnLevel)

	path, _ := filepath.Abs("app/log/log.txt")

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	wrt := io.MultiWriter(os.Stdout, f)

	log.SetOutput(wrt)

	defer f.Close()

	return l
}
