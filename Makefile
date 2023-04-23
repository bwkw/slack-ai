.PHONY: build

build:
	GOOS=linux GOARCH=amd64 go build -o main src/main.go && zip main.zip main
