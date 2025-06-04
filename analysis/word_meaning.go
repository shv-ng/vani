package analysis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// structs that are use to fetch api
type Definition struct {
	Definition string   `json:"definition"`
	Example    string   `json:"example"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
}

type DictionaryEntry struct {
	Word     string    `json:"word"`
	Meanings []Meaning `json:"meanings"`
}

// Temp cache to store word meaning
// TODO: change it to permanent db
var meaningsCache = make(map[string]string)

// get and add word meaning form/to the cache
func GetWordMeaningCache(word string) (string, error) {
	lowerWord := strings.ToLower(word)
	if meaning, ok := meaningsCache[lowerWord]; ok {
		return meaning, nil
	}
	go func() {
		meaning, err := GetWordMeaning(lowerWord)
		if err != nil {
			return
		}
		if len(meaning) > 1 {
			meaningsCache[lowerWord] = meaning
		}
	}()
	return "No definition found", nil
}

// Call api to get meaning and more info of the word
func GetWordMeaning(word string) (string, error) {
	url := fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	var entries []DictionaryEntry
	if err := json.Unmarshal(body, &entries); err != nil {
		return "", fmt.Errorf("failed to parse JSON: %v", err)
	}

	var sb strings.Builder
	for _, entry := range entries {
		sb.WriteString(fmt.Sprintf("## Word: **%s**\n\n", entry.Word))
		for _, meaning := range entry.Meanings {
			sb.WriteString(fmt.Sprintf("### Part of Speech: _%s_\n", meaning.PartOfSpeech))
			for i, def := range meaning.Definitions {
				sb.WriteString(fmt.Sprintf("- **Definition %d:** %s\n", i+1, def.Definition))
				if def.Example != "" {
					sb.WriteString(fmt.Sprintf("  - _Example:_ \"%s\"\n", def.Example))
				}
			}
			sb.WriteString("\n")
		}
	}

	return sb.String(), nil
}
