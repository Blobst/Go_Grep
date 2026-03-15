package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const Version = "1.0.4"

type Release struct {
	TagName string `json:"tag_name"`
	Body    string `json:"body"`
	URL     string `json:"html_url"`
}

func isNewerVersion(latest, current string) bool {
	latestParts := strings.Split(latest, ".")
	currentParts := strings.Split(current, ".")

	for i := 0; i < len(latestParts) && i < len(currentParts); i++ {
		l, _ := strconv.Atoi(latestParts[i])
		c, _ := strconv.Atoi(currentParts[i])

		if l > c {
			return true
		}
		if l < c {
			return false
		}
	}

	return false
}

func CheckForUpdates() {
	url := "https://api.github.com/repos/Blobst/Go_Grep/releases/latest"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Update check failed:", err)
		return
	}
	defer resp.Body.Close()

	var release Release

	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		fmt.Println("Failed to parse response:", err)
		return
	}

	latest := strings.TrimPrefix(release.TagName, "v")

	if isNewerVersion(latest, Version) {
		fmt.Println("New version available:", latest)
		fmt.Println("Current version:", Version)
		fmt.Println("Release notes:\n", release.Body)
		fmt.Println("Download here:", release.URL)
	} else {
		fmt.Println("You are using the latest version.")
	}
}
