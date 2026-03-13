package internal

import (
	"fmt"
	"os"
	"strings"
)

const (
	ColorFound = "\x1b[32m"
	ColorReset = "\x1b[0m"
)

func FindSimilarFiles(fileName string) {
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

func FindWordSimilar(word string) {

}
