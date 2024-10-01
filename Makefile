PROG=bin/WASM-IPFS-Serverless
SRCS=./cmd/WASM-IPFS-Serverless

INSTALL_PREFIX=/usr/local/WASM-IPFS-Serverless
CONF_INSTALL_PREFIX=/usr/local/WASM-IPFS-Serverless

# git commit hash
COMMIT_HASH=$(shell git rev-parse --short HEAD || echo "GitNotFound")

# build time
BUILD_DATE=$(shell date '+%Y-%m-%d %H:%M:%S')

# build opts
CFLAGS = -ldflags "-s -w  -X \"main.BuildVersion=${COMMIT_HASH}\" -X \"main.BuildDate=$(BUILD_DATE)\""

# Default target
all: build

# Create the bin directory if it doesn't exist
$(shell mkdir -p bin)

# Build for the current platform
build:
	go build $(CFLAGS) -o $(PROG) $(SRCS)

# Build with race detector
race:
	go build $(CFLAGS) -race -o $(PROG) $(SRCS)

# Build for Linux
build-linux:
	GOOS=linux GOARCH=amd64 go build $(CFLAGS) -o $(PROG)-linux $(SRCS)

# Build for ARM (e.g., Raspberry Pi)
build-arm:
	GOOS=linux GOARCH=arm go build $(CFLAGS) -o $(PROG)-arm $(SRCS)

# Build for Windows
build-windows:
	GOOS=windows GOARCH=amd64 go build $(CFLAGS) -o $(PROG)-windows.exe $(SRCS)

# Clean up build artifacts
clean:
	rm -f $(PROG) $(PROG)-linux $(PROG)-arm $(PROG)-windows.exe

# Install the binary to the specified prefix
install:
	install -d $(INSTALL_PREFIX)
	install $(PROG) $(INSTALL_PREFIX)

# Uninstall the binary
uninstall:
	rm -f $(INSTALL_PREFIX)/$(PROG)

.PHONY: all build race build-linux build-arm build-windows clean install uninstall
