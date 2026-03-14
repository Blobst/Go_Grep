package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const Version = "1.0.1"

type Release struct {
	TagName string `json:"tag_name"`
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

	latest := release.TagName

	if latest != Version {
		fmt.Println("New version available:", latest)
		fmt.Println("Current version:", Version)
	} else {
		fmt.Println("You are using the latest version.")
	}
}
