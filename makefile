binaryName := jwt-verify
binaryLocation := bin

help:
	@echo Usage:
	@echo
	@echo "  make compile                  Compile for darwin and linux"
	@echo "  make compile-darwin           Compile for all darwin OS's"
	@echo "  make compile-darwin-amd       Compile for AMD64 darwin"
	@echo "  make compile-darwin-arm       Compile for ARM64 darwin"
	@echo "  make compile-linux            Compile for amd64 & 386 linux"
	@echo "  make compile-linux-64         Compile for amd64 linux"
	@echo "  make compile-linux-386        Compile for 386 linux"
	@echo


compile: compile-darwin compile-linux

compile-darwin: compile-darwin-amd compile-darwin-arm

compile-darwin-amd:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o $(binaryLocation)/$(binaryName)-darwin-amd

compile-darwin-arm:
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s -w" -o $(binaryLocation)/$(binaryName)-darwin-arm

compile-linux: compile-linux-386 compile-linux-64

compile-linux-64:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(binaryLocation)/$(binaryName)-linux-amd64

compile-linux-386:
	GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o $(binaryLocation)/$(binaryName)-linux-386