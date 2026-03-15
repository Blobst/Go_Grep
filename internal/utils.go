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

func newMatcher(query string, ignoreCase bool) (*regexp.Regexp, error) {
	if strings.TrimSpace(query) == "" {
		return nil, errors.New("search pattern cannot be empty")
	}

	if ignoreCase {
		query = `(?i)` + query
	}

	return regexp.Compile(query)
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
		b.WriteString(colorize(text[start:end], ColorMatch))
		last = end
	}
	b.WriteString(text[last:])

	return b.String(), true
}

func findWordSimilarInFile(displayName string, filePath string, matcher *regexp.Regexp) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, fmt.Errorf("failed to open file %q: %w", filePath, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to close file %q: %v\n", filePath, err)
		}
	}()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	foundInThisFile := false

	for scanner.Scan() {
		line := scanner.Text()
		highlighted, ok := highlightAllMatches(line, matcher)
		if ok {
			foundInThisFile = true
			fmt.Printf("%s:%s:%s\n", colorize(displayName, ColorPath), colorize(strconv.Itoa(lineNum), ColorLine), highlighted)
		}
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		return foundInThisFile, fmt.Errorf("failed while reading file %q: %w", filePath, err)
	}

	return foundInThisFile, nil
}

func findWordSimilar(root string, matcher *regexp.Regexp) (bool, error) {
	foundAny := false

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

		foundInFile, fileErr := findWordSimilarInFile(displayPath, path, matcher)
		if fileErr != nil {
			_, _ = fmt.Fprintln(os.Stderr, fileErr)
		}
		if foundInFile {
			foundAny = true
		}
		return nil
	})
	if err != nil {
		return foundAny, fmt.Errorf("error searching files: %w", err)
	}

	return foundAny, nil
}

func Search(query string, path string, ignoreCase bool) (bool, error) {
	matcher, err := newMatcher(query, ignoreCase)
	if err != nil {
		return false, fmt.Errorf("invalid regular expression: %w", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		return false, fmt.Errorf("path %q is not accessible: %w", path, err)
	}

	if !info.IsDir() {
		return findWordSimilarInFile(path, path, matcher)
	}

	return findWordSimilar(path, matcher)
}
