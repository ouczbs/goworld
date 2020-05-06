package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
)

func Build(sid ServerID) {
	ShowMsg("Building server %s ...", sid)

	BuildServer(sid)
	BuildDispatcher()
	BuildGate()
}

func BuildServer(sid ServerID) {
	serverPath := sid.Path()
	ShowMsg("server directory is %s ...", serverPath)
	if !isdir(serverPath) {
		ShowMsgAndQuit("wrong server id: %s, using '\\' instead of '/'?", sid)
	}

	ShowMsg("go Build %s ...", sid)
	BuildDirectory(serverPath)
}

func BuildDispatcher() {
	ShowMsg("go Build dispatcher ...")
	BuildDirectory(filepath.Join(env.GoWorldRoot, "components", "dispatcher"))
}

func BuildGate() {
	ShowMsg("go Build gate ...")
	BuildDirectory(filepath.Join(env.GoWorldRoot, "components", "gate"))
}

func BuildDirectory(dir string) {
	var err error
	var curdir string
	curdir, err = os.Getwd()
	checkErrorOrQuit(err, "")

	err = os.Chdir(dir)
	checkErrorOrQuit(err, "")

	defer os.Chdir(curdir)

	cmd := exec.Command("go", "Build", ".")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err = cmd.Run()
	checkErrorOrQuit(err, "")
	return
}
