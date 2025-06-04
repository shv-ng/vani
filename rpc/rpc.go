package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

/*
converts any struct/data type to this format

Content-Length: ...\r\n
\r\n

		{
			"jsonrpc": "2.0",
			"id": 1,
			"method": "textDocument/completion",
			"params": {
				...
			}
		}
	 or error
*/
func EncodeMessage(msg any) (encodedMsg string, err error) {
	content, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content), nil
}

// Each request/response contains method
type BaseMessage struct {
	Method string `json:"method"`
}

/*
converts any string of this format

Content-Length: ...\r\n
\r\n

	{
		"jsonrpc": "2.0",
		"id": 1,
		"method": "textDocument/completion",
		"params": {
			...
		}
	}

to (method, content in []byte, err)
*/
func DecodeMessage(msg []byte) (method string, content []byte, err error) {
	header, content, found := bytes.Cut(msg, []byte("\r\n\r\n"))
	if !found {
		return "", nil, errors.New("DecodeMessage failed: separator not found")
	}
	contentLengthBytes := len([]byte("Content-Length: "))
	contentLength, err := strconv.Atoi(string(header[contentLengthBytes:]))
	if err != nil {
		return "", nil, err
	}
	var baseMessage BaseMessage
	err = json.Unmarshal(content, &baseMessage)
	if err != nil {
		return "", nil, err
	}
	return baseMessage.Method, content[:contentLength], nil
}

// split the buf input into bytes of length given by
// contentLength to get next input, it is kind of saperator of buf inputs
func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, nil, err
	}

	if len(content) < contentLength {
		return 0, nil, nil
	}

	totalLength := len(header) + 4 + contentLength

	return totalLength, data[:totalLength], nil
}
