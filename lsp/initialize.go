package lsp

import "github.com/shv-ng/vani/data"

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialize
type InitializeRequest struct {
	Request
	Params initializeParams `json:"params"`
}

type initializeParams struct {
	ClientInfo *clientInfo `json:"clientInfo"`
}

type clientInfo struct {
	// The name of the client as defined by the client.
	Name string `json:"name"`
	// The client's version as defined by the client.
	Version string `json:"version"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initializeResult
type InitializeResponse struct {
	Response
	Result initializeResult `json:"result"`
}

type initializeResult struct {
	// Information about the server.
	ServerInfo serverInfo `json:"serverInfo"`
	// The capabilities the language server provides.
	Capabilities serverCapabilities `json:"capabilities"`
}

type serverInfo struct {
	// The name of the server as defined by the server.
	Name string `json:"name"`

	// The server's version as defined by the server.
	Version string `json:"version"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#serverCapabilities
type serverCapabilities struct {
	// Defines how text documents are synced. Is either a detailed structure
	// defining each notification or for backwards compatibility the
	// TextDocumentSyncKind number. If omitted it defaults to
	// `TextDocumentSyncKind.None`.
	// 0 -> None, 1 -> Full, 2 -> incremental
	TextDocumentSync int `json:"textDocumentSync"`

	// The server provides hover support.
	HoverProvider bool `json:"hoverProvider"`

	// Add docs
	CompletionProvider map[string]any `json:"completionProvider"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		Response: Response{
			ID:  &id,
			RPC: "2.0",
		},
		Result: initializeResult{
			ServerInfo: serverInfo{
				Name:    data.GetData().ServerName,
				Version: data.GetData().Version,
			},
			Capabilities: serverCapabilities{
				TextDocumentSync:   1,
				HoverProvider:      true,
				CompletionProvider: map[string]any{},
			},
		},
	}
}
