package service

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/kardianos/osext"
	"github.com/op/go-logging"
)

var logger *logging.Logger

var formatLog = logging.MustStringFormatter(
	`%{color}%{level:.1s} %{time:2006-01-02 15:04:05.000} %{shortpkg} %{shortfile} ▶ %{color:reset} %{message}`,
)
var formatSys = logging.MustStringFormatter(
	`%{level:.1s} %{time:2006-01-02 15:04:05.000} %{shortfile} ▶ %{message}`,
)

func initLogger() {
	logger = logging.MustGetLogger(Config.Name)
	logger.ExtraCalldepth = 1
}

// consoleLogger -
func consoleLogger() {
	logging.SetBackend(
		logging.NewBackendFormatter(
			logging.NewLogBackend(os.Stderr, "", 0), formatLog,
		),
	)
}

// sysLogger -
func sysLogger() {
	bin, _ := osext.Executable()
	p, f := path.Split(bin)
	logDir := path.Join(p, "log")
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, os.ModePerm)
	}
	lf, err := os.OpenFile(path.Join(logDir, f+".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	b := logging.AddModuleLevel(logging.NewLogBackend(os.Stderr, "", 0))
	b.SetLevel(logging.ERROR, Config.Name)
	logging.SetBackend(
		logging.NewBackendFormatter(
			logging.NewLogBackend(lf, "", 0), formatLog,
		),
		logging.NewBackendFormatter(b, formatSys),
	)
}

// Debug -
func Debug(v ...interface{}) {
	logger.Debug(fmt.Sprint(v...))
}

// Debugln -
func Debugln(a ...interface{}) {
	logger.Debug(fmt.Sprintln(a...))
}

// Debugf -
func Debugf(format string, a ...interface{}) {
	logger.Debug(fmt.Sprintf(format, a...))
}

// Info -
func Info(v ...interface{}) {
	logger.Info(fmt.Sprint(v...))
}

// Infoln -
func Infoln(a ...interface{}) {
	logger.Info(fmt.Sprintln(a...))
}

// Infof -
func Infof(format string, a ...interface{}) {
	logger.Info(fmt.Sprintf(format, a...))
}

// Notice -
func Notice(v ...interface{}) {
	logger.Notice(fmt.Sprint(v...))
}

// Noticeln -
func Noticeln(a ...interface{}) {
	logger.Notice(fmt.Sprintln(a...))
}

// Noticef -
func Noticef(format string, a ...interface{}) {
	logger.Notice(fmt.Sprintf(format, a...))
}

// Warning -
func Warning(v ...interface{}) {
	logger.Warning(fmt.Sprint(v...))
}

// Warningln -
func Warningln(a ...interface{}) {
	logger.Warning(fmt.Sprintln(a...))
}

// Warningf -
func Warningf(format string, a ...interface{}) {
	logger.Warning(fmt.Sprintf(format, a...))
}

// Error -
func Error(v ...interface{}) {
	logger.Error(fmt.Sprint(v...))
}

// Errorln -
func Errorln(a ...interface{}) {
	logger.Error(fmt.Sprintln(a...))
}

// Errorf -
func Errorf(format string, a ...interface{}) {
	logger.Error(fmt.Sprintf(format, a...))
}

// Fatal -
func Fatal(v ...interface{}) {
	logger.Critical(fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalln -
func Fatalln(a ...interface{}) {
	logger.Critical(fmt.Sprintln(a...))
	os.Exit(1)
}

// Fatalf -
func Fatalf(format string, a ...interface{}) {
	logger.Critical(fmt.Sprintf(format, a...))
	os.Exit(1)
}
