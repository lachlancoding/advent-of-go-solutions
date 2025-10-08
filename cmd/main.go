package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Get INPUT_FILE environment variable
	inputFile := os.Getenv("INPUT_FILE")
	if inputFile == "" {
		log.Fatal("INPUT_FILE environment variable is not set")
	}

	// Get SOLUTION_PATH environment variable
	solutionPath := os.Getenv("SOLUTION_PATH")
	if solutionPath == "" {
		log.Fatal("SOLUTION_PATH environment variable is not set")
	}

	// Read from input file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file %s: %v", inputFile, err)
	}

	// Log what we read (for debugging)
	log.Printf("Read %d bytes from input file: %s", len(content), inputFile)
	log.Printf("Input content: %s", string(content))

	// Write answer to file
	answerFile := filepath.Join(solutionPath, "answer")
	err = os.WriteFile(answerFile, []byte("the answer is: correct :)"), 0644)
	if err != nil {
		log.Fatalf("Error writing answer file %s: %v", answerFile, err)
	}

	log.Printf("Successfully wrote answer to: %s", answerFile)
}
