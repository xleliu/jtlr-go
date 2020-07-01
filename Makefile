EXE = $(shell go env GOEXE)

build:
	go build -o jtlr$(EXE) cmd/jtlr/main.go

parser:
	antlr -Dlanguage=Go -o parser JSON.g4

.PHONY: build parser