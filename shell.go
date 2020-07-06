package jtlr

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

type hanlder func(string)
type readerWriter struct{}

func (rw *readerWriter) Read(b []byte) (int, error) {
	return os.Stdin.Read(b)
}

func (rw *readerWriter) Write(b []byte) (int, error) {
	return os.Stdout.Write(b)
}

func BasicShell(fn hanlder) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">>> ")
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if text == "\n" || text == "\r\n" {
			continue
		}
		fn(text)
	}
}

func AdvancedShell(fn hanlder) {
	oldState, _ := terminal.MakeRaw(int(os.Stdin.Fd()))
	restore := func() {
		terminal.Restore(int(os.Stdin.Fd()), oldState)
	}
	defer restore()

	term := terminal.NewTerminal(&readerWriter{}, COLOR_Reset+">>> ")
	for {
		text, err := term.ReadLine()
		if err == io.EOF {
			restore()
			break
		}
		if text == "" || text == "\n" || text == "\r\n" {
			continue
		}
		fn(text)
	}
}
