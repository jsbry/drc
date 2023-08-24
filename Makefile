BINARY=bin/drc.exe

PHONY: release
release:
	go build -tags release -o ${BINARY}
