EXE:=$(shell go env GOEXE)
LDFLAGS:=-ldflags '-w -s'

build:
	go build -o jtlr$(EXE) cmd/jtlr/main.go

release:
	go build $(LDFLAGS) -o jtlr$(EXE) cmd/jtlr/main.go

parser:
	antlr -Dlanguage=Go -o parser/ JSON.g4

.PHONY: build parser release