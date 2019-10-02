package service

import (
	"fmt"
	"os"
	"path"

	base "github.com/bendersilver/service"
	"github.com/joho/godotenv"
	"github.com/kardianos/osext"
)

// Config -
var Config base.Config

type program struct {
	BotService
}

// BotService -
type BotService interface {
	StartService()
	StopService()
	Init()
}

// Start -
func (p *program) Start(s base.Service) error {
	go p.StartService()
	return nil
}

// Stop -
func (p *program) Stop(s base.Service) error {
	p.StopService()
	return nil
}

// Init -
func Init(name, userName string) {
	Config.Name = name
	Config.DisplayName = name
	Config.Description = fmt.Sprintf("System servise %s", name)
	Config.UserName = userName
	Config.Dependencies = []string{
		"Requires=network.target",
		"After=network-online.target syslog.target"}
	if p, err := osext.ExecutableFolder(); err == nil {
		options := make(base.KeyValue)
		options["EnvFile"] = path.Join(p, ".env")
		Config.Option = options
	}
	// options["HasOutputFileSupport"] = true
}

// Run -
func Run(b BotService) {

	p := &program{b}
	s, err := base.New(p, &Config)
	if err != nil {
		Fatal(err)
	}
	args := os.Args
	p.Init()
	if len(args) > 1 {
		switch args[1] {
		case "start", "stop", "restart", "install", "uninstall":
			if args[1] == "uninstall" {
				base.Control(s, "stop")
			}
			err := base.Control(s, args[1])
			if err != nil {
				Fatal(err)
			}
			if args[1] == "install" {
				base.Control(s, "start")
			}
		default:
			Fatalf("Valid actions: %q\n", base.ControlAction)
		}
	} else {
		err = s.Run()
		if err != nil {
			Error(err)
		}
	}
}
func exists(name string) (exists bool) {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	return
}

func init() {
	initLogger()
	if base.Interactive() {
		consoleLogger()
		var f string
		p0, _ := osext.ExecutableFolder()
		p1, _ := os.Getwd()
		for _, i := range []string{p0, p1} {
			if exists(path.Join(i, ".env")) {
				f = path.Join(i, ".env")
				break
			}
		}
		if err := godotenv.Load(f); err != nil {
			Error("No .env file found")
		}

	} else {
		sysLogger()
	}
}
