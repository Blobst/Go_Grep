package internal

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

const (
	ColorFound = "\x1b[32m"
	ColorReset = "\x1b[0m"
)

func newInsensitiveMatcher(query string) *regexp.Regexp {
	if query == "" {
		return nil
	}

	return regexp.MustCompile(`(?i)` + regexp.QuoteMeta(query))
}

func highlightFirstMatch(text string, matcher *regexp.Regexp) (string, bool) {
	if matcher == nil {
		return text, false
	}

	loc := matcher.FindStringIndex(text)
	if loc == nil {
		return text, false
	}

	start, end := loc[0], loc[1]
	return text[:start] + ColorFound + text[start:end] + ColorReset + text[end:], true
}

func FindSimilarFiles(fileName string) {
	found := false
	matcher := newInsensitiveMatcher(fileName)

	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		name := d.Name()

		if highlighted, ok := highlightFirstMatch(name, matcher); ok {
			fmt.Println("Found:", filepath.Join(filepath.Dir(path), highlighted))
			found = true
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error searching files:", err)
	}

	if !found {
		fmt.Println("No matching files")
	}
}

func findWordSimilarInFile(fileName string, matcher *regexp.Regexp) bool {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", fileName)
		return false
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", fileName)
		}
	}()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		if highlighted, ok := highlightFirstMatch(line, matcher); ok {
			fmt.Printf("%s:%d: %s\n", fileName, lineNum, highlighted)
			found = true
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", fileName)
	}

	return found
}

func FindWordSimilar(root string, word string) {
	found := false
	matcher := newInsensitiveMatcher(word)

	if matcher == nil {
		fmt.Println("No matches found")
		return
	}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.IsDir() {
			return nil
		}

		if findWordSimilarInFile(path, matcher) {
			found = true
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error searching files:", err)
	}

	if !found {
		fmt.Println("No matches found")
	}
}
