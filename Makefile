.PHONY: all build test fmt vet lint clean

all: fmt vet test

build:
	go build ./...

test:
	go test ./...

fmt:
	gofmt -w .

vet:
	go vet ./...

lint:
	golangci-lint run ./...

clean:
	rm -f beispiel.txt person.json todos.json test-output.log
