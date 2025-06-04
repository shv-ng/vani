package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentItem
type textDocumentItem struct {
	// The text document's URI.
	URI string `json:"uri"`

	// The text document's language identifier.
	LanguageId string `json:"languageId"`

	// The version number of this document (it will increase after each
	// change, including undo/redo).
	Version int `json:"version"`

	// The content of the opened text document.
	Text string `json:"text"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentPositionParams
type textDocumentPosition struct {
	// The text document.
	TextDocument textDocumentIdentifier `json:"textDocument"`

	// The position inside the text document.
	Position Position `json:"position"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentIdentifier
type textDocumentIdentifier struct {
	// The text document's URI.
	URI string `json:"uri"`
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#Position
type Position struct {
	// Line position in a document (zero-based).
	Line int `json:"line"`

	// Character offset on a line in a document (zero-based). The meaning of this
	// offset is determined by the negotiated `PositionEncodingKind`.
	//
	// If the character value is greater than the line length it defaults back
	// to the line length.
	Character int `json:"character"`
}

type VersionTextDocumentIdentifier struct {
	textDocumentIdentifier
	Version int `json:"version"`
}
