package main

import (
	"testing"
	"io"
	"bytes"
	"strings"
)

func TestEncodeUnquoted(t *testing.T) {
	encoded := Encode("GET key")
	expected := "*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n"
	if encoded != expected {
		t.Fatal("Encode should handle unquoted args")
	}
}

func TestEncodeQuoted(t *testing.T) {
	encoded := Encode(`GET "key"`)
	expected := "*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n"
	if encoded != expected {
		t.Fatal("Encode should handle quoted args")
	}
}

func TestEncodeSingleCharUnquoted(t *testing.T) {
	encoded := Encode("GET A")
	expected := "*2\r\n$3\r\nGET\r\n$1\r\nA\r\n"
	if encoded != expected {
		t.Fatal("Encode should handle one char unquoted args")
	}
}


func BenchmarkEncode(b *testing.B) {
	var text string = `
HSET foo 3 1.2
HSET foo 4 3.0
HSET foo 5 4.88
HSET bar 3 3.4
HSET foo 12 1.9
HSET foo 3 3.12
`
	var raw io.Reader = strings.NewReader(text)
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		EncodeStream(raw, &buf)
	}
}
