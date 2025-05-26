package service

import (
	"URL_OPENER/logger"
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func GenerateLocalOllamaResponse(model string, prompt string) (string, error) {
	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	response, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
	if err != nil {
		return "", err
	}

	if len(response) == 0 {
		return "", fmt.Errorf("no response - ollama local llm")
	}

	return response, nil
}

func TestLocalOllama(model string) {
	prompt := "You know everything about computer science. Create a short informative description of what Go is"

	response, err := GenerateLocalOllamaResponse(model, prompt)
	if err != nil {
		logger.ErrorF("Error generating response: %v", err)
	}

	logger.InfoF("Ollama local llm response: %s", response)
}
