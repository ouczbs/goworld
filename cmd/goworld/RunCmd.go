package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	cmdlib "github.com/ouczbs/goworld/cmd/goworld/cmd"
)

var arguments struct {
	runInDaemonMode bool
}

func parseArgs() {
	//flag.StringVar(&arguments.configFile, "configfile", "", "set config file path")
	flag.BoolVar(&arguments.runInDaemonMode, "d", false, "run in daemon mode")
	flag.Parse()
}

func main() {
	runcmd := "cmd -d true Start ./goworld"
	os.Args = strings.Split(runcmd, " ")
	parseArgs()
	args := flag.Args()
	cmdlib.ShowMsg("arguments: %s", strings.Join(args, " "))

	cmdlib.DetectGoWorldPath()

	if len(args) == 0 {
		cmdlib.ShowMsg("no command to execute")
		flag.Usage()
		fmt.Fprintf(os.Stderr, "\tgoworld <Build|Start|Stop|Kill|Reload|Status> [server-id]\n")
		os.Exit(1)
	}

	cmd := args[1]

	if cmd == "Build" || cmd == "Start" || cmd == "Stop" || cmd == "Reload" || cmd == "Kill" {
		if len(args) != 3 {
			cmdlib.ShowMsgAndQuit("server id is not given")
		}
	}
	if cmd == "Build" {
		cmdlib.Build(cmdlib.ServerID(args[2]))
	} else if cmd == "Start" {
		cmdlib.Start(cmdlib.ServerID(args[2]))
	} else if cmd == "Stop" {
		if runtime.GOOS == "windows" {
			cmdlib.ShowMsgAndQuit("Stop does not work on Windows, use Kill instead (will lose player data)")
		}

		cmdlib.Stop(cmdlib.ServerID(args[2]))
	} else if cmd == "Reload" {
		cmdlib.Reload(cmdlib.ServerID(args[2]))
	} else if cmd == "Kill" {
		cmdlib.Kill(cmdlib.ServerID(args[2]))
	} else if cmd == "Status" {
		cmdlib.Status()
	} else {
		cmdlib.ShowMsgAndQuit("unknown command: %s", cmd)
	}
}
