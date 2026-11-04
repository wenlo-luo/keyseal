.PHONY: build test vet race

BINDIR ?= dist

build:
	mkdir -p "$(BINDIR)"
	go build -trimpath -o "$(BINDIR)/broker" ./broker/cmd/broker
	go build -trimpath -o "$(BINDIR)/keyGet" ./cli/cmd/keyget

test:
	go test ./...

vet:
	go vet ./...

race:
	go test -race ./...
