package main

import (
	"errors"
	"fmt"
	"os"
)

const version float64 = 1.0
const (
	ColorFound = "\x1b[31m"
	ColorReset = "\x1b[0m"
)

func fileExits(fileName string) {
	info, err := os.Stat(fileName)
	if errors.Is(err, os.ErrNotExist) {
		return
	}
	fmt.Printf("File Found: %v%s%v", ColorFound, info.Name(), ColorReset)
}

func main() {
	fmt.Printf("Go-Grep v[%.1f]", version)
	fileExits("main.go")
}
