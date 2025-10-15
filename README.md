# Advent of Go Solutions

A modular Go-based solution system for Advent of Code problems.

## Structure

```
solutions/
├── solutions.go              # Main router
├── solutions/
│   └── {YEAR}/
│       └── day{N}/
│           └── part{N}.go          # package day{N}_{YEAR}
├── test_inputs/
├── test_solutions/
├── run_test.sh
```

## Package Naming Convention

To avoid package name collisions between years, we use the format:
- `package day{N}_{YEAR}`

## Usage

Set environment variables and run:

```bash
# For 2024 Day 1 Part 1
$env:YEAR="2024"
$env:DAY="day1"
$env:PART="part1"
$env:INPUT_FILE="input.txt"
$env:SOLUTION_PATH="./output"

./solve
```

## Adding New Solutions

1. Create a new directory: `solutions/{YEAR}/day{N}/`
2. Create `part1.go` with package name `day{N}_{YEAR}`
3. Implement `SolvePart1(input string) string` and `SolvePart2(input string) string`
4. Add import and case to `solutions/solutions.go`

Example for 2025 Day 3:
```go
// solutions/2025/day3/part1.go
package day3_2025

func SolvePart1(input string) string {
    // Your solution here
    return "answer"
}

func SolvePart2(input string) string {
    // Your solution here  
    return "answer"
}
```

Then add to `solutions/solutions.go`:
```go
import (
    day3_2025 "github.com/lachlancoding/advent-of-go-solutions/solutions/2025/day3"
)

// In the switch statement:
case "2025/day3/part1":
    return day3_2025.SolvePart1(input)
case "2025/day3/part2":
    return day3_2025.SolvePart2(input)
```

Testing
-------

The test_inputs folder has a folder for each year/day/part which contains a number of pairs of files answer{N} and case{N}. You can execute these tests by using the `test` make command and providing a SOLUTION_KEY variable of the form `{year}-{day}-{part}-{case}`

e.g. `make test 2024-1-1-1` # tests 2024 day1 part1 testcase 1

The resulting outputs will be stored in the ./test_solutions directory, and the answer will be compared to the answer provided by the test case.
