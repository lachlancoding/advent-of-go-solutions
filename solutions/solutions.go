package solutions

import (
	"log"

	day1_2024 "github.com/lachlancoding/advent-of-go-solutions/solutions/2024/day1"
)

// Solve routes to the appropriate solution based on the solution key (year/day/part)
func Solve(solutionKey, input string) string {
	log.Printf("Solving for key: %s", solutionKey)

	switch solutionKey {
	case "2024/day1/part1":
		return day1_2024.SolvePart1(input)
	case "2024/day1/part2":
		return day1_2024.SolvePart2(input)

	// Add more cases as you create more solutions
	default:
		log.Fatalf("No solution found for key: %s", solutionKey)
		return ""
	}
}
