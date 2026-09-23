package ai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ClaudeClientProvider struct {
	ApiKey  string
	BaseURL string
	Model   string
	client  *http.Client
}

func NewClaudeClientProvider(apiKey, baseURL, model string) *ClaudeClientProvider {
	if baseURL == "" {
		baseURL = "https://api.anthropic.com"
	}
	if model == "" {
		model = "claude-3-5-haiku-20241022"
	}
	return &ClaudeClientProvider{
		ApiKey:  apiKey,
		BaseURL: baseURL,
		Model:   model,
		client:  &http.Client{},
	}
}

type ClaudeMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ClaudeRequest struct {
	Model     string          `json:"model"`
	MaxTokens int             `json:"max_tokens"`
	System    string          `json:"system,omitempty"`
	Messages  []ClaudeMessage `json:"messages"`
	Stream    bool            `json:"stream"`
}

type ClaudeStreamEvent struct {
	Type  string `json:"type"`
	Delta struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"delta"`
}

func (p *ClaudeClientProvider) GenerateStream(ctx context.Context, req AIRequest) (<-chan string, <-chan error) {
	out := make(chan string)
	errCh := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errCh)

		sysPrompt := buildSystemPrompt(req.Personality, req.Mode, req.UserNickname, req.UserMood, req.Memories)

		var msgs []ClaudeMessage
		for _, m := range req.Messages {
			role := m.Role
			if role == "system" {
				continue
			}
			if role != "user" && role != "assistant" {
				role = "user"
			}
			msgs = append(msgs, ClaudeMessage{
				Role:    role,
				Content: m.Content,
			})
		}

		if len(msgs) == 0 {
			msgs = append(msgs, ClaudeMessage{Role: "user", Content: "Halo"})
		}

		bodyObj := ClaudeRequest{
			Model:     p.Model,
			MaxTokens: 1024,
			System:    sysPrompt,
			Messages:  msgs,
			Stream:    true,
		}

		bodyBytes, err := json.Marshal(bodyObj)
		if err != nil {
			errCh <- err
			return
		}

		url := fmt.Sprintf("%s/v1/messages", strings.TrimRight(p.BaseURL, "/"))
		httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(bodyBytes))
		if err != nil {
			errCh <- err
			return
		}

		httpReq.Header.Set("Content-Type", "application/json")
		httpReq.Header.Set("x-api-key", p.ApiKey)
		httpReq.Header.Set("anthropic-version", "2023-06-01")

		resp, err := p.client.Do(httpReq)
		if err != nil {
			errCh <- err
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			b, _ := io.ReadAll(resp.Body)
			errCh <- fmt.Errorf("claude api returned status %d: %s", resp.StatusCode, string(b))
			return
		}

		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					errCh <- err
				}
				return
			}

			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "data: ") {
				payload := strings.TrimPrefix(line, "data: ")
				if payload == "[DONE]" {
					return
				}

				var event ClaudeStreamEvent
				if err := json.Unmarshal([]byte(payload), &event); err == nil {
					if event.Type == "content_block_delta" && event.Delta.Text != "" {
						select {
						case <-ctx.Done():
							errCh <- ctx.Err()
							return
						case out <- event.Delta.Text:
						}
					} else if event.Type == "message_stop" {
						return
					}
				}
			}
		}
	}()

	return out, errCh
}
