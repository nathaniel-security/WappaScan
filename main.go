package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"

	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <url>")
		os.Exit(1)
	}
	url := os.Args[1]

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	wappalyzerClient, err := wappalyzer.New()
	if err != nil {
		log.Fatalf("Error initializing Wappalyzer: %v", err)
	}

	fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)

	// Extract technology names
	var techs []string
	for tech := range fingerprints {
		techs = append(techs, tech)
	}
	sort.Strings(techs)

	// Build final JSON object
	output := map[string]interface{}{
		"url":          url,
		"technologies": techs,
	}

	// Print JSON
	jsonOutput, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	fmt.Println(string(jsonOutput))
}
