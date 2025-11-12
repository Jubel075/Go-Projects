package ai

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"google.golang.org/genai"
)

type GeminiClient struct {
	client    *genai.Client
	modelName string
	ctx       context.Context
	history   []*genai.Content // New: conversation history
}

var (
	geminiClient *GeminiClient
	aiEnabled    bool
)

// Initialize sets up the Gemini AI client
func Initialize() error {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		aiEnabled = false
		return fmt.Errorf("GEMINI_API_KEY not found in environment")
	}

	enabled := os.Getenv("AI_ENABLED")
	if enabled == "false" {
		aiEnabled = false
		return nil
	}
	ctx := context.Background()

	// Use ClientConfig for explicit API key and backend (avoids redundant env set)
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		aiEnabled = false
		return fmt.Errorf("failed to create Gemini client: %w", err)
	}

	modelName := os.Getenv("AI_MODEL")
	if modelName == "" {
		modelName = "gemini-2.5-flash" // Updated to match guide
	}

	geminiClient = &GeminiClient{
		client:    client,
		modelName: modelName,
		ctx:       ctx,
		history:   []*genai.Content{}, // Initialize empty history
	}

	aiEnabled = true
	return nil
}

// IsEnabled returns whether AI is enabled
func IsEnabled() bool {
	return aiEnabled
}

// StreamResponse sends a prompt to Gemini and streams the response chunk by chunk.
// onChunk is called for each real chunk from the API.
// onComplete receives the full final string when streaming finishes.
func StreamResponse(prompt string, onChunk func(string), onComplete func(string)) error {
	if !aiEnabled || geminiClient == nil {
		return fmt.Errorf("AI is not enabled or initialized")
	}

	systemPrompt := "You are a helpful CLI assistant. Provide clear, concise responses. " +
		"Use markdown sparingly - prefer plain text with occasional formatting. " +
		"Keep responses focused and terminal-friendly. Be conversational but brief."

	// Build contents with history + new user prompt
	contents := append(append([]*genai.Content{}, geminiClient.history...),
		&genai.Content{
			Role:  "user",
			Parts: []*genai.Part{{Text: prompt}},
		})

	// If first message, prepend system prompt to the user prompt
	if len(geminiClient.history) == 0 {
		contents[0].Parts[0].Text = systemPrompt + "\n\nUser: " + contents[0].Parts[0].Text
	}

	// Use real streaming
	iter := geminiClient.client.Models.GenerateContentStream(
		geminiClient.ctx,
		geminiClient.modelName,
		contents,
		nil,
	)

	var builder strings.Builder

	for result, err := range iter {
		if err != nil {
			return fmt.Errorf("streaming error: %w", err)
		}

		// Extract text from the partial response (manual extraction)
		if len(result.Candidates) == 0 || result.Candidates[0].Content == nil || len(result.Candidates[0].Content.Parts) == 0 {
			continue
		}
		chunk := result.Candidates[0].Content.Parts[0].Text
		if chunk == "" {
			continue
		}

		if onChunk != nil {
			onChunk(chunk)
		}
		builder.WriteString(chunk)

		// Optional small delay for visual typing effect (remove for max speed)
		time.Sleep(20 * time.Millisecond)
	}

	fullText := builder.String()
	if onComplete != nil {
		onComplete(fullText)
	}

	// On success, append user and model to history
	geminiClient.history = append(geminiClient.history, contents[len(contents)-1]) // user
	geminiClient.history = append(geminiClient.history,
		&genai.Content{
			Role:  "model",
			Parts: []*genai.Part{{Text: fullText}},
		})

	return nil
}

// GetResponse sends a prompt to Gemini and returns the response (non-streaming fallback)
func GetResponse(prompt string) (string, error) {
	if !aiEnabled || geminiClient == nil {
		return "", fmt.Errorf("AI is not enabled or initialized")
	}

	systemPrompt := "You are a helpful CLI assistant. Provide clear, concise responses. " +
		"Use markdown sparingly - prefer plain text with occasional formatting. " +
		"Keep responses focused and terminal-friendly. Be conversational but brief."

	// Build contents with history + new user prompt
	contents := append(append([]*genai.Content{}, geminiClient.history...),
		&genai.Content{
			Role:  "user",
			Parts: []*genai.Part{{Text: prompt}},
		})

	// If first message, prepend system prompt to the user prompt
	if len(geminiClient.history) == 0 {
		contents[0].Parts[0].Text = systemPrompt + "\n\nUser: " + contents[0].Parts[0].Text
	}

	result, err := geminiClient.client.Models.GenerateContent(
		geminiClient.ctx,
		geminiClient.modelName,
		contents,
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	// Extract text manually
	if len(result.Candidates) == 0 || result.Candidates[0].Content == nil || len(result.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no response generated")
	}
	text := result.Candidates[0].Content.Parts[0].Text
	if text == "" {
		return "", fmt.Errorf("no response generated")
	}

	// On success, append user and model to history
	geminiClient.history = append(geminiClient.history, contents[len(contents)-1]) // user
	geminiClient.history = append(geminiClient.history,
		&genai.Content{
			Role:  "model",
			Parts: []*genai.Part{{Text: text}},
		})

	return text, nil
}

// Close closes the Gemini client connection (no-op for current SDK)
func Close() error {
	return nil
}
