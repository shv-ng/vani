package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_completion
type CompletionRequest struct {
	Request
	Params completionParams `json:"params"`
}

type completionParams struct {
	textDocumentPosition
}

type CompletionResponse struct {
	Response
	Result []CompletionItem `json:"result"`
}

type CompletionItem struct {
	Label string `json:"label"`
	// TODO: add this
	// Documentation string `json:"documentation"`
}
