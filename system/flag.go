package system

import (
	"flag"
	"fmt"
	"os"
)

//
var (
	Flag  func(*flag.FlagSet)
	Usage string
)

//Icon Icon
var Icon = `
         __   __               _____    ____  
         \ \ / /              |_   _|  / __ \ 
          \ V /     ______      | |   | |  | |
           > <     |______|     | |   | |  | |
          / . \                _| |_  | |__| |
         /_/ \_\              |_____|  \____/ 

`

//UsageTemplate UsageTemplate
var UsageTemplate = `
%s
Usage: %s [options]
%s
Logging Options:
	-l, --log <file>                 File to redirect log output

Common Options:
	-c, --config <file>              Configuration file	
	-h, --help                       Show this message
	-v, --version                    Show version
`

//Bind Bind
func Bind(config interface{}) error {
	fs := flag.NewFlagSet("main", flag.ExitOnError)

	fs.Usage = func() {
		fmt.Printf(UsageTemplate, Icon, Name, Usage)
		os.Exit(0)
	}

	var (
		showVersion bool
		configFile  string
		logDir      string
	)

	fs.BoolVar(&showVersion, "v", false, "Show version")
	fs.BoolVar(&showVersion, "version", false, "Show version")
	fs.StringVar(&logDir, "l", "", "File to redirect log output")
	fs.StringVar(&logDir, "log", "", "File to redirect log output")
	fs.StringVar(&configFile, "c", "", "Configuration file.")
	fs.StringVar(&configFile, "config", "", "Configuration file.")

	if Flag != nil {
		Flag(fs)
	}

	if err := fs.Parse(os.Args[1:]); err != nil {
		return err
	}

	if showVersion {
		version()
		os.Exit(0)
		return nil
	}

	if logDir != "" {
		setLogOutput(logDir)
	}

	// Parse config if given
	if configFile == "" {
		configFile = "config.yaml"
	}
	err := LoadConfig(config, configFile, "/")
	if err != nil {
		return err
	}

	return fs.Parse(os.Args[1:])
}
