package main

import (
	"github.com/alimy/mir-covid19/cmd"
)

func main() {
	cmd.Setup(
		"covid",                       // command name
		"covid19 information service", // command short describe
		"covid19 information service", // command long describe
	)
	// execute start application
	cmd.Execute()
}
