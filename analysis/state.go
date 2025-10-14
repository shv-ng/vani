package analysis

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/shv-ng/vani/data"
	"github.com/shv-ng/vani/logger"
	"github.com/shv-ng/vani/lsp"
)

// State store data of each files, in form of map.
// file path (URI) --> file content
type State struct {
	Documents map[string]string
}

func NewState() State {
	return State{
		Documents: map[string]string{},
	}
}

// textDocument/openDocument
func (s *State) OpenDocument(uri, text string) {
	s.Documents[uri] = text
}

// textDocument/didChange
// textDocument/didSave
func (s *State) UpdateDocument(uri, text string) {
	s.Documents[uri] = text
}

// textDocument/hover
func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	word := getWord(s.Documents[uri], position)
	def, err := GetWordMeaningCache(word)
	if err != nil {
		logger.Error(fmt.Sprintf("error on definition: %v", err))
	}
	return lsp.HoverResponse{
		HoverResult: lsp.HoverResult{
			Contents: hoverText(word, def),
		},
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
	}
}

func getWord(contents string, position lsp.Position) string {
	lines := strings.Split(contents, "\n")
	if position.Line < 0 || position.Line > len(lines) {
		return ""
	}

	line := lines[position.Line]
	if position.Character < 0 || position.Character > len(line) {
		return ""
	}
	// Expand left
	start := position.Character
	for start > 0 && isAlpha(line[start-1]) {
		start--
	}

	// Expand right
	end := position.Character
	for end < len(line) && isAlpha(line[end]) {
		end++
	}
	return line[start:end]
}

// check is [a-zA-Z]
func isAlpha(r byte) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

// markdown hover Response for better ux
func hoverText(word, definition string) string {
	return fmt.Sprintf(
		"```\n%s\n```\n---\n%s\n---\n[%s](%s)",
		word,
		definition,
		data.GetData().ServerName,
		data.GetData().RepositoryLink,
	)
}

func (s *State) TextDocumentCompletion(id int, uri string) lsp.CompletionResponse {
	if items == nil {
		LoadCompletionItems("/home/shivang/personal/vani/data/words-large.txt")
		LoadCompletionItems("/home/shivang/personal/vani/data/words-medium.txt")
	}
	return lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			ID:  &id,
		},
		Result: items,
	}
}

var items []lsp.CompletionItem

func LoadCompletionItems(filename string) []lsp.CompletionItem {
	file, err := os.Open(filename)
	if err != nil {
		logger.Error(fmt.Sprintf("LoadCompletionItems failed: %v", err))
		return nil
	}
	data, err := io.ReadAll(file)
	if err != nil {
		logger.Error(fmt.Sprintf("LoadCompletionItems on read failed: %v", err))
		return nil
	}

	content := string(data)
	lines := strings.SplitSeq(content, "\n")
	for word := range lines {
		items = append(items, lsp.CompletionItem{
			Label: word,
		})
	}
	return items
}
