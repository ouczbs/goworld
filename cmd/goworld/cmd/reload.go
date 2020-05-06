package cmd

import (
	"os"

	"github.com/ouczbs/goworld/engine/binutil"
	"github.com/ouczbs/goworld/engine/config"
)

func Reload(sid ServerID) {
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

	if ss.NumGamesRunning == 0 {
		ShowMsgAndQuit("no game is running")
	} else if ss.NumGamesRunning != config.GetDeployment().DesiredGames {
		ShowMsgAndQuit("found %d games, but should have %d", ss.NumGamesRunning, config.GetDeployment().DesiredGames)
	}

	StopGames(ss, binutil.FreezeSignal)
	StartGames(sid, true)
}
