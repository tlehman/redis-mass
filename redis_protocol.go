package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type parserFunc func(*parserState, rune, bool)

type parserState struct {
	parser parserFunc
	args   []string
	buffer bytes.Buffer
}

func waitingArgStart(state *parserState, c rune, eol bool) {
	if c != ' ' {
		if c == '"' {
			state.parser = waitingQuotedArgumentEnd
		} else {
			state.parser = waitingUnquotedArgumentEnd
			state.parser(state, c, eol)
		}
	}
}

func waitingArgEnd(endToken rune, state *parserState, c rune, eol bool) {
	if c != endToken {
		state.buffer.WriteRune(c)
	}
	if eol || c == endToken {
		state.args = append(state.args, state.buffer.String())
		state.buffer.Reset()
		state.parser = waitingArgStart
	}
}

func waitingQuotedArgumentEnd(state *parserState, c rune, eol bool) {
	waitingArgEnd('"', state, c, eol)
}

func waitingUnquotedArgumentEnd(state *parserState, c rune, eol bool) {
	waitingArgEnd(' ', state, c, eol)
}

// parse a redis command
func parse(command string) []string {
	var length int = len(command)
	var eol bool
	var state parserState
	state.parser = waitingArgStart
	state.args = []string{}

	for i, c := range command {
		eol = (i == length-1)
		state.parser(&state, c, eol)
	}
	return state.args
}

func EncodeStream(raw io.Reader, enc io.Writer) {
	var args []string
	var length int

	scanner := bufio.NewScanner(raw)

	for scanner.Scan() {
		command := strings.TrimSpace(scanner.Text())
		args = parse(command)
		length = len(args)
		if length > 0 {
			io.WriteString(enc, fmt.Sprintf("*%d\r\n", length))
			for _, arg := range args {
				io.WriteString(enc, fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg))
			}
		}
	}
}

func Encode(text string) string {
	var raw io.Reader = strings.NewReader(text)
	var buf bytes.Buffer
	enc := bufio.NewWriter(&buf)

	EncodeStream(raw, enc)

	err := enc.Flush()
	if err != nil {
		fmt.Printf("error flushing encoded writer: %v", err)
	}

	return buf.String()
}
