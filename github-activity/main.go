package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Repo struct {
	Name string `json:"name"`
}

type Event struct {
	Type string `json:"type"`
	Repo Repo   `json:"repo"`
}

func printBorder() {
	fmt.Println("+-----+----------------------+--------------------------------+")
}

func describeEvent(event Event) string {
	switch event.Type {

	case "PushEvent":
		return "Pushed commits"

	case "WatchEvent":
		return "Starred"

	case "IssuesEvent":
		return "Opened issue"

	case "ForkEvent":
		return "Forked repository"

	case "CreateEvent":
		return "Created repository"

	default:
		return event.Type
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		os.Exit(1)
	}

	username := os.Args[1]

	fmt.Println("Fetching GitHub user activity for", username)

	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Could not fetch user activity for %s\n", username)
		os.Exit(1)
	}

	var events []Event

	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	if len(events) == 0 {
		fmt.Println("No recent activity found.")
		return
	}

	// if len(events) > 10 {
	// 	events = events[:10]
	// }

	printBorder()
	fmt.Printf("| %-3s | %-20s | %-30s |\n", "#", "Event", "Repository")
	printBorder()

	for i, event := range events {
		description := describeEvent(event)
		fmt.Printf("| %-3d | %-20s | %-30s |\n", i+1, description, event.Repo.Name)
	}

	printBorder()
}
