package model

import (
	"time"
)

type User struct {
	ID           string     `json:"id"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

type Profile struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Nickname  string    `json:"nickname"`
	AvatarURL string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Conversation struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Title       string    `json:"title"`
	Mode        string    `json:"mode"`        // vent, casual, perspective, night
	Personality string    `json:"personality"` // pendengar, santai, bijak, motivator
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Message struct {
	ID             string    `json:"id"`
	ConversationID string    `json:"conversation_id"`
	Role           string    `json:"role"` // user, assistant, system
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
}

type Memory struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Content    string    `json:"content"`
	Category   string    `json:"category"` // preference, personal, interest, goal, other
	Importance int       `json:"importance"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Mood struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Mood      int       `json:"mood"` // 1: very sad, 2: sad, 3: neutral, 4: good, 5: very good
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}

type Reflection struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	CreatedAt time.Time `json:"created_at"`
}

type NotificationSettings struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	DailyReminder      bool      `json:"daily_reminder"`
	ReminderTime       string    `json:"reminder_time"`
	ReflectionReminder bool      `json:"reflection_reminder"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type UserSettings struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	MemoryEnabled   bool      `json:"memory_enabled"`
	Theme           string    `json:"theme"` // light, dark, system
	DefaultChatMode string    `json:"default_chat_mode"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
