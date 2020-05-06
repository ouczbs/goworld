package cmd

import (
	"fmt"
	"os"
)

func ShowMsgAndQuit(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "! "+format+"\n", a...)
	os.Exit(2)
}

func ShowMsg(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "> "+format+"\n", a...)
}

func checkErrorOrQuit(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "! %s: %v\n", msg, err)
		os.Exit(2)
	}
}
