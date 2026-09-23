package ai

import (
	"context"
)

type ChatMessage struct {
	Role    string `json:"role"` // system, user, assistant
	Content string `json:"content"`
}

type AIRequest struct {
	SystemPrompt string
	Personality  string
	Mode         string
	UserNickname string
	UserMood     string
	Memories     []string
	Messages     []ChatMessage
}

type AIProvider interface {
	GenerateStream(ctx context.Context, req AIRequest) (<-chan string, <-chan error)
}
