build:
	go build -o jtlr cmd/main.go

parser:
	antlr -Dlanguage=Go -o parser JSON.g4

.PHONY: build parser