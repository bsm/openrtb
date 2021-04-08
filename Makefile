default: vet test

.PHONY: test
test:
	go test -a ./...

.PHONY: bench
bench:
	go test ./... -bench=. -run=NONE

.PHONY: vet
vet:
	go vet ./...
