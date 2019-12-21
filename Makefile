# Makefile
build:
	go get -d
	go build -o pollyctl polly.go
