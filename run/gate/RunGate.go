package main

import (
	"github.com/ouczbs/goworld/components/gate"
	"github.com/ouczbs/goworld/engine/config"
	"os"
	"strconv"
)

func main() {
	desiredGates := config.GetDeployment().DesiredGates
	for gateid := uint16(1); int(gateid) <= desiredGates; gateid++ {
		var cmdList []string
		cmdList = append(cmdList, "cmd")
		cmdList = append(cmdList, "-gid")
		cmdList = append(cmdList, strconv.Itoa(int(gateid)))
		os.Args = cmdList
		gate.RunGate()
		break
	}
}
