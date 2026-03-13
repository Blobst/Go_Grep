package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

const version float64 = 1.0
const (
	ColorFound = "\x1b[32m"
	ColorReset = "\x1b[0m"
)

type Commands struct {
	CommandExit     string
	CommandHelp     string
	CommandVersion  string
	CommandPrevious string
}

func NewCommands() *Commands {
	return &Commands{
		CommandExit:     "{exit",
		CommandHelp:     "{help",
		CommandVersion:  "{ver",
		CommandPrevious: "{pre",
	}
}

func fileExists(fileName string) {
	info, err := os.Stat(fileName)
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	fmt.Printf("File Found: %v%s%v\n", ColorFound, info.Name(), ColorReset)
}

func main() {
	progIsRunning := true
	scanner := bufio.NewScanner(os.Stdin)
	cmd := NewCommands()
	re := regexp.MustCompile(`^[a-zA-Z0-9 _-]+$`)
	fmt.Printf("go_grep v[%.1f]\n", version)

	for progIsRunning {
		fmt.Print("> ")
		scanner.Scan()

		switch scanner.Text() {
		case cmd.CommandExit:
			fmt.Println("Exiting...")
			progIsRunning = false
		case cmd.CommandHelp:
			fmt.Print("Commands: {help, {exit, {ver\n")
		case cmd.CommandVersion:
			fmt.Printf("go_grep v[%.1f]\n", version)

		default:
			if re.MatchString(scanner.Text()) {
				fileExists(scanner.Text())
			}
		}
	}
}
