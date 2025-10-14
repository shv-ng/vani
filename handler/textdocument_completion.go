package handler

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shv-ng/vani/analysis"
	"github.com/shv-ng/vani/logger"
	"github.com/shv-ng/vani/lsp"
)

func handleTextDocumentCompletion(state analysis.State, contents []byte) {
	var request lsp.CompletionRequest
	if err := json.Unmarshal(contents, &request); err != nil {
		logger.Error(fmt.Sprintf("handleTextDocumentCompletion failed: %v", err))
		return
	}

	response := state.TextDocumentCompletion(request.ID, request.Params.TextDocument.URI)

	writer := os.Stdout
	if err := WriteResponse(writer, response); err != nil {
		logger.Error(fmt.Sprintf("EncodeMessage failed: %v", err))
	}
}
