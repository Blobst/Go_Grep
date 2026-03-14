package main

import (
	"bufio"
	"fmt"
	ut "gogrep/internal"
	"os"
)

func main() {
	progIsRunning := true
	scanner := bufio.NewScanner(os.Stdin)
	cmd := ut.NewCommands()
	fmt.Printf("go_grep v[%s]\n", ut.Version)

	for progIsRunning {
		fmt.Print(":> ")
		scanner.Scan()

		switch scanner.Text() {
		case cmd.CommandExit:
			fmt.Println("Exiting...")
			progIsRunning = false
		case cmd.CommandHelp:
			fmt.Print("Commands: {help, {exit, {ver, {upd\n")
		case cmd.CommandVersion:
			fmt.Printf("go_grep v[%s]\n", ut.Version)
		case cmd.CommandUpdateCheck:
			ut.CheckForUpdates()

		default:
			ut.Search(scanner.Text())
		}
	}
}
