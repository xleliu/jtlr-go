package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/xiaoler/jtlr-go"
)

const (
	INTRO   = "jtlr (v%s) - JSON Tools by Language Recognition."
	VERSION = "0.02"
)

var (
	f_help        bool
	f_indent      string
	f_stdin       bool
	f_interactive bool
	f_file        string
)

func init() {
	flag.BoolVar(&f_help, "h", false, "This help")
	flag.StringVar(&f_indent, "t", "", "Set indent characters")
	flag.BoolVar(&f_stdin, "s", false, "Read json from stdin")
	flag.BoolVar(&f_interactive, "a", false, "Run as interactive shell")
	flag.StringVar(&f_file, "f", "", "Read json from file")

	flag.Usage = usage
}

func main() {
	flag.Parse()
	// -h
	if flag.NArg()|flag.NFlag() == 0 || f_help {
		flag.Usage()
		os.Exit(0)
	}
	// Graceful shutdown for -s/-a
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		os.Exit(0)
	}()

	// -t
	if f_indent != "" {
		if f_indent == "\\t" {
			f_indent = "\t"
		}
		jtlr.IDENT_CHAR = f_indent
	}
	// -f
	if f_file != "" {
		file, err := os.Open(f_file)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			jtlr.PrettyPrint(scanner.Text())
		}
		os.Exit(0)
	}
	// -s
	if f_stdin {
		scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
		for scanner.Scan() {
			jtlr.PrettyPrint(scanner.Text())
		}
		os.Exit(0)
	}
	// -a
	if f_interactive {
		fmt.Printf(jtlr.COLOR_Blue+INTRO+jtlr.COLOR_Reset+jtlr.CRLF, VERSION)
		if runtime.GOOS == "windows" {
			jtlr.BasicShell(jtlr.PrettyPrint)
		} else {
			jtlr.AdvancedShell(jtlr.PrettyPrint)
		}
		os.Exit(0)
	}
	// default
	jtlr.PrettyPrint(flag.Arg(0))
}

func usage() {
	fmt.Printf(INTRO+"\nOptions:\n", VERSION)
	flag.PrintDefaults()
}
