build:
	go build -o solve cmd/main.go

solve: build
	./solve
