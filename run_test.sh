#!/bin/bash
SOLUTION_KEY=$1

year=$(echo $SOLUTION_KEY | cut -d "-" -f1)
day_num=$(echo $SOLUTION_KEY | cut -d "-" -f2)
part_num=$(echo $SOLUTION_KEY | cut -d "-" -f3)
case=$(echo $SOLUTION_KEY | cut -d "-" -f4)

# Format day and part with prefixes as expected by the Go program
day="day$day_num"
part="part$part_num"

export INPUT_FILE=./test_inputs/$year/$day/$part/case$case
export SOLUTION_PATH=./test_solutions/
export YEAR=$year
export DAY=$day
export PART=$part

echo "INPUT_FILE: $INPUT_FILE"
echo "SOLUTION_PATH: $SOLUTION_PATH"
echo "YEAR: $YEAR"
echo "DAY: $DAY"
echo "PART: $PART"

# Build the solver first
go build -o solve cmd/main.go

# Run the solver
./solve

# Compare result with expected answer
expected_answer="./test_inputs/$year/$day/$part/answer$case"
actual_answer="./test_solutions/answer"

if [ -f "$expected_answer" ]; then
    if cmp -s "$expected_answer" "$actual_answer"; then
        echo "✅ Test PASSED: Output matches expected answer"
    else
        echo "❌ Test FAILED: Output differs from expected answer"
        echo "Expected:"
        cat "$expected_answer"
        echo "Actual:"
        cat "$actual_answer"
        echo "Diff:"
        diff "$expected_answer" "$actual_answer"
    fi
else
    echo "⚠️  No expected answer file found at $expected_answer"
    echo "Actual output:"
    cat "$actual_answer"
fi