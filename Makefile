WASM_BIN = lib.wasm

all: client server

.PHONY: all client server cli

test:
	GOCACHE=off go test -v -race -tags test $(shell go list ./... | grep -v 'vendor\|cmd')

client:
	GOOS=js GOARCH=wasm go build -ldflags "-s -w" -o ${WASM_BIN} ./cmd/client/main.go

server:
	go build -ldflags "-s -w" -o server ./cmd/server/main.go

cli:
	go build -ldflags "-s -w" -o belgrade ./cmd/cli/main.go

run:
	./server

clean:
	rm -rf ${WASM_BIN} server belgrade
