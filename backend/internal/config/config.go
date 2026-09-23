package config

import (
	"bufio"
	"os"
	"strings"
)

type Config struct {
	AppEnv         string
	Port           string
	DatabaseURL    string
	JWTSecret      string
	AIProviderType string
	AIApiKey       string
	AIBaseURL      string
	AIModel        string
	RedisURL       string
	AllowedOrigins string
}

func loadDotEnv() {
	f, err := os.Open(".env")
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			k := strings.TrimSpace(parts[0])
			v := strings.TrimSpace(parts[1])
			os.Setenv(k, v)
		}
	}
}

func LoadConfig() *Config {
	loadDotEnv()

	provider := strings.ToLower(getEnv("AI_PROVIDER", "groq"))
	defaultBaseURL := "https://api.groq.com/openai/v1"
	defaultModel := "openai/gpt-oss-120b"

	if provider == "claude" {
		defaultBaseURL = "https://api.anthropic.com"
		defaultModel = "claude-3-5-haiku-20241022"
	} else if provider == "gemini" {
		defaultBaseURL = "https://generativelanguage.googleapis.com/v1beta/openai"
		defaultModel = "gemini-1.5-flash"
	}

	baseURL := getEnv("AI_BASE_URL", defaultBaseURL)
	// Auto-fix if user put web console url instead of api url
	if strings.Contains(baseURL, "console.groq.com") || baseURL == "" {
		baseURL = "https://api.groq.com/openai/v1"
	}

	model := getEnv("AI_MODEL", defaultModel)
	if provider == "groq" && strings.Contains(model, "claude") {
		model = "openai/gpt-oss-120b"
	}

	return &Config{
		AppEnv:         getEnv("APP_ENV", "development"),
		Port:           getEnv("PORT", "8080"),
		DatabaseURL:    getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/temenin?sslmode=disable"),
		JWTSecret:      getEnv("JWT_SECRET", "super-secret-temenin-jwt-key-change-in-prod"),
		AIProviderType: provider,
		AIApiKey:       getEnv("AI_API_KEY", ""),
		AIBaseURL:      baseURL,
		AIModel:        model,
		RedisURL:       getEnv("REDIS_URL", ""),
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "*"),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
