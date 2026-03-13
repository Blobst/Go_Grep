package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const version float64 = 1.0
const (
	ColorFound = "\x1b[32m"
	ColorReset = "\x1b[0m"
)

func findSimilarFiles(fileName string) {
	files, _ := os.ReadDir(".")
	found := false
	lowerInput := strings.ToLower(fileName)

	for _, f := range files {
		name := f.Name()
		lowerName := strings.ToLower(name)

		if strings.Contains(lowerName, lowerInput) {
			// Find start index of match
			start := strings.Index(lowerName, lowerInput)
			end := start + len(fileName)

			// Build highlighted string
			highlighted := name[:start] + ColorFound + name[start:end] + ColorReset + name[end:]
			fmt.Println("Found:", highlighted)
			found = true
		}
	}

	if !found {
		fmt.Println("No matching files")
	}
}

func main() {
	progIsRunning := true
	scanner := bufio.NewScanner(os.Stdin)
	cmd := NewCommands()
	re := regexp.MustCompile(`^[a-zA-Z0-9 _\-.]+$`)
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
				findSimilarFiles(scanner.Text())
			}
		}
	}
}
