package main

import (
	"bufio"
	"os"
	"strings"
)

func parseCommand(input string) {
	switch {
	case input == "q":
		GoodbyeMessage()
		os.Exit(0)
	}
}

func askForCommand() string {
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	return strings.TrimSpace(response)
}

func executeCommand() {
	command := askForCommand()
	parseCommand(command)
}
