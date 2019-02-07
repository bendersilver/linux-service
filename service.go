package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/logger"
	"github.com/kardianos/osext"
	base "github.com/kardianos/service"
)

var prog *program

// program -
type program struct {
	cnf  *base.Config
	ifce Ifce
}

func (p *program) setName(name string) {
	p.cnf.Name = name
}

// SetName -
func SetName(name string) {
	prog.setName(name)
	prog.setDisplayName(name)
	prog.setDescription(fmt.Sprintf("System servise %s", name))
}

// SetUserName -
func SetUserName(name string) {
	prog.setUserName(name)
}

func (p *program) setDisplayName(name string) {
	p.cnf.DisplayName = name
}

func (p *program) setDescription(txt string) {
	p.cnf.Description = txt
}

func (p *program) setUserName(name string) {
	p.cnf.UserName = name
}

// Ifce -
type Ifce interface {
	Start()
	Stop()
}

// Start -
func (p *program) Start(s base.Service) error {
	go p.ifce.Start()
	return nil
}

// Stop -
func (p *program) Stop(s base.Service) error {
	p.ifce.Stop()
	return nil
}

func setLogger() {
	if base.Interactive() {
		logger.Init(prog.cnf.Name, true, false, ioutil.Discard)
		logger.SetFlags(log.LstdFlags | log.Llongfile)
	} else {
		bin, _ := osext.Executable()
		logPath := fmt.Sprintf("%s.log", bin)
		lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
		if err != nil {
			logger.Fatalf("Failed to open log file: %v", err)
		}
		logger.Init(prog.cnf.Name, false, true, lf)
		logger.SetFlags(log.LstdFlags)
	}
}

// Run -
func Run(i Ifce) {
	setLogger()
	prog.ifce = i
	s, err := base.New(prog, prog.cnf)
	if err != nil {
		logger.Fatal(err)
	}
	errs := make(chan error, 5)

	go func() {
		for {
			err := <-errs
			if err != nil {
				logger.Error(err)
			}
		}
	}()

	args := os.Args
	if len(args) > 1 {
		arg := args[1]
		if arg == "start" || arg == "stop" || arg == "restart" || arg == "install" || arg == "uninstall" {
			err := base.Control(s, arg)
			if err != nil {
				logger.Fatal(err)
			}
		} else {
			logger.Fatalf("Valid actions: %q\n", base.ControlAction)
		}
	} else {
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	}
}

func init() {
	prog = &program{
		cnf: new(base.Config),
	}
}
