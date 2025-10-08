package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/lachlancoding/advent-of-go-solutions/solutions"
)

func main() {
	// Get required environment variables
	requiredEnvVars := []string{"INPUT_FILE", "SOLUTION_PATH", "DAY", "PART", "YEAR"}
	envVars := make(map[string]string)

	for _, envVar := range requiredEnvVars {
		value := os.Getenv(envVar)
		if value == "" {
			log.Fatalf("Environment variable %s is not set", envVar)
		}
		envVars[envVar] = value
	}

	// Create solution key for module path
	solutionKey := fmt.Sprintf("%s/%s/%s", envVars["YEAR"], envVars["DAY"], envVars["PART"])
	log.Printf("Looking for solution module: %s", solutionKey)

	// Read from input file
	inputFile := envVars["INPUT_FILE"]
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file %s: %v", inputFile, err)
	}

	// Log what we read (for debugging)
	log.Printf("Read %d bytes from input file: %s", len(content), inputFile)
	log.Printf("Input content: %s", string(content))

	// Create the solution directory if it doesn't exist
	solutionPath := envVars["SOLUTION_PATH"]
	err = os.MkdirAll(solutionPath, 0755)
	if err != nil {
		log.Fatalf("Error creating solution directory %s: %v", solutionPath, err)
	}

	// Call the appropriate solution function
	answer := solutions.Solve(solutionKey, string(content))

	// Write answer to file
	answerFile := filepath.Join(solutionPath, "answer")
	err = os.WriteFile(answerFile, []byte(answer), 0644)
	if err != nil {
		log.Fatalf("Error writing answer file %s: %v", answerFile, err)
	}

	log.Printf("Successfully wrote answer '%s' to: %s", answer, answerFile)
}
