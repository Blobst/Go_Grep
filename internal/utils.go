package internal

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
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

	return regexp.MustCompile("(?i)" + regexp.QuoteMeta(query))
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

	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
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
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	lineNum := 1
	found := false
	matcher := newInsensitiveMatcher(word)

	for scanner.Scan() {
		line := scanner.Text()
		if highlighted, ok := highlightFirstMatch(line, matcher); ok {
			fmt.Printf("%s:%d: %s\n", fileName, lineNum, highlighted)
			found = true
		}
		lineNum++
	}

	if !found {
		fmt.Println("No matches found in", fileName)
	}
}
