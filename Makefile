default: vet test

test:
	go test ./...

bench:
	go test ./... -bench=. -run=NONE

vet:
	go vet ./...
