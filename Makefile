BINARY_NAME=ice

init: install clean mod tidy vet

install:
	go install github.com/AlexBeauchemin/gobadge@latest

clean:
	go clean
	rm -f ${BINARY_NAME}
	rm *.out

mod:
	go mod download

tidy:
	go mod tidy

vet:
	go vet ./...

build:
	go build -o ${BINARY_NAME} .

lint:
	golangci-lint run

fix:
	golangci-lint run --fix

test:
	go test ./... -covermode=count -coverprofile=coverage.out fmt

coverage: test
	go tool cover -html=coverage.out

badge: test
	go tool cover -func=coverage.out -o=coverage.out
	gobadge -filename=coverage.out

run:
	go run .
