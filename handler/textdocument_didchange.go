package handler

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shv-ng/vani/analysis"
	"github.com/shv-ng/vani/logger"
	"github.com/shv-ng/vani/lsp"
)

func handleTextDocumentDidChange(state analysis.State, contents []byte) {
	var request lsp.TextDocumentDidChangeNotification
	if err := json.Unmarshal(contents, &request); err != nil {
		logger.Error(fmt.Sprintf("textDocument/didChange: %s", err))
		return
	}

	logger.Info(fmt.Sprintf("Changed: %v", request.Params.TextDocument.URI))
	if !strings.HasPrefix(request.Params.TextDocument.URI, "file") {
		return
	}
	for _, change := range request.Params.ContentChanges {
		state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
	}
}
