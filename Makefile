run:
	go run main.go

compile:
	goxc -build-ldflags "-s"

test:
	go test

build: test compile
