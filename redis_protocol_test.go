package main

import (
	"fmt"
	"testing"
)

func TestEncodeUnquoted(t *testing.T) {
	encoded := Encode("GET key")
	expected := "*2\r\n$3\r\nGET\r\n$3\r\nkey\r\n"
	fmt.Println(encoded)
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
