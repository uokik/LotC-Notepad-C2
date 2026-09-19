AGENT_BINARY=agent.exe
SERVER_BINARY=server

.PHONY: all build build-agent build-server test clean

all: clean test build
build: build-agent build-server

build-agent:
	@echo "Compiling agent for windows without CLI (amd64)..."
	env GOOS=windows GOARCH=amd64 go build -ldflags="-H=windowsgui -s -w" -o bin/$(AGENT_BINARY) cmd/agent/main.go
	@echo "Agent successfully compiled (bin/$(AGENT_BINARY))"

build-server:
	@echo "Compiling server for linux..."
	go build -o bin/$(SERVER_BINARY) cmd/server/main.go
	@echo "Server successfully compiled (bin/$(AGENT_BINARY))"

agent-debug:
	@echo "Compiling agent for windows with CLI (amd64)..."
	env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/$(AGENT_BINARY) cmd/agent/main.go
	@echo "Agent successfully compiled (bin/$(AGENT_BINARY))"

test:
	@echo "Testing.."
	go test -v ./internal/...

clean:
	@echo "Cleaning bin/..."
	rm -rf bin/