package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get application name from user
	fmt.Print("Enter application name (e.g., blog, store, dashboard): ")
	appName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Clean up the input
	appName = strings.TrimSpace(appName)
	if appName == "" {
		fmt.Println("Application name cannot be empty")
		os.Exit(1)
	}

	// Generate secure random bytes
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Printf("Error generating random bytes: %v\n", err)
		os.Exit(1)
	}

	// Create the API key
	randomString := strings.TrimRight(base64.URLEncoding.EncodeToString(bytes), "=")
	apiKey := fmt.Sprintf("sk_%s_%s", strings.ToLower(appName), randomString)

	// Display the result
	fmt.Printf("\n✓ Generated API key for '%s':\n", appName)
	fmt.Printf("%s\n", apiKey)
	fmt.Println("\nKeep this key secure and don't share it publicly!")
}
