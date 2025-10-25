package ai

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/genai"
)

type GeminiClient struct {
	client    *genai.Client
	modelName string
	ctx       context.Context
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

	// Set the API key in environment for the client
	os.Setenv("GEMINI_API_KEY", apiKey)

	// Create client - it will automatically use GEMINI_API_KEY from environment
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		aiEnabled = false
		return fmt.Errorf("failed to create Gemini client: %w", err)
	}

	modelName := os.Getenv("AI_MODEL")
	if modelName == "" {
		modelName = "gemini-2.0-flash-exp" // Using latest model
	}

	geminiClient = &GeminiClient{
		client:    client,
		modelName: modelName,
		ctx:       ctx,
	}

	aiEnabled = true
	return nil
}

// IsEnabled returns whether AI is enabled
func IsEnabled() bool {
	return aiEnabled
}

// GetResponse sends a prompt to Gemini and returns the response
func GetResponse(prompt string) (string, error) {
	if !aiEnabled || geminiClient == nil {
		return "", fmt.Errorf("AI is not enabled or initialized")
	}

	// Use the new API format: client.Models.GenerateContent()
	result, err := geminiClient.client.Models.GenerateContent(
		geminiClient.ctx,
		geminiClient.modelName,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	// Extract text from result
	text := result.Text()
	if text == "" {
		return "", fmt.Errorf("no response generated")
	}

	return text, nil
}

// Close closes the Gemini client connection
// Note: The new genai SDK doesn't require explicit closing
func Close() error {
	// No-op: new SDK handles cleanup automatically
	return nil
}
