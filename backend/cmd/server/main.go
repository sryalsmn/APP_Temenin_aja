package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"temenin-backend/internal/ai"
	"temenin-backend/internal/config"
	"temenin-backend/internal/handler"
	"temenin-backend/internal/middleware"
	"temenin-backend/internal/repository"
	"temenin-backend/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 1. Initialize Repository
	repo, err := repository.NewRepository(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 2. Initialize AI Provider (Claude AI / OpenAI / Local Simulator)
	var aiProv ai.AIProvider
	if cfg.AIApiKey != "" {
		if cfg.AIProviderType == "claude" || strings.Contains(strings.ToLower(cfg.AIModel), "claude") {
			fmt.Printf(" [AI] Live Anthropic Claude AI Provider activated (Model: %s)\n", cfg.AIModel)
			aiProv = ai.NewClaudeClientProvider(cfg.AIApiKey, cfg.AIBaseURL, cfg.AIModel)
		} else {
			fmt.Printf(" [AI] Live OpenAI/Gemini Provider activated (Model: %s)\n", cfg.AIModel)
			aiProv = ai.NewLLMClientProvider(cfg.AIApiKey, cfg.AIBaseURL, cfg.AIModel)
		}
	} else {
		fmt.Println(" [AI] Local Indonesian Empathetic AI Simulator activated (zero-cost offline mode)")
		aiProv = ai.NewLocalSimulatorProvider()
	}

	// 3. Initialize Services & Handlers
	services := service.NewServices(cfg, repo, aiProv)
	handlers := handler.NewHandlers(services)

	// 4. Setup Router
	r := gin.Default()

	// CORS Setup
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	r.Use(cors.New(corsConfig))

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "TEMENIN API",
			"version": "1.0.0",
		})
	})

	// API v1 Routes (PRD Section 32)
	v1 := r.Group("/api/v1")
	{
		// Authentication (Public)
		auth := v1.Group("/auth")
		{
			auth.POST("/register", handlers.Auth.Register)
			auth.POST("/login", handlers.Auth.Login)
		}

		// Authenticated Routes
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		protected.Use(middleware.RateLimitMiddleware(60)) // 60 req/min
		{
			// User & Profile
			protected.GET("/me", handlers.User.GetMe)
			protected.DELETE("/me", handlers.User.DeleteAccount)

			// Conversations & Chat
			protected.GET("/conversations", handlers.Chat.ListConversations)
			protected.POST("/conversations", handlers.Chat.CreateConversation)
			protected.GET("/conversations/:id", handlers.Chat.GetConversation)
			protected.DELETE("/conversations/:id", handlers.Chat.DeleteConversation)
			protected.POST("/conversations/:id/stream", handlers.Chat.StreamMessage)

			// Mood Tracking
			protected.POST("/moods", handlers.Mood.LogMood)
			protected.GET("/moods", handlers.Mood.GetMoods)
			protected.GET("/moods/summary", handlers.Mood.GetSummary)

			// AI Memory
			protected.GET("/memories", handlers.Memory.GetMemories)
			protected.POST("/memories", handlers.Memory.AddMemory)
			protected.DELETE("/memories/:id", handlers.Memory.DeleteMemory)
			protected.DELETE("/memories", handlers.Memory.ClearAll)

			// Daily Reflection
			protected.GET("/reflections/question", handlers.Reflection.GetQuestion)
			protected.POST("/reflections", handlers.Reflection.SubmitAnswer)
			protected.GET("/reflections", handlers.Reflection.GetHistory)
		}
	}

	addr := fmt.Sprintf(":%s", cfg.Port)
	fmt.Printf("\n TEMENIN Backend running at http://localhost%s\n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
