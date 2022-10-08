single:
	go test -v -cover ./$(package)

all:
	go test -v ./... -race

benchmark:
	go test -bench=. ./...