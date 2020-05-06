package cmd

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/ouczbs/goworld/engine/config"
	"github.com/ouczbs/goworld/engine/consts"
)
var arguments struct {
	runInDaemonMode bool
}
func Start(sid ServerID) {
	err := os.Chdir(env.GoWorldRoot)
	checkErrorOrQuit(err, "chdir to goworld directory failed")

	ss := detectServerStatus()
	if ss.NumDispatcherRunning > 0 || ss.NumGatesRunning > 0 {
		Status()
		ShowMsgAndQuit("server is already running, can not Start multiple servers")
	}

	StartDispatchers()
	//StartGames(sid, false)
	StartGates()
}

func StartDispatchers() {
	ShowMsg("Start dispatchers ...")
	dispatcherIds := config.GetDispatcherIDs()
	ShowMsg("dispatcher ids: %v", dispatcherIds)
	for _, dispid := range dispatcherIds {
		StartDispatcher(dispid)
	}
}

func StartDispatcher(dispid uint16) {
	cfg := config.GetDispatcher(dispid)
	args := []string{"-dispid", strconv.Itoa(int(dispid))}
	if arguments.runInDaemonMode {
		args = append(args, "-d")
	}
	cmd := exec.Command(env.GetDispatcherBinary(), args...)
	err := runCmdUntilTag(cmd, cfg.LogFile, consts.DISPATCHER_STARTED_TAG, time.Second*10)
	checkErrorOrQuit(err, "Start dispatcher failed, see dispatcher.log for error")
}

func StartGames(sid ServerID, isRestore bool) {
	ShowMsg("Start games ...")
	desiredGames := config.GetDeployment().DesiredGames
	ShowMsg("desired games = %d", desiredGames)
	for gameid := uint16(1); int(gameid) <= desiredGames; gameid++ {
		StartGame(sid, gameid, isRestore)
	}
}

func StartGame(sid ServerID, gameid uint16, isRestore bool) {
	ShowMsg("Start game %d ...", gameid)

	gameExePath := filepath.Join(sid.Path(), sid.Name()+BinaryExtension)
	args := []string{"-gid", strconv.Itoa(int(gameid))}
	if isRestore {
		args = append(args, "-restore")
	}
	if arguments.runInDaemonMode {
		args = append(args, "-d")
	}
	cmd := exec.Command(gameExePath, args...)
	err := runCmdUntilTag(cmd, config.GetGame(gameid).LogFile, consts.GAME_STARTED_TAG, time.Second*600)
	checkErrorOrQuit(err, "Start game failed, see game.log for error")
}

func StartGates() {
	ShowMsg("Start gates ...")
	desiredGates := config.GetDeployment().DesiredGates
	ShowMsg("desired gates = %d", desiredGates)
	for gateid := uint16(1); int(gateid) <= desiredGates; gateid++ {
		StartGate(gateid)
	}
}

func StartGate(gateid uint16) {
	ShowMsg("Start gate %d ...", gateid)

	args := []string{"-gid", strconv.Itoa(int(gateid))}
	if arguments.runInDaemonMode {
		args = append(args, "-d")
	}
	cmd := exec.Command(env.GetGateBinary(), args...)
	err := runCmdUntilTag(cmd, config.GetGate(gateid).LogFile, consts.GATE_STARTED_TAG, time.Second*10)
	checkErrorOrQuit(err, "Start gate failed, see gate.log for error")
}

func runCmdUntilTag(cmd *exec.Cmd, logFile string, tag string, timeout time.Duration) (err error) {
	clearLogFile(logFile)
	err = cmd.Start()
	if err != nil {
		return
	}

	timeoutTime := time.Now().Add(timeout)
	for time.Now().Before(timeoutTime) {
		time.Sleep(time.Millisecond * 200)
		if isTagInFile(logFile, tag) {
			cmd.Process.Release()
			return
		}
	}

	err = errors.Errorf("wait Started tag timeout")
	return
}

func clearLogFile(logFile string) {
	ioutil.WriteFile(logFile, []byte{}, 0644)
}

func isTagInFile(filename string, tag string) bool {
	data, err := ioutil.ReadFile(filename)
	checkErrorOrQuit(err, "read file error")
	return strings.Contains(string(data), tag)
}
