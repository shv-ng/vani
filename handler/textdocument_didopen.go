package handler

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/shv-ng/vani/analysis"
	"github.com/shv-ng/vani/logger"
	"github.com/shv-ng/vani/lsp"
)

func handleTextDocumentDidOpen(state analysis.State, contents []byte) {
	var request lsp.TextDocumentDidOpenNotification
	if err := json.Unmarshal(contents, &request); err != nil {
		logger.Error(fmt.Sprintf("handleTextDocumentDidOpen failed: %v", err))
		return
	}
	if strings.HasPrefix(request.Params.TextDocument.URI, "file") {
		logger.Info(fmt.Sprintf("Opened: %v", request.Params.TextDocument.URI))
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)
	}
}
