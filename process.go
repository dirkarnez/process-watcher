package processwatcher

import (
	"fmt"
	"os"
)

func IsProcessRunning(processID int) (bool, error) {
	process, err := os.FindProcess(processID)
	if err != nil {
		return false, fmt.Errorf("no such process")
	}

	return isProcessRunning(process)
}
