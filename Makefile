dependencies:
	docker-compose -f http-server/docker-compose.yml up -d

test: dependencies
	go test -v -cover ./$(package)

tests: dependencies
	go test -v ./... -race

benchmarks:
	go test -bench=. ./...