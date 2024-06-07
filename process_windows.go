//go:build windows
// +build windows

package processwatcher

import (
	"fmt"
	"os"
	"syscall"
)

func isProcessRunning(p *os.Process) (bool, error) {
	// On Windows, use the 'GetExitCodeProcess' Windows API function to check if the process is running
	var exitCode uint32
	err := syscall.GetExitCodeProcess(syscall.Handle(p.Pid), &exitCode)
	if err != nil {
		return false, fmt.Errorf("cannot run GetExitCodeProcess")
	} else {
		return exitCode != 0, nil
	}
}
