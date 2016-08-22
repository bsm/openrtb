default: vet test

test:
	go test ./...

vet:
	go vet ./...
