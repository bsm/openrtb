default: test

deps:
	go get -t ./...

test: deps
	go test ./...
