package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/xiaoler/jtlr-go"
)

const (
	INTRO = "jtlr - JSON Tools by Language Recognition.\n"
)

var (
	f_help        bool
	f_indent      string
	f_stdin       bool
	f_interactive bool
)

func init() {
	flag.BoolVar(&f_help, "h", false, "This help")
	flag.StringVar(&f_indent, "t", "", "Set indent characters")
	flag.BoolVar(&f_stdin, "s", false, "Read json from stdin")
	flag.BoolVar(&f_interactive, "a", false, "Run as interactive shell")

	flag.Usage = usage
}

func main() {
	flag.Parse()
	// -h
	if flag.NArg()|flag.NFlag() == 0 || f_help {
		flag.Usage()
		os.Exit(0)
	}
	// -t
	if f_indent != "" {
		jtlr.IDENT_CHAR = f_indent
	}
	// -s
	if f_stdin {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		jtlr.PrettyPrint(text)
		os.Exit(0)
	}
	// -a
	if f_interactive {
		// graceful shutdown
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-sigs
			os.Exit(0)
		}()

		fmt.Print(jtlr.COLOR_Blue, INTRO, jtlr.COLOR_Reset)
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print(">>> ")
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			if err == io.EOF {
				os.Exit(0)
			}
			if text == "\n" || text == "\r\n" {
				continue
			}
			jtlr.PrettyPrint(text)
		}
	}
	// default
	jtlr.PrettyPrint(flag.Arg(0))
}

func usage() {
	fmt.Fprintf(os.Stderr, INTRO+"\nOptions:\n")
	flag.PrintDefaults()
}
