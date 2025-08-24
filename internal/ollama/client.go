package ollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// OllamaResponse represents the structure of Ollama's API response
type OllamaResponse struct {
	Response string `json:"response"`
}

// FormattedResponse represents the structured output we requested
type FormattedResponse struct {
	Message string `json:"message"`
}

func GetMessage(diff string) (string,error){

	prompt := `
	You are a helpful assistant that can help with generating a git commit message.
	You are given a diff of a codebase and you need to generate a commit message for the changes.
	diffs are in the format of a git diff command.
	The message has to be short and concise.
	The message has to be in the present tense.
	The output should be a single line with the generated commit message.
	Do not include any other text in the output.


	Here is the diff of the changes:

	<diff>

	` + diff + `

	</diff>
	`

	requestBody := map[string]interface{}{
		"model": "llama3.1:8b",
		"prompt": prompt,
		"stream": false,
		"format": map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"message": map[string]interface{}{
					"type": "string",
				},
			},
			"required": []string{"message"},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	response,err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewReader(jsonData))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var ollamaResp OllamaResponse
	if err := json.Unmarshal(body, &ollamaResp); err != nil {
		return "", fmt.Errorf("failed to parse Ollama response: %w", err)
	}

	// Parse the formatted JSON response from Ollama
	var formattedResp FormattedResponse
	if err := json.Unmarshal([]byte(ollamaResp.Response), &formattedResp); err != nil {
		return "", fmt.Errorf("failed to parse formatted response: %w", err)
	}

	return formattedResp.Message, nil
}