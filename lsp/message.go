package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#requestMessage
type Request struct {
	RPC string `json:"jsonrpc"`
	// The request id.
	ID int `json:"id"`

	// The method to be invoked.
	Method string `json:"method"`

	// The method's params.
	// params
}

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#responseMessage
type Response struct {
	RPC string `json:"jsonrpc"`
	// The request id.
	ID *int `json:"id,omitempty"`
}

type Notification struct {
	RPC string `json:"jsonrpc"`
	// The method to be invoked.
	Method string `json:"method"`
}
