package internal

import (
	"bufio"
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

func FindWordSimilar(fileName string, word string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", fileName)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	lineNum := 1
	wordLower := strings.ToLower(word)
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), wordLower) {
			// Highlight the matching word(s)
			highlighted := strings.ReplaceAll(line, word, ColorFound+word+ColorReset)
			fmt.Printf("%s:%d: %s\n", fileName, lineNum, highlighted)
			found = true
		}
		lineNum++
	}

	if !found {
		fmt.Println("No matches found in", fileName)
	}
}
