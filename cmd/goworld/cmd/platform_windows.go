// +Build windows

package cmd

import (
	"syscall"

	_ "github.com/go-ole/go-ole" // so that dep can resolve versions correctly
)

const (
	// BinaryExtension extension used on windows
	BinaryExtension = ".exe"
	// StopSignal syscall used to Stop server
	StopSignal = syscall.SIGKILL
)
