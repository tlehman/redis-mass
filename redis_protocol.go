package main

import (
	"bytes"
	"fmt"
	"strings"
)

func Encode(text string) string {
	var splitText []string = strings.Split(text, "\n")
	var commands []string = splitText
	var protocol bytes.Buffer

	var args []string
	var length int

	for _, command := range commands {
		args = parseRedisCommand(strings.TrimSpace(command))
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

func parseRedisCommand(command string) []string {
	return []string{"foo", "bar"}
}
