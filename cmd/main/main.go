package main

import (
	"fmt"

	"github.com/dirkarnez/processwatcher"
)

func main() {
	isRunning, err := processwatcher.IsProcessRunning(47584)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(isRunning)
	}
}
