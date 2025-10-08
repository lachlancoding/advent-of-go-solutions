package day1_2024

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// SolvePart1 solves Advent of Code 2024 Day 1 Part 1
func SolvePart1(input string) string {
	inputLines := strings.Split(input, "\n")
	leftList := []int{}
	rightList := []int{}
	for _, line := range inputLines {
		parts := strings.Fields(line)
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}
	sort.Ints(leftList)
	sort.Ints(rightList)

	total_diff := 0
	for i := 0; i < len(leftList); i++ {
		total_diff += int(math.Abs(float64(leftList[i] - rightList[i])))
	}

	return strconv.Itoa(total_diff)
}

// SolvePart2 solves Advent of Code 2024 Day 1 Part 2
func SolvePart2(input string) string {
	// Placeholder implementation for Part 2
	// You can implement the actual solution here
	return "Solution for 2024 Day 1 Part 2 not yet implemented"
}
