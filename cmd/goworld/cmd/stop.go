package cmd

import (
	"os"
	"syscall"
	"time"

	"github.com/ouczbs/goworld/cmd/goworld/process"
)

func Stop(sid ServerID) {
	StopWithSignal(sid, StopSignal)
}

func StopWithSignal(sid ServerID, signal syscall.Signal) {
	err := os.Chdir(env.GoWorldRoot)
	checkErrorOrQuit(err, "chdir to goworld directory failed")

	ss := detectServerStatus()
	showServerStatus(ss)
	if !ss.IsRunning() {
		// server is not running
		ShowMsgAndQuit("no server is running currently")
	}

	if ss.ServerID != "" && ss.ServerID != sid {
		ShowMsgAndQuit("another server is running: %s", ss.ServerID)
	}

	StopGates(ss, signal)
	StopGames(ss, signal)
	StopDispatcher(ss, signal)
}

func StopGames(ss *ServerStatus, signal syscall.Signal) {
	if ss.NumGamesRunning == 0 {
		return
	}

	ShowMsg("Stop %d games ...", ss.NumGamesRunning)
	for _, proc := range ss.GameProcs {
		StopProc(proc, signal)
	}
}

func StopDispatcher(ss *ServerStatus, signal syscall.Signal) {
	if ss.NumDispatcherRunning == 0 {
		return
	}

	ShowMsg("Stop dispatcher ...")
	for _, proc := range ss.DispatcherProcs {
		StopProc(proc, signal)
	}
}

func StopGates(ss *ServerStatus, signal syscall.Signal) {
	if ss.NumGatesRunning == 0 {
		return
	}

	ShowMsg("Stop %d gates ...", ss.NumGatesRunning)
	for _, proc := range ss.GateProcs {
		StopProc(proc, signal)
	}
}

func StopProc(proc process.Process, signal syscall.Signal) {
	ShowMsg("Stop process %s pid=%d", proc.Executable(), proc.Pid())

	proc.Signal(signal)
	for {
		time.Sleep(time.Millisecond * 100)
		if !checkProcessRunning(proc) {
			break
		}
	}
}

func checkProcessRunning(proc process.Process) bool {
	pid := proc.Pid()
	procs, err := process.Processes()
	checkErrorOrQuit(err, "list processes failed")
	for _, _proc := range procs {
		if _proc.Pid() == pid {
			return true
		}
	}
	return false
}
