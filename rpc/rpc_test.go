package rpc_test

import (
	"testing"

	"github.com/shv-ng/vani/rpc"
)

type EncodingExample struct {
	Testing bool
}

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	found, err := rpc.EncodeMessage(EncodingExample{Testing: true})
	if err != nil {
		t.Fatalf("Failed to encode message: %v", err)
	}

	if expected != found {
		t.Errorf("Expected: %v, found: %v", expected, found)
	}
}

func TestDecoding(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	contentLength := len(content)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}
	if contentLength != 15 {
		t.Fatalf("Expected: 15, Found: %v", contentLength)
	}
	if method != "hi" {
		t.Fatalf("Expected: hi, Found: %v", method)
	}
}
