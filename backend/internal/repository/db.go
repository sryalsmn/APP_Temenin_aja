package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"temenin-backend/internal/model"
)

type Repository struct {
	db       *sql.DB
	isMemory bool
	mu       sync.RWMutex

	// In-memory tables for local development without external PostgreSQL
	memUsers         map[string]*model.User
	memProfiles      map[string]*model.Profile
	memConversations map[string]*model.Conversation
	memMessages      map[string][]*model.Message
	memMemories      map[string][]*model.Memory
	memMoods         map[string][]*model.Mood
	memReflections   map[string][]*model.Reflection
	memUserSettings  map[string]*model.UserSettings
	memNotifSettings map[string]*model.NotificationSettings
}

func NewRepository(databaseURL string) (*Repository, error) {
	repo := &Repository{
		isMemory:         true,
		memUsers:         make(map[string]*model.User),
		memProfiles:      make(map[string]*model.Profile),
		memConversations: make(map[string]*model.Conversation),
		memMessages:      make(map[string][]*model.Message),
		memMemories:      make(map[string][]*model.Memory),
		memMoods:         make(map[string][]*model.Mood),
		memReflections:   make(map[string][]*model.Reflection),
		memUserSettings:  make(map[string]*model.UserSettings),
		memNotifSettings: make(map[string]*model.NotificationSettings),
	}

	if databaseURL != "" {
		db, err := sql.Open("postgres", databaseURL)
		if err == nil && db.Ping() == nil {
			repo.db = db
			repo.isMemory = false
			fmt.Println(" Connected to PostgreSQL database successfully")
			return repo, nil
		}
	}

	fmt.Println(" Running with in-memory repository (PostgreSQL connection bypassed/offline)")
	return repo, nil
}

// User methods
func (r *Repository) CreateUser(email, passwordHash string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user := &model.User{
		ID:           uuid.New().String(),
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	r.memUsers[user.ID] = user
	r.memProfiles[user.ID] = &model.Profile{
		ID:        uuid.New().String(),
		UserID:    user.ID,
		Nickname:  "Teman",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	r.memUserSettings[user.ID] = &model.UserSettings{
		ID:              uuid.New().String(),
		UserID:          user.ID,
		MemoryEnabled:   true,
		Theme:           "system",
		DefaultChatMode: "curhat",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	r.memNotifSettings[user.ID] = &model.NotificationSettings{
		ID:                 uuid.New().String(),
		UserID:             user.ID,
		DailyReminder:      true,
		ReminderTime:       "20:00",
		ReflectionReminder: true,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	return user, nil
}

func (r *Repository) GetUserByEmail(email string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, u := range r.memUsers {
		if u.Email == email && u.DeletedAt == nil {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *Repository) GetUserByID(id string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if u, ok := r.memUsers[id]; ok && u.DeletedAt == nil {
		return u, nil
	}
	return nil, errors.New("user not found")
}

func (r *Repository) DeleteUser(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if u, ok := r.memUsers[id]; ok {
		now := time.Now()
		u.DeletedAt = &now
		delete(r.memProfiles, id)
		delete(r.memConversations, id)
		delete(r.memMemories, id)
		delete(r.memMoods, id)
		delete(r.memReflections, id)
		return nil
	}
	return errors.New("user not found")
}

// Profile methods
func (r *Repository) GetProfile(userID string) (*model.Profile, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if p, ok := r.memProfiles[userID]; ok {
		return p, nil
	}
	return &model.Profile{
		ID:        uuid.New().String(),
		UserID:    userID,
		Nickname:  "Teman",
		CreatedAt: time.Now(),
	}, nil
}

func (r *Repository) UpdateProfile(userID, nickname, avatarURL string) (*model.Profile, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	p, ok := r.memProfiles[userID]
	if !ok {
		p = &model.Profile{
			ID:        uuid.New().String(),
			UserID:    userID,
			CreatedAt: time.Now(),
		}
		r.memProfiles[userID] = p
	}
	if nickname != "" {
		p.Nickname = nickname
	}
	if avatarURL != "" {
		p.AvatarURL = avatarURL
	}
	p.UpdatedAt = time.Now()
	return p, nil
}

// Conversation methods
func (r *Repository) CreateConversation(userID, title, mode, personality string) (*model.Conversation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if title == "" {
		title = "Percakapan Baru"
	}
	if mode == "" {
		mode = "curhat"
	}
	if personality == "" {
		personality = "pendengar"
	}

	conv := &model.Conversation{
		ID:          uuid.New().String(),
		UserID:      userID,
		Title:       title,
		Mode:        mode,
		Personality: personality,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	r.memConversations[conv.ID] = conv
	return conv, nil
}

func (r *Repository) GetConversations(userID string) ([]*model.Conversation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []*model.Conversation
	for _, c := range r.memConversations {
		if c.UserID == userID {
			result = append(result, c)
		}
	}
	return result, nil
}

func (r *Repository) GetConversationByID(id, userID string) (*model.Conversation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if c, ok := r.memConversations[id]; ok {
		return c, nil
	}

	// Auto-provision conversation if not found
	conv := &model.Conversation{
		ID:          id,
		UserID:      userID,
		Title:       "Obrolan Temenin",
		Mode:        "curhat",
		Personality: "pendengar",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	r.memConversations[id] = conv
	return conv, nil
}

func (r *Repository) DeleteConversation(id, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if c, ok := r.memConversations[id]; ok && c.UserID == userID {
		delete(r.memConversations, id)
		delete(r.memMessages, id)
		return nil
	}
	return errors.New("conversation not found")
}

// Message methods
func (r *Repository) CreateMessage(conversationID, role, content string) (*model.Message, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	msg := &model.Message{
		ID:             uuid.New().String(),
		ConversationID: conversationID,
		Role:           role,
		Content:        content,
		CreatedAt:      time.Now(),
	}
	r.memMessages[conversationID] = append(r.memMessages[conversationID], msg)
	return msg, nil
}

func (r *Repository) GetMessages(conversationID string) ([]*model.Message, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.memMessages[conversationID], nil
}

// Memory methods
func (r *Repository) CreateMemory(userID, content, category string, importance int) (*model.Memory, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	mem := &model.Memory{
		ID:         uuid.New().String(),
		UserID:     userID,
		Content:    content,
		Category:   category,
		Importance: importance,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	r.memMemories[userID] = append(r.memMemories[userID], mem)
	return mem, nil
}

func (r *Repository) GetMemories(userID string) ([]*model.Memory, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.memMemories[userID], nil
}

func (r *Repository) DeleteMemory(userID, memoryID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	list := r.memMemories[userID]
	var updated []*model.Memory
	for _, m := range list {
		if m.ID != memoryID {
			updated = append(updated, m)
		}
	}
	r.memMemories[userID] = updated
	return nil
}

func (r *Repository) DeleteAllMemories(userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.memMemories[userID] = []*model.Memory{}
	return nil
}

// Mood methods
func (r *Repository) CreateMood(userID string, moodVal int, note string) (*model.Mood, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	m := &model.Mood{
		ID:        uuid.New().String(),
		UserID:    userID,
		Mood:      moodVal,
		Note:      note,
		CreatedAt: time.Now(),
	}
	r.memMoods[userID] = append(r.memMoods[userID], m)
	return m, nil
}

func (r *Repository) GetMoods(userID string) ([]*model.Mood, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.memMoods[userID], nil
}

// Reflection methods
func (r *Repository) CreateReflection(userID, question, answer string) (*model.Reflection, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	ref := &model.Reflection{
		ID:        uuid.New().String(),
		UserID:    userID,
		Question:  question,
		Answer:    answer,
		CreatedAt: time.Now(),
	}
	r.memReflections[userID] = append(r.memReflections[userID], ref)
	return ref, nil
}

func (r *Repository) GetReflections(userID string) ([]*model.Reflection, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.memReflections[userID], nil
}

// Settings methods
func (r *Repository) GetUserSettings(userID string) (*model.UserSettings, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if s, ok := r.memUserSettings[userID]; ok {
		return s, nil
	}
	return &model.UserSettings{
		UserID:        userID,
		MemoryEnabled: true,
		Theme:         "system",
	}, nil
}

func (r *Repository) UpdateUserSettings(userID string, memoryEnabled *bool, theme, defaultMode string) (*model.UserSettings, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	s, ok := r.memUserSettings[userID]
	if !ok {
		s = &model.UserSettings{
			ID:            uuid.New().String(),
			UserID:        userID,
			MemoryEnabled: true,
			Theme:         "system",
		}
		r.memUserSettings[userID] = s
	}

	if memoryEnabled != nil {
		s.MemoryEnabled = *memoryEnabled
	}
	if theme != "" {
		s.Theme = theme
	}
	if defaultMode != "" {
		s.DefaultChatMode = defaultMode
	}
	s.UpdatedAt = time.Now()
	return s, nil
}
