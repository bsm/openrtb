default: vet test

test:
	go test -a ./...

bench:
	go test ./... -bench=. -run=NONE

vet:
	go vet ./...
