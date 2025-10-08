.PHONY: build dev-run solve

SOLUTION_KEY=2024-1-1-1

test:
	./run_test.sh $(SOLUTION_KEY)

build:
	go build -o ./bin/solve cmd/main.go

solve: build
	./bin/solve
