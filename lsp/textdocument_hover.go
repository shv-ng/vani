package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_hover
type HoverRequest struct {
	Request
	HoverParams hoverParams `json:"params"`
}

type hoverParams struct {
	textDocumentPosition
}

type HoverResponse struct {
	Response
	HoverResult HoverResult `json:"result"`
}
type HoverResult struct {
	// The hover's content
	Contents string `json:"contents"`
}
