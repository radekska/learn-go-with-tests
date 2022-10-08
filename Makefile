test:
	go test -v -cover ./$(package)

tests:
	go test -v ./... -race

benchmarks:
	go test -bench=. ./...