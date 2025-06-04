package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_didOpen
type TextDocumentDidOpenNotification struct {
	Notification Notification
	Params       didOpenTextDocumentParams `json:"params"`
}

type didOpenTextDocumentParams struct {
	TextDocument textDocumentItem `json:"textDocument"`
}
