run:
	go run main.go

compile:
	goxc

test:
	go test

build: test compile
