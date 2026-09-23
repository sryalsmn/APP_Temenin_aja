package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"temenin-backend/internal/ai"
	"temenin-backend/internal/config"
	"temenin-backend/internal/model"
	"temenin-backend/internal/repository"
	"temenin-backend/pkg/jwt"
)

type Services struct {
	Auth       *AuthService
	Chat       *ChatService
	Mood       *MoodService
	Memory     *MemoryService
	Reflection *ReflectionService
}

func NewServices(cfg *config.Config, repo *repository.Repository, aiProv ai.AIProvider) *Services {
	return &Services{
		Auth:       NewAuthService(cfg, repo),
		Chat:       NewChatService(cfg, repo, aiProv),
		Mood:       NewMoodService(repo),
		Memory:     NewMemoryService(repo),
		Reflection: NewReflectionService(repo),
	}
}

// ==================== AUTH SERVICE ====================

type AuthService struct {
	cfg  *config.Config
	repo *repository.Repository
}

func NewAuthService(cfg *config.Config, repo *repository.Repository) *AuthService {
	return &AuthService{cfg: cfg, repo: repo}
}

type AuthResult struct {
	User         *model.User    `json:"user"`
	Profile      *model.Profile `json:"profile"`
	AccessToken  string         `json:"access_token"`
	RefreshToken string         `json:"refresh_token"`
}

func (s *AuthService) Register(email, password, nickname string) (*AuthResult, error) {
	if _, err := s.repo.GetUserByEmail(email); err == nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.CreateUser(email, string(hashed))
	if err != nil {
		return nil, err
	}

	profile, _ := s.repo.UpdateProfile(user.ID, nickname, "")

	access, err := jwt.GenerateAccessToken(user.ID, user.Email, s.cfg.JWTSecret)
	if err != nil {
		return nil, err
	}

	refresh, err := jwt.GenerateRefreshToken(user.ID, user.Email, s.cfg.JWTSecret)
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		User:         user,
		Profile:      profile,
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *AuthService) Login(email, password string) (*AuthResult, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	profile, _ := s.repo.GetProfile(user.ID)

	access, err := jwt.GenerateAccessToken(user.ID, user.Email, s.cfg.JWTSecret)
	if err != nil {
		return nil, err
	}

	refresh, err := jwt.GenerateRefreshToken(user.ID, user.Email, s.cfg.JWTSecret)
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		User:         user,
		Profile:      profile,
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *AuthService) GetMe(userID string) (*model.User, *model.Profile, *model.UserSettings, error) {
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return nil, nil, nil, err
	}
	profile, _ := s.repo.GetProfile(userID)
	settings, _ := s.repo.GetUserSettings(userID)
	return user, profile, settings, nil
}

func (s *AuthService) DeleteAccount(userID string) error {
	return s.repo.DeleteUser(userID)
}

// ==================== CHAT SERVICE ====================

type ChatService struct {
	cfg        *config.Config
	repo       *repository.Repository
	aiProvider ai.AIProvider
}

func NewChatService(cfg *config.Config, repo *repository.Repository, aiProv ai.AIProvider) *ChatService {
	return &ChatService{cfg: cfg, repo: repo, aiProvider: aiProv}
}

func (s *ChatService) ListConversations(userID string) ([]*model.Conversation, error) {
	return s.repo.GetConversations(userID)
}

func (s *ChatService) CreateConversation(userID, title, mode, personality string) (*model.Conversation, error) {
	return s.repo.CreateConversation(userID, title, mode, personality)
}

func (s *ChatService) GetConversation(convID, userID string) (*model.Conversation, []*model.Message, error) {
	conv, err := s.repo.GetConversationByID(convID, userID)
	if err != nil {
		return nil, nil, err
	}
	msgs, err := s.repo.GetMessages(convID)
	return conv, msgs, err
}

func (s *ChatService) DeleteConversation(convID, userID string) error {
	return s.repo.DeleteConversation(convID, userID)
}

type StreamChunk struct {
	Token          string             `json:"token"`
	IsSafetyAlert  bool               `json:"is_safety_alert"`
	CrisisHotlines []ai.CrisisHotline `json:"crisis_hotlines,omitempty"`
	Done           bool               `json:"done"`
}

func (s *ChatService) StreamMessage(ctx context.Context, userID, convID, userContent, reqMode, reqPersonality, reqNickname string) (<-chan StreamChunk, error) {
	conv, err := s.repo.GetConversationByID(convID, userID)
	if err != nil {
		return nil, err
	}

	if reqMode != "" {
		conv.Mode = reqMode
	}
	if reqPersonality != "" {
		conv.Personality = reqPersonality
	}

	// 1. Save user message to database
	_, err = s.repo.CreateMessage(convID, "user", userContent)
	if err != nil {
		return nil, err
	}

	out := make(chan StreamChunk)

	// 2. Perform Safety Crisis Detection Check
	safety := ai.CheckSafety(userContent)
	if safety.IsCrisis {
		go func() {
			defer close(out)
			// Save the safe response to conversation
			s.repo.CreateMessage(convID, "assistant", safety.SafeResponse)

			out <- StreamChunk{
				Token:          safety.SafeResponse,
				IsSafetyAlert:  true,
				CrisisHotlines: safety.CrisisHotlines,
				Done:           true,
			}
		}()
		return out, nil
	}

	// 3. Load context: profile, memories, recent messages
	profile, _ := s.repo.GetProfile(userID)
	settings, _ := s.repo.GetUserSettings(userID)

	userNick := "kamu"
	if profile != nil && profile.Nickname != "" {
		userNick = profile.Nickname
	}
	if reqNickname != "" {
		userNick = reqNickname
	}

	var memories []string
	if settings == nil || settings.MemoryEnabled {
		memObjs, _ := s.repo.GetMemories(userID)
		for _, m := range memObjs {
			memories = append(memories, m.Content)
		}
	}

	existingMsgs, _ := s.repo.GetMessages(convID)
	var aiMsgs []ai.ChatMessage
	for _, m := range existingMsgs {
		aiMsgs = append(aiMsgs, ai.ChatMessage{Role: m.Role, Content: m.Content})
	}

	// Load latest user mood if available
	recentMoods, _ := s.repo.GetMoods(userID)
	var latestMoodStr string
	if len(recentMoods) > 0 {
		m := recentMoods[0]
		moodLabels := map[int]string{
			1: "Sangat Sedih / Down (1/5)",
			2: "Kurang Baik / Lelah (2/5)",
			3: "Biasa Saja / Netral (3/5)",
			4: "Cukup Baik / Tenang (4/5)",
			5: "Sangat Senang / Bersemangat (5/5)",
		}
		notePart := ""
		if m.Note != "" {
			notePart = fmt.Sprintf(" - Catatan: \"%s\"", m.Note)
		}
		latestMoodStr = fmt.Sprintf("%s%s", moodLabels[m.Mood], notePart)
	}

	aiReq := ai.AIRequest{
		Personality:  conv.Personality,
		Mode:         conv.Mode,
		UserNickname: userNick,
		UserMood:     latestMoodStr,
		Memories:     memories,
		Messages:     aiMsgs,
	}

	tokensCh, errCh := s.aiProvider.GenerateStream(ctx, aiReq)

	go func() {
		defer close(out)
		var fullAssistantReply string

		for {
			select {
			case <-ctx.Done():
				return
			case err, ok := <-errCh:
				if !ok {
					errCh = nil
					continue
				}
				if err != nil {
					errMsg := err.Error()
					if strings.Contains(errMsg, "credit balance is too low") {
						out <- StreamChunk{
							Token: "⚠️ [Anthropic Claude]: Saldo kredit API kamu di console.anthropic.com saat ini masih $0. Silakan isi billing Claude di console.anthropic.com/settings/billing, atau beralih ke Google Gemini (gratis tanpa kartu kredit) di file backend/.env.\n\n",
						}
					}
					// Fallback to dynamic empathetic response so user is never stuck
					fallbackText := fmt.Sprintf("Gue dengerin dan paham banget apa yang kamu rasain, %s. Kadang mengeluarkan isi pikiran itu hal yang paling melegakan. Mau cerita lebih detail lagi? Gue siap nemenin.", profile.Nickname)
					for _, w := range strings.Split(fallbackText, " ") {
						out <- StreamChunk{Token: w + " "}
					}
					s.repo.CreateMessage(convID, "assistant", fallbackText)
				}
				return
			case token, ok := <-tokensCh:
				if !ok {
					// Done streaming, save full message
					if fullAssistantReply != "" {
						s.repo.CreateMessage(convID, "assistant", fullAssistantReply)
					}
					out <- StreamChunk{Done: true}
					return
				}
				fullAssistantReply += token
				out <- StreamChunk{Token: token, Done: false}
			}
		}
	}()

	return out, nil
}

// ==================== MOOD SERVICE ====================

type MoodService struct {
	repo *repository.Repository
}

func NewMoodService(repo *repository.Repository) *MoodService {
	return &MoodService{repo: repo}
}

func (s *MoodService) LogMood(userID string, moodVal int, note string) (*model.Mood, error) {
	if moodVal < 1 || moodVal > 5 {
		return nil, errors.New("mood value must be between 1 and 5")
	}
	return s.repo.CreateMood(userID, moodVal, note)
}

func (s *MoodService) GetMoods(userID string) ([]*model.Mood, error) {
	return s.repo.GetMoods(userID)
}

type MoodSummary struct {
	TotalLogs  int            `json:"total_logs"`
	Counts     map[int]int    `json:"counts"`
	Insight    string         `json:"insight"`
	RecentLogs []*model.Mood  `json:"recent_logs"`
}

func (s *MoodService) GetSummary(userID string) (*MoodSummary, error) {
	moods, err := s.repo.GetMoods(userID)
	if err != nil {
		return nil, err
	}

	counts := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}
	for _, m := range moods {
		counts[m.Mood]++
	}

	var insight string
	if counts[4]+counts[5] > counts[1]+counts[2] {
		insight = "Minggu ini mood kamu cenderung positif dan cerah. Pertahankan ritme yang bikin hatimu nyaman ya!"
	} else if counts[1]+counts[2] > 0 {
		insight = "Tampaknya beberapa hari ini terasa cukup berat. Jangan ragu buat istirahat dan curhat kapan pun kamu butuh ruang."
	} else {
		insight = "Catat mood harianmu untuk melihat pola perasaan dan perkembangan diri kamu sepanjang minggu."
	}

	return &MoodSummary{
		TotalLogs:  len(moods),
		Counts:     counts,
		Insight:    insight,
		RecentLogs: moods,
	}, nil
}

// ==================== MEMORY SERVICE ====================

type MemoryService struct {
	repo *repository.Repository
}

func NewMemoryService(repo *repository.Repository) *MemoryService {
	return &MemoryService{repo: repo}
}

func (s *MemoryService) GetMemories(userID string) ([]*model.Memory, error) {
	return s.repo.GetMemories(userID)
}

func (s *MemoryService) AddMemory(userID, content, category string) (*model.Memory, error) {
	if category == "" {
		category = "personal"
	}
	return s.repo.CreateMemory(userID, content, category, 1)
}

func (s *MemoryService) DeleteMemory(userID, memoryID string) error {
	return s.repo.DeleteMemory(userID, memoryID)
}

func (s *MemoryService) ClearAll(userID string) error {
	return s.repo.DeleteAllMemories(userID)
}

// ==================== REFLECTION SERVICE ====================

type ReflectionService struct {
	repo *repository.Repository
}

func NewReflectionService(repo *repository.Repository) *ReflectionService {
	return &ReflectionService{repo: repo}
}

var DailyQuestions = []string{
	"Apa satu hal kecil yang membuatmu tersenyum hari ini?",
	"Apa hal yang paling memenuhi pikiranmu seharian ini?",
	"Apa satu hal yang berhasil kamu selesaikan dan patut kamu syukuri?",
	"Kalau harimu bisa diulang, bagian mana yang ingin kamu jalani lebih santai?",
	"Siapa orang atau momen yang bikin harimu terasa lebih hangat?",
}

func (s *ReflectionService) GetTodayQuestion() string {
	dayOfYear := time.Now().YearDay()
	idx := dayOfYear % len(DailyQuestions)
	return DailyQuestions[idx]
}

func (s *ReflectionService) SubmitAnswer(userID, question, answer string) (*model.Reflection, error) {
	return s.repo.CreateReflection(userID, question, answer)
}

func (s *ReflectionService) GetHistory(userID string) ([]*model.Reflection, error) {
	return s.repo.GetReflections(userID)
}
