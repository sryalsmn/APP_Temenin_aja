package middleware

import (
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"temenin-backend/pkg/jwt"
	"temenin-backend/pkg/response"
)

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// In local dev/demo environment, allow demo-user so client can stream without friction
			c.Set("user_id", "demo-user")
			c.Set("email", "demo@temenin.id")
			c.Next()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.Unauthorized(c, "Invalid authorization format, must be Bearer token")
			c.Abort()
			return
		}

		claims, err := jwt.ValidateToken(parts[1], secret)
		if err != nil {
			response.Unauthorized(c, "Invalid or expired token")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("email", claims.Email)
		c.Next()
	}
}

type visitor struct {
	lastSeen time.Time
	count    int
}

var (
	visitors = make(map[string]*visitor)
	mu       sync.Mutex
)

func RateLimitMiddleware(limitPerMin int) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		mu.Lock()
		v, exists := visitors[ip]
		if !exists || time.Since(v.lastSeen) > time.Minute {
			visitors[ip] = &visitor{lastSeen: time.Now(), count: 1}
			mu.Unlock()
			c.Next()
			return
		}

		v.count++
		if v.count > limitPerMin {
			mu.Unlock()
			response.Error(c, 429, "Too many requests. Please slow down.")
			c.Abort()
			return
		}
		mu.Unlock()
		c.Next()
	}
}
