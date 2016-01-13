package main

import (
	"bytes"
	"fmt"
	"strings"
)

type parserFunc func(parserState, rune, bool)

type parserState struct {
	parser parserFunc
	args   []string
	buffer bytes.Buffer
}

func waitingArgStart(state parserState, c rune, eol bool) {
	if c != ' ' {
		if c == '"' {
			state.parser = waitingQuotedArgumentEnd
		} else {
			state.parser = waitingUnquotedArgumentEnd
			state.parser(state, c, eol)
		}
	}
}

func waitingArgEnd(endToken rune, state parserState, c rune, eol bool) {
	if c != endToken {
		state.buffer.WriteRune(c)
	}
	if eol || c == endToken {
		state.args = append(state.args, state.buffer.String())
	}
}

func waitingQuotedArgumentEnd(state parserState, c rune, eol bool) {
}

func waitingUnquotedArgumentEnd(state parserState, c rune, eol bool) {
}

// parse a redis command
func parse(command string) []string {
	var args []string
	var length int = len(command)
	var eol bool
	var state parserState
	state.parser = waitingArgStart
 	state.args = []string{}

	for i, c := range command {
		eol = (i == length-1)
		state.parser(state, c, eol)
	}
	return args
}

func Encode(text string) string {
	var splitText []string = strings.Split(text, "\r\n")
	var commands []string = splitText
	var protocol bytes.Buffer

	var args []string
	var length int

	for _, command := range commands {
		args = parse(strings.TrimSpace(command))
		length = len(args)
		if length > 0 {
			protocol.WriteString(fmt.Sprintf("*%d\r\n", length))
			for _, arg := range args {
				encoded := fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg)
				protocol.WriteString(encoded)
			}
		}
	}

	return protocol.String()
}
