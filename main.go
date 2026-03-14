package main

import (
	"bufio"
	"fmt"
	ut "gogrep/internal"
	"os"
	"regexp"
)

func main() {
	progIsRunning := true
	scanner := bufio.NewScanner(os.Stdin)
	cmd := ut.NewCommands()
	re := regexp.MustCompile(`^[a-zA-Z0-9 _\-.]+$`)
	fmt.Printf("go_grep v[%s]\n", ut.Version)

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
			fmt.Printf("go_grep v[%s]\n", ut.Version)
		case cmd.CommandUpdateCheck:
			ut.CheckForUpdates()

		default:
			if re.MatchString(scanner.Text()) {
				ut.FindSimilarFiles(scanner.Text())
				files, _ := os.ReadDir(".")
				for _, f := range files {
					if !f.IsDir() {
						ut.FindWordSimilar(f.Name(), scanner.Text())
					}
				}
			}
		}
	}
}
