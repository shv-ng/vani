package handler

import (
	"encoding/json"
	"fmt"

	"github.com/ShivangSrivastava/vani/analysis"
	"github.com/ShivangSrivastava/vani/logger"
	"github.com/ShivangSrivastava/vani/lsp"
)

func handleTextDocumentDidChange(state analysis.State, contents []byte) {
	var request lsp.TextDocumentDidChangeNotification
	if err := json.Unmarshal(contents, &request); err != nil {
		logger.Error(fmt.Sprintf("textDocument/didChange: %s", err))
		return
	}

	logger.Info(fmt.Sprintf("Changed: %v", request.Params.TextDocument.URI))
	for _, change := range request.Params.ContentChanges {
		state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
	}
}
