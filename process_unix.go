//go:build freebsd || linux || netbsd || openbsd || solaris || dragonfly
// +build freebsd linux netbsd openbsd solaris dragonfly

package processwatcher

import (
	"os"
	"syscall"
)

func isProcessRunning(p *os.Process) (bool, error) {
	// On Unix-like systems, use the 'kill' system call with a signal 0 to check if the process is running
	if err := p.Signal(syscall.Signal(0)); err != nil {
		if err == syscall.ESRCH {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
