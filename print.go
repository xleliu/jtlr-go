package jtlr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/xiaoler/jtlr-go/parser"
)

var (
	IDENT_CHAR = strings.Repeat(" ", 4)
	CRLF       = "\r\n"
)

type PrettyPrintListener struct {
	indent     int
	arrayLevel int
	arrayPause bool
	*parser.BaseJSONListener
}

type PrettyPrintErrorListener struct {
	*antlr.DefaultErrorListener
}

func NewPrettyPrintErrorListener() *PrettyPrintErrorListener {
	return new(PrettyPrintErrorListener)
}

func (c *PrettyPrintErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line,
	column int,
	msg string,
	e antlr.RecognitionException,
) {
	fmt.Print("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg + CRLF)
}

func PrettyPrint(input string) {
	// Setup the input
	is := antlr.NewInputStream(input)
	// Create the Lexer
	lexer := parser.NewJSONLexer(is)
	lexer.RemoveErrorListeners()
	// lexer.AddErrorListener(NewPrettyPrintErrorListener())

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	parser := parser.NewJSONParser(stream)
	// Finally parse the expression
	parser.RemoveErrorListeners()
	parser.AddErrorListener(NewPrettyPrintErrorListener())

	antlr.ParseTreeWalkerDefault.Walk(NewPrettyPrintListener(), parser.Json())
}

func NewPrettyPrintListener() *PrettyPrintListener {
	return &PrettyPrintListener{
		indent:     0,
		arrayLevel: 0,
		arrayPause: false,
	}
}

// VisitTerminal is called when a terminal node is visited.
func (s *PrettyPrintListener) VisitTerminal(node antlr.TerminalNode) {
	t := node.GetText()
	// unescape unicode, better way?
	if node.GetSymbol().GetTokenType() == parser.JSONLexerSTRING {
		if unquoted, err := strconv.Unquote(t); err == nil {
			t = "\"" + unquoted + "\""
		}
	}

	switch t {
	case "]":
		fmt.Print(CRLF, strings.Repeat(IDENT_CHAR, (s.indent-1)), t)
	case "}":
		fmt.Print(CRLF, strings.Repeat(IDENT_CHAR, s.indent), t)
	case ":":
		fmt.Print(COLOR_Reset, t, " ")
	case "true", "false":
		fmt.Print(COLOR_White, t, COLOR_Reset)
	case "null":
		fmt.Print(COLOR_Dark_Gray, t, COLOR_Reset)
	default:
		fmt.Print(t)
	}
}

// ExitJson is called when production json is exited.
func (s *PrettyPrintListener) ExitJson(ctx *parser.JsonContext) {
	fmt.Print(CRLF)
}

// EnterPair is called when production pair is entered.
func (s *PrettyPrintListener) EnterPair(ctx *parser.PairContext) {
	s.indent++
	fmt.Print(CRLF, strings.Repeat(IDENT_CHAR, s.indent), COLOR_Blue)
}

// ExitPair is called when production pair is exited.
func (s *PrettyPrintListener) ExitPair(ctx *parser.PairContext) {
	s.indent--
}

// EnterArray is called when production array is entered.
func (s *PrettyPrintListener) EnterArray(ctx *parser.ArrayContext) {
	s.arrayLevel++
	s.arrayPause = false
	s.indent++
}

// ExitArray is called when production array is exited.
func (s *PrettyPrintListener) ExitArray(ctx *parser.ArrayContext) {
	s.arrayLevel--
	s.indent--
}

// EnterObject is called when production object is entered.
func (s *PrettyPrintListener) EnterObject(ctx *parser.ObjectContext) {
	s.arrayPause = true
}

// ExitObject is called when production object is exited.
func (s *PrettyPrintListener) ExitObject(ctx *parser.ObjectContext) {
	s.arrayPause = false
}

// EnterValue is called when production value is entered.
func (s *PrettyPrintListener) EnterValue(ctx *parser.ValueContext) {
	if s.arrayLevel > 0 && !s.arrayPause {
		fmt.Print(CRLF, strings.Repeat(IDENT_CHAR, s.indent))
	}
}
