package internal

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	ColorPath  = "\x1b[32m"
	ColorLine  = "\x1b[34m"
	ColorMatch = "\x1b[31m"
	ColorReset = "\x1b[0m"
)

func newInsensitiveMatcher(query string) (*regexp.Regexp, error) {
	if strings.TrimSpace(query) == "" {
		return nil, errors.New("search pattern cannot be empty")
	}

	return regexp.Compile(`(?i)` + query)
}

func colorize(text, color string) string {
	if os.Getenv("NO_COLOR") != "" {
		return text
	}
	return color + text + ColorReset
}

func highlightAllMatches(text string, matcher *regexp.Regexp) (string, bool) {
	if matcher == nil {
		return text, false
	}

	locs := matcher.FindAllStringIndex(text, -1)
	if len(locs) == 0 {
		return text, false
	}

	var b strings.Builder
	last := 0
	for _, loc := range locs {
		start, end := loc[0], loc[1]
		if start == end {
			continue
		}
		b.WriteString(text[last:start])
		b.WriteString(ColorMatch)
		b.WriteString(text[start:end])
		b.WriteString(ColorReset)
		last = end
	}
	b.WriteString(text[last:])

	return b.String(), true
}

func findWordSimilarInFile(displayName string, filePath string, matcher *regexp.Regexp, hasPrintedAnyFile *bool) bool {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", filePath)
		return false
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", filePath)
		}
	}()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	foundInThisFile := false

	for scanner.Scan() {
		line := scanner.Text()
		highlighted, ok := highlightAllMatches(line, matcher)
		if ok {
			if !foundInThisFile {
				if *hasPrintedAnyFile {
					fmt.Println()
				}
				fmt.Println(colorize(displayName, ColorPath))
				*hasPrintedAnyFile = true
				foundInThisFile = true
			}

			fmt.Printf("%s:%s\n", colorize(strconv.Itoa(lineNum), ColorLine), highlighted)
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", filePath)
	}

	return foundInThisFile
}

func findWordSimilar(root string, matcher *regexp.Regexp) {
	hasPrintedAnyFile := false

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		displayPath := path
		if rel, relErr := filepath.Rel(root, path); relErr == nil {
			displayPath = rel
		}

		findWordSimilarInFile(displayPath, path, matcher, &hasPrintedAnyFile)
		return nil
	})
	if err != nil {
		fmt.Println("Error searching files:", err)
	}
}

func Search(query string) {
	matcher, err := newInsensitiveMatcher(query)
	if err != nil {
		fmt.Println("Invalid regular expression:", err)
		return
	}

	findWordSimilar(".", matcher)
}
