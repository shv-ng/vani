package handler

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/shv-ng/vani/analysis"
	"github.com/shv-ng/vani/logger"
	"github.com/shv-ng/vani/lsp"
)

func handleTextDocumentHover(state analysis.State, contents []byte) {
	var request lsp.HoverRequest
	if err := json.Unmarshal(contents, &request); err != nil {
		logger.Error(fmt.Sprintf("handleTextDocumentHover failed: %v", err))
		return
	}

	response := state.Hover(request.ID, request.HoverParams.TextDocument.URI, request.HoverParams.Position)

	writer := os.Stdout
	if err := WriteResponse(writer, response); err != nil {
		logger.Error(fmt.Sprintf("EncodeMessage failed: %v", err))
	}
}
