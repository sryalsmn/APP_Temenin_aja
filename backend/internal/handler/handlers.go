package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"temenin-backend/internal/service"
	"temenin-backend/pkg/response"
)

type Handlers struct {
	Auth       *AuthHandler
	User       *UserHandler
	Chat       *ChatHandler
	Mood       *MoodHandler
	Memory     *MemoryHandler
	Reflection *ReflectionHandler
}

func NewHandlers(services *service.Services) *Handlers {
	return &Handlers{
		Auth:       &AuthHandler{svc: services.Auth},
		User:       &UserHandler{svc: services.Auth},
		Chat:       &ChatHandler{svc: services.Chat},
		Mood:       &MoodHandler{svc: services.Mood},
		Memory:     &MemoryHandler{svc: services.Memory},
		Reflection: &ReflectionHandler{svc: services.Reflection},
	}
}

// ==================== AUTH HANDLER ====================

type AuthHandler struct {
	svc *service.AuthService
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Nickname string `json:"nickname" binding:"required"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.svc.Register(req.Email, req.Password, req.Nickname)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Registrasi berhasil", result)
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.svc.Login(req.Email, req.Password)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Login berhasil", result)
}

// ==================== USER HANDLER ====================

type UserHandler struct {
	svc *service.AuthService
}

func (h *UserHandler) GetMe(c *gin.Context) {
	userID := c.GetString("user_id")
	user, profile, settings, err := h.svc.GetMe(userID)
	if err != nil {
		response.NotFound(c, "User not found")
		return
	}

	response.Success(c, http.StatusOK, "Data profil didapatkan", gin.H{
		"user":     user,
		"profile":  profile,
		"settings": settings,
	})
}

func (h *UserHandler) DeleteAccount(c *gin.Context) {
	userID := c.GetString("user_id")
	if err := h.svc.DeleteAccount(userID); err != nil {
		response.InternalError(c, "Gagal menghapus akun")
		return
	}

	response.Success(c, http.StatusOK, "Akun dan semua data berhasil dihapus", nil)
}

// ==================== CHAT HANDLER ====================

type ChatHandler struct {
	svc *service.ChatService
}

func (h *ChatHandler) ListConversations(c *gin.Context) {
	userID := c.GetString("user_id")
	list, err := h.svc.ListConversations(userID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Daftar percakapan", list)
}

type CreateConversationRequest struct {
	Title       string `json:"title"`
	Mode        string `json:"mode"`        // curhat, casual, perspective, night
	Personality string `json:"personality"` // pendengar, santai, bijak, motivator
}

func (h *ChatHandler) CreateConversation(c *gin.Context) {
	userID := c.GetString("user_id")
	var req CreateConversationRequest
	_ = c.ShouldBindJSON(&req)

	conv, err := h.svc.CreateConversation(userID, req.Title, req.Mode, req.Personality)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Percakapan baru dibuat", conv)
}

func (h *ChatHandler) GetConversation(c *gin.Context) {
	userID := c.GetString("user_id")
	convID := c.Param("id")

	conv, msgs, err := h.svc.GetConversation(convID, userID)
	if err != nil {
		response.NotFound(c, "Percakapan tidak ditemukan")
		return
	}

	response.Success(c, http.StatusOK, "Detail percakapan", gin.H{
		"conversation": conv,
		"messages":     msgs,
	})
}

func (h *ChatHandler) DeleteConversation(c *gin.Context) {
	userID := c.GetString("user_id")
	convID := c.Param("id")

	if err := h.svc.DeleteConversation(convID, userID); err != nil {
		response.NotFound(c, "Percakapan tidak ditemukan")
		return
	}

	response.Success(c, http.StatusOK, "Percakapan berhasil dihapus", nil)
}

type SendMessageRequest struct {
	Content     string `json:"content" binding:"required"`
	Mode        string `json:"mode,omitempty"`
	Personality string `json:"personality,omitempty"`
	Nickname    string `json:"nickname,omitempty"`
}

// SSE Streaming message endpoint
func (h *ChatHandler) StreamMessage(c *gin.Context) {
	userID := c.GetString("user_id")
	convID := c.Param("id")

	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Pesan tidak boleh kosong")
		return
	}

	chunkCh, err := h.svc.StreamMessage(c.Request.Context(), userID, convID, req.Content, req.Mode, req.Personality, req.Nickname)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	// Setup Server-Sent Events headers
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Transfer-Encoding", "chunked")

	c.Stream(func(w io.Writer) bool {
		chunk, ok := <-chunkCh
		if !ok {
			fmt.Fprintf(w, "data: [DONE]\n\n")
			return false
		}

		data, _ := json.Marshal(chunk)
		fmt.Fprintf(w, "data: %s\n\n", string(data))
		return true
	})
}

// ==================== MOOD HANDLER ====================

type MoodHandler struct {
	svc *service.MoodService
}

type LogMoodRequest struct {
	Mood int    `json:"mood" binding:"required,min=1,max=5"`
	Note string `json:"note"`
}

func (h *MoodHandler) LogMood(c *gin.Context) {
	userID := c.GetString("user_id")
	var req LogMoodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	mood, err := h.svc.LogMood(userID, req.Mood, req.Note)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Mood hari ini berhasil dicatat", mood)
}

func (h *MoodHandler) GetMoods(c *gin.Context) {
	userID := c.GetString("user_id")
	list, err := h.svc.GetMoods(userID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Riwayat mood", list)
}

func (h *MoodHandler) GetSummary(c *gin.Context) {
	userID := c.GetString("user_id")
	summary, err := h.svc.GetSummary(userID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Ringkasan mood mingguan", summary)
}

// ==================== MEMORY HANDLER ====================

type MemoryHandler struct {
	svc *service.MemoryService
}

func (h *MemoryHandler) GetMemories(c *gin.Context) {
	userID := c.GetString("user_id")
	list, err := h.svc.GetMemories(userID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Daftar memori AI", list)
}

type AddMemoryRequest struct {
	Content  string `json:"content" binding:"required"`
	Category string `json:"category"`
}

func (h *MemoryHandler) AddMemory(c *gin.Context) {
	userID := c.GetString("user_id")
	var req AddMemoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	mem, err := h.svc.AddMemory(userID, req.Content, req.Category)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Memori baru disimpan", mem)
}

func (h *MemoryHandler) DeleteMemory(c *gin.Context) {
	userID := c.GetString("user_id")
	memID := c.Param("id")

	_ = h.svc.DeleteMemory(userID, memID)
	response.Success(c, http.StatusOK, "Memori dihapus", nil)
}

func (h *MemoryHandler) ClearAll(c *gin.Context) {
	userID := c.GetString("user_id")
	_ = h.svc.ClearAll(userID)
	response.Success(c, http.StatusOK, "Semua memori dihapus", nil)
}

// ==================== REFLECTION HANDLER ====================

type ReflectionHandler struct {
	svc *service.ReflectionService
}

func (h *ReflectionHandler) GetQuestion(c *gin.Context) {
	q := h.svc.GetTodayQuestion()
	response.Success(c, http.StatusOK, "Pertanyaan refleksi hari ini", gin.H{"question": q})
}

type SubmitReflectionRequest struct {
	Question string `json:"question" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

func (h *ReflectionHandler) SubmitAnswer(c *gin.Context) {
	userID := c.GetString("user_id")
	var req SubmitReflectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	ref, err := h.svc.SubmitAnswer(userID, req.Question, req.Answer)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Refleksi tersimpan", ref)
}

func (h *ReflectionHandler) GetHistory(c *gin.Context) {
	userID := c.GetString("user_id")
	history, err := h.svc.GetHistory(userID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Riwayat refleksi", history)
}
