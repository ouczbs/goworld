package main

import (
	"github.com/ouczbs/goworld/components/dispatcher"
	"github.com/ouczbs/goworld/engine/config"
	"os"
	"strconv"
)

func main() {
	dispatcherIds := config.GetDispatcherIDs()
	for _, dispid := range dispatcherIds {
		var cmdList []string
		cmdList = append(cmdList, "cmd")
		cmdList = append(cmdList, "-dispid")
		cmdList = append(cmdList, strconv.Itoa(int(dispid)))
		os.Args = cmdList
		dispatcher.RunDispatvher()
		break
	}
}