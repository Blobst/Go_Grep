package internal

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const Version float64 = 1.0

func CheckForUpdates() {
	updatehistory, err := os.ReadFile(".updatehist")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Checking github for updates...\n")

	converthistory, err := strconv.ParseFloat(string(updatehistory), 64)
	if err != nil {
		log.Fatal(err)
	}

	if converthistory < Version {
		fmt.Printf("A new version of go_grep is available! (v[%.1f])\n", Version)
		fmt.Println("Please visit the GitHub repository to download the latest version.")
	}
}
