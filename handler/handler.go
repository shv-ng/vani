package handler

import (
	"fmt"
	"io"

	"github.com/shv-ng/vani/analysis"
	"github.com/shv-ng/vani/logger"
	"github.com/shv-ng/vani/rpc"
)

var handlers = map[string]func(analysis.State, []byte){
	"initialize":              handleInitialize,
	"textDocument/didOpen":    handleTextDocumentDidOpen,
	"textDocument/hover":      handleTextDocumentHover,
	"textDocument/completion": handleTextDocumentCompletion,
	"textDocument/didChange":  handleTextDocumentDidChange,
}

func Handle(state analysis.State, method string, contents []byte) {
	// say in log to what we got
	if handleFunc, ok := handlers[method]; ok {
		logger.Info(fmt.Sprintf("Method recieved: %v", method))
		handleFunc(state, contents)
	} else {
		logger.Warn(fmt.Sprintf("Unknown method recieved: %v", method))
	}
}

func WriteResponse(w io.Writer, msg any) error {
	reply, err := rpc.EncodeMessage(msg)
	if err != nil {
		return err
	}
	w.Write([]byte(reply))
	return nil
}
