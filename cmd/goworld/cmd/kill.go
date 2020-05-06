package cmd

import "syscall"

func Kill(sid ServerID) {
	StopWithSignal(sid, syscall.SIGKILL)
}
