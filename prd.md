# PRD — TEMENIN

### AI Companion & Teman Curhat untuk Android

**Version:** 1.0
**Platform:** Android
**Frontend:** React Native + Expo + TypeScript
**Backend:** Go
**Database:** PostgreSQL
**AI Engine:** LLM API melalui backend
**Target:** Google Play Store

---

# 1. Product Overview

## 1.1 Nama Produk

**TEMENIN**

### Tagline

> **"Lagi pengen cerita? Sini, gue dengerin."**

TEMENIN adalah aplikasi Android berbasis AI yang dirancang sebagai **teman ngobrol virtual** untuk pengguna yang ingin bercerita, ngobrol santai, meminta perspektif, atau sekadar memiliki teman ketika sedang sendirian.

TEMENIN bukan aplikasi terapi dan bukan pengganti psikolog, konselor, keluarga, pasangan, atau teman di dunia nyata.

Fokus utama produk:

* Percakapan natural
* Teman curhat
* Teman ngobrol
* Refleksi diri
* Mood tracking
* AI personality
* Memory yang dikontrol pengguna
* Pengalaman chat yang nyaman dan personal

---

# 2. Problem Statement

Banyak orang ingin bercerita tetapi tidak selalu memiliki seseorang yang bisa diajak bicara pada saat tertentu.

Beberapa kondisi yang sering terjadi:

* sedang sendirian
* ingin mengeluarkan isi pikiran
* takut mengganggu teman
* tidak tahu harus cerita kepada siapa
* ingin mendapatkan perspektif lain
* ingin ngobrol ringan
* ingin melakukan refleksi setelah menjalani hari

Namun aplikasi AI sering terasa terlalu formal dan seperti chatbot biasa.

TEMENIN mencoba membuat pengalaman yang lebih:

**natural → hangat → personal → sederhana.**

---

# 3. Product Vision

Menjadi aplikasi companion AI yang terasa seperti **tempat aman untuk bercerita dan berbicara**, tanpa membuat pengguna bergantung secara tidak sehat kepada AI.

---

# 4. Target User

## Primary User

Usia:

**17–35 tahun**

Target awal:

* mahasiswa
* pekerja muda
* freelancer
* developer
* pengguna yang sering bekerja sendiri
* pengguna yang suka journaling
* pengguna yang ingin ngobrol dengan AI

## User Characteristics

Pengguna cenderung:

* aktif menggunakan smartphone
* familiar dengan WhatsApp/Telegram/ChatGPT
* menyukai percakapan casual
* membutuhkan tempat untuk menuangkan pikiran
* tertarik dengan AI
* menyukai aplikasi dengan UI modern

---

# 5. Unique Value Proposition

TEMENIN bukan sekadar chatbot.

Konsep utamanya:

> **AI yang bisa menemani percakapan sehari-hari dengan gaya komunikasi yang bisa dipilih pengguna.**

Contohnya:

**Pengguna:**

> "Hari ini kerjaan gue banyak banget."

TEMENIN:

> "Waduh, kedengerannya capek banget 😭
> Yang paling bikin berat bagian mana?"

Bukan:

> "Saya memahami bahwa Anda sedang mengalami tekanan akibat banyaknya pekerjaan."

Bahasa harus terasa natural.

---

# 6. Core Features

## MVP Features

MVP wajib memiliki:

1. Onboarding
2. Login/Register
3. Home
4. AI Chat
5. Personality
6. Chat Mode
7. Conversation History
8. Basic AI Memory
9. Mood Tracking
10. Daily Reflection
11. Profile
12. Settings
13. Notification
14. Safety System
15. Delete Account/Data
16. Dark Mode

---

# 7. App Navigation

Gunakan bottom navigation.

```text
┌─────────────────────────────────────┐
│                                     │
│              CONTENT                │
│                                     │
├─────────────────────────────────────┤
│  Home     Chat      Mood    Profile │
└─────────────────────────────────────┘
```

Navigation:

### Home

Dashboard dan shortcut.

### Chat

Daftar percakapan + mulai percakapan baru.

### Mood

Mood tracking dan reflection.

### Profile

Profile, personality, memory, notification, theme, privacy.

---

# 8. Onboarding

Saat pertama kali membuka aplikasi.

## Screen 1 — Welcome

Logo TEMENIN.

Headline:

> **"Kadang kita cuma butuh seseorang untuk mendengarkan."**

Subheadline:

> TEMENIN siap menemani kamu ngobrol kapan saja.

CTA:

**Mulai Sekarang**

---

## Screen 2 — Nickname

Pertanyaan:

> "Kamu biasanya dipanggil apa?"

Input:

```text
Nama panggilan
```

CTA:

**Lanjut**

---

## Screen 3 — Tujuan

> "Biasanya kamu datang ke sini buat apa?"

Pilihan:

* Curhat
* Ngobrol santai
* Minta perspektif
* Refleksi diri
* Menemani malam

Bisa memilih lebih dari satu.

---

## Screen 4 — Personality

> "Kamu lebih nyaman ngobrol sama yang seperti apa?"

Pilihan:

### 🫂 Pendengar

Lebih banyak mendengarkan.

### 😎 Santai

Casual dan banyak bercanda.

### 🧠 Bijak

Lebih reflektif.

### 🔥 Motivator

Lebih memberikan dorongan.

---

## Screen 5 — Privacy

Informasikan:

* percakapan diproses oleh AI
* data akun disimpan secara aman
* pengguna dapat menghapus data
* memory dapat dimatikan
* TEMENIN bukan tenaga profesional kesehatan mental

CTA:

**Saya Mengerti**

---

# 9. Home Screen

Home menjadi halaman utama.

Contoh:

```text
Good evening, Surya 👋

Gimana kabarmu hari ini?

┌─────────────────────────────┐
│ 🙂 Lumayan baik             │
│                             │
│ Tap untuk check-in mood     │
└─────────────────────────────┘

Mau ngobrol tentang apa?

┌──────────────┐ ┌──────────────┐
│ 🫂 Curhat    │ │ 😎 Santai    │
└──────────────┘ └──────────────┘

┌──────────────┐ ┌──────────────┐
│ 🧠 Perspektif│ │ 🌙 Night Talk│
└──────────────┘ └──────────────┘

──────────────────────────────

💭 Lanjutkan percakapan

"Kerjaan hari ini..."

──────────────────────────────

✨ Daily Reflection

"Apa satu hal kecil yang membuatmu
tersenyum hari ini?"
```

---

# 10. Chat System

Ini adalah fitur utama TEMENIN.

## Chat UI

Konsep seperti aplikasi messenger modern.

User bubble:

```text
Gue hari ini capek banget.
```

AI:

```text
Kelihatannya hari ini lumayan berat ya.

Mau cerita bagian yang paling bikin
capeknya apa?
```

---

# 11. Chat Modes

## Mode 1 — Curhat

AI fokus:

* mendengarkan
* memberikan respons empatik
* tidak langsung memberi solusi
* bertanya secara natural

---

## Mode 2 — Ngobrol Santai

Untuk:

* random conversation
* jokes
* cerita
* hobi
* film
* game
* kehidupan sehari-hari

---

## Mode 3 — Perspektif

Pengguna meminta sudut pandang.

Contoh:

> "Menurut kamu gue harus ngomong ke dia atau enggak?"

AI:

* memahami konteks
* memberikan beberapa perspektif
* menjelaskan trade-off
* tidak memutuskan hidup pengguna

---

## Mode 4 — Night Talk

Mode percakapan malam.

Tone:

* lebih tenang
* pendek
* hangat
* tidak terlalu ramai

Contoh:

> "Hari ini sudah kamu lewati.
> Ada sesuatu yang masih kepikiran sebelum tidur?"

---

# 12. AI Personality

Pengguna dapat mengganti personality.

## Pendengar

Tone:

> Empathetic, patient, reflective.

## Santai

Tone:

> Casual, friendly, playful.

## Bijak

Tone:

> Calm, thoughtful, balanced.

## Motivator

Tone:

> Encouraging, energetic, constructive.

Personality tidak boleh mengubah safety rules.

---

# 13. AI Memory

TEMENIN memiliki memory sederhana.

Contoh:

Pengguna:

> "Gue suka main Free Fire."

AI dapat menyimpan:

```text
User likes playing Free Fire.
```

Kemudian beberapa hari berikutnya:

> "Kemarin katanya mau push rank lagi. Jadi main?"

---

## Memory Controls

User dapat:

* melihat memory
* menghapus memory tertentu
* menghapus semua memory
* mematikan memory

Contoh:

```text
AI Memory

☑ Remember things about me

Memories:

• Kamu suka bermain game
• Kamu bekerja sebagai developer

[Manage Memory]
```

---

# 14. AI Context System

Backend harus menyusun context:

```text
System Prompt
+
Personality
+
Chat Mode
+
User Profile
+
Relevant Memories
+
Recent Conversation
+
Safety Rules
```

Contoh:

```text
SYSTEM

You are TEMENIN.

You are a friendly AI companion.

Your purpose is to:
- listen
- converse naturally
- help users reflect
- provide balanced perspectives

Never claim to be a human.
Never claim to be a therapist.
Never encourage emotional dependency.
Never manipulate the user.
```

---

# 15. Conversation History

User dapat melihat percakapan sebelumnya.

Contoh:

```text
Chats

Today
────────────────

Kerjaan lagi berat
09:32

Ngobrol malam
00:14

Yesterday
────────────────

Tentang kuliah

Mau resign?
```

User dapat:

* membuka chat
* rename conversation
* delete conversation

---

# 16. Mood Tracking

Halaman:

**"Gimana perasaanmu hari ini?"**

Mood:

```text
😭
😔
😐
🙂
😄
```

User memilih mood.

Optional:

```text
Apa yang membuatmu merasa seperti ini?

[_____________________]
```

---

# 17. Mood Dashboard

Tampilkan:

```text
Your Mood

This Week

😄  🙂  😐  😔  🙂
```

Statistik:

```text
Happy       3 days
Neutral     2 days
Sad         1 day
```

Insight sederhana:

> "Minggu ini mood kamu terlihat lebih stabil dibanding minggu lalu."

Hindari diagnosis psikologis.

---

# 18. Daily Reflection

Setiap hari satu pertanyaan.

Contoh:

> "Apa hal kecil yang kamu syukuri hari ini?"

Pertanyaan lainnya:

* Apa yang paling membuatmu senang hari ini?
* Apa yang ingin kamu lakukan lebih baik besok?
* Apa hal yang paling memenuhi pikiranmu?
* Apa satu hal yang berhasil kamu selesaikan?

Jawaban disimpan sebagai private reflection.

---

# 19. Notification

Push notification bersifat optional.

Contoh:

### Morning

> "Pagi 👋 Semoga harimu berjalan dengan baik."

### Evening

> "Hari ini gimana? Mau cerita sebentar?"

### Reflection

> "Ada waktu 1 menit buat refleksi hari ini?"

User dapat mengatur:

```text
Notifications

☑ Daily Reminder

Time
20:00

☑ Reflection Reminder
```

Tidak boleh menggunakan notification yang manipulatif seperti:

> "TEMENIN kangen kamu."

atau:

> "Kok kamu nggak datang hari ini?"

---

# 20. Safety System

Karena aplikasi berkaitan dengan curhat dan emosi, safety harus menjadi bagian inti.

## AI Tidak Boleh

AI tidak boleh:

* mengaku sebagai manusia
* mengaku sebagai psikolog
* mendiagnosis penyakit
* memberikan instruksi berbahaya
* mendorong self-harm
* mendorong kekerasan
* mendorong ketergantungan emosional
* mengatakan hanya AI yang memahami pengguna
* meminta pengguna menjauhi manusia nyata
* memanipulasi pengguna agar terus menggunakan aplikasi

---

# 21. Crisis Detection

Jika pengguna mengindikasikan kondisi berbahaya, sistem masuk ke **Safety Response**.

Contoh sinyal:

* keinginan menyakiti diri
* keinginan bunuh diri
* ancaman terhadap orang lain
* kondisi darurat

AI harus:

1. merespons dengan empati
2. tidak menghakimi
3. mendorong pengguna mencari bantuan manusia
4. menyarankan layanan darurat yang relevan
5. mendorong pengguna untuk tidak sendirian jika berada dalam bahaya langsung

Jangan memberikan instruksi teknis terkait self-harm.

---

# 22. Profile

Profile screen:

```text
┌───────────────────────────┐
│        👤                 │
│       Surya               │
│                           │
│   Edit Profile            │
└───────────────────────────┘

Your Companion
────────────────────

Personality
🫂 Pendengar

Memory
On

Theme
System

Notifications
On

────────────────────

Privacy & Security

Delete Account
Logout
```

---

# 23. Settings

## Account

* Edit profile
* Change nickname
* Logout
* Delete account

## Companion

* Personality
* Memory
* Default chat mode

## Appearance

* Light
* Dark
* System

## Notifications

* Enable/disable
* Reminder time

## Privacy

* Privacy policy
* Terms
* Data deletion

## About

* Version
* Contact
* Feedback

---

# 24. Design System

Visual direction:

**Warm + Modern + Premium + Calm**

Jangan menggunakan UI chatbot generik.

## Color

Primary:

```text
Deep Violet
#6C5CE7
```

Secondary:

```text
Soft Purple
#A78BFA
```

Background:

```text
#F8F7FC
```

Dark:

```text
#121118
```

Accent:

```text
#F4A7C1
```

---

# 25. UI Style

Gunakan:

* rounded card
* large typography
* soft shadows
* subtle gradients
* glass-like elements secukupnya
* smooth transitions
* micro interaction
* animated typing indicator
* animated mood selector

Border radius:

```text
12–24px
```

---

# 26. Mobile UX

Karena Android menjadi platform utama:

* touch target minimal ±44px
* keyboard-aware chat
* safe area
* Android back button handling
* dark mode
* responsive berbagai ukuran layar
* loading skeleton
* offline state
* network error state

Chat input harus tetap nyaman ketika keyboard terbuka.

---

# 27. Animation

Gunakan React Native Reanimated.

Contoh:

### Chat Message

Fade + slide up.

### Mood

Scale animation saat dipilih.

### Personality

Card selection animation.

### AI Typing

Three-dot animation.

### Page Transition

Subtle slide/fade.

Animasi jangan berlebihan.

---

# 28. Frontend Technology

## React Native

Framework utama aplikasi.

## Expo

Digunakan untuk:

* development
* native APIs
* notifications
* secure storage
* build Android
* EAS

## TypeScript

Wajib.

## Expo Router

Untuk navigation.

## Zustand

Global state.

State:

```text
auth
user
settings
chat
companion
```

## TanStack Query

Untuk:

* API request
* caching
* loading state
* refetch
* mutation

## React Native Reanimated

Untuk animation.

## Expo SecureStore

Untuk token sensitif.

---

# 29. Frontend Folder Structure

```text
temenin/
│
├── app/
│   ├── _layout.tsx
│   │
│   ├── index.tsx
│   │
│   ├── onboarding/
│   │   ├── welcome.tsx
│   │   ├── nickname.tsx
│   │   ├── purpose.tsx
│   │   ├── personality.tsx
│   │   └── privacy.tsx
│   │
│   ├── auth/
│   │   ├── login.tsx
│   │   └── register.tsx
│   │
│   ├── (tabs)/
│   │   ├── _layout.tsx
│   │   ├── home.tsx
│   │   ├── chat.tsx
│   │   ├── mood.tsx
│   │   └── profile.tsx
│   │
│   ├── conversation/
│   │   └── [id].tsx
│   │
│   ├── reflection/
│   │   └── index.tsx
│   │
│   └── settings/
│       ├── index.tsx
│       ├── personality.tsx
│       ├── memory.tsx
│       ├── notification.tsx
│       ├── privacy.tsx
│       └── about.tsx
│
├── components/
│   ├── chat/
│   ├── mood/
│   ├── home/
│   ├── common/
│   └── ui/
│
├── services/
│   ├── api.ts
│   ├── auth.ts
│   ├── chat.ts
│   ├── mood.ts
│   └── memory.ts
│
├── stores/
│   ├── authStore.ts
│   ├── chatStore.ts
│   ├── userStore.ts
│   └── settingsStore.ts
│
├── hooks/
│
├── types/
│
├── constants/
│
├── utils/
│
├── assets/
│   ├── images/
│   ├── icons/
│   └── fonts/
│
└── package.json
```

---

# 30. Backend Architecture

Backend menggunakan:

```text
Go
Gin
PostgreSQL
JWT
Redis (optional)
AI API
```

Architecture:

```text
React Native
      │
      │ HTTPS
      ▼
┌───────────────┐
│   Go API      │
│    Gin        │
└───────┬───────┘
        │
 ┌──────┼───────────┐
 │      │           │
 ▼      ▼           ▼
Auth   PostgreSQL   AI API
 │
 ▼
Redis
(optional)
```

AI API **tidak boleh dipanggil langsung dari aplikasi Android**.

API key AI harus berada di server.

---

# 31. Backend Folder Structure

```text
backend/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── middleware/
│   ├── model/
│   ├── ai/
│   └── config/
│
├── migrations/
│
├── pkg/
│   ├── jwt/
│   ├── logger/
│   └── response/
│
├── .env
├── go.mod
└── Dockerfile
```

---

# 32. API Design

Base:

```text
/api/v1
```

---

## Authentication

```http
POST /auth/register
POST /auth/login
POST /auth/refresh
POST /auth/logout
```

---

## User

```http
GET /me
PATCH /me
DELETE /me
```

---

## Companion

```http
GET /companion/settings
PATCH /companion/settings
```

---

## Chat

```http
GET /conversations
POST /conversations
GET /conversations/:id
DELETE /conversations/:id
POST /conversations/:id/messages
```

Streaming:

```http
GET /conversations/:id/stream
```

atau menggunakan SSE.

---

## Memory

```http
GET /memories
POST /memories
DELETE /memories/:id
DELETE /memories
PATCH /memory/settings
```

---

## Mood

```http
POST /moods
GET /moods
GET /moods/summary
```

---

## Reflection

```http
GET /reflections/question
POST /reflections
GET /reflections
```

---

# 33. Database

PostgreSQL.

Core tables:

```text
users
profiles
refresh_tokens
conversations
messages
memories
moods
reflections
notification_settings
user_settings
```

---

# 34. Users

```text
users
----------------
id
email
password_hash
created_at
updated_at
deleted_at
```

---

# 35. Profiles

```text
profiles
----------------
id
user_id
nickname
avatar_url
created_at
updated_at
```

---

# 36. Conversations

```text
conversations
----------------
id
user_id
title
mode
personality
created_at
updated_at
```

Mode:

```text
vent
casual
perspective
night
```

---

# 37. Messages

```text
messages
----------------
id
conversation_id
role
content
created_at
```

Role:

```text
user
assistant
system
```

---

# 38. Memories

```text
memories
----------------
id
user_id
content
category
importance
created_at
updated_at
```

Category:

```text
preference
personal
interest
goal
other
```

---

# 39. Mood

```text
moods
----------------
id
user_id
mood
note
created_at
```

Mood:

```text
1 = very sad
2 = sad
3 = neutral
4 = good
5 = very good
```

---

# 40. AI Request Flow

Saat user mengirim pesan:

```text
User
 ↓
React Native
 ↓
Go API
 ↓
Authentication
 ↓
Load conversation
 ↓
Load recent messages
 ↓
Load relevant memories
 ↓
Load personality
 ↓
Load chat mode
 ↓
Build AI prompt
 ↓
AI Provider
 ↓
Stream response
 ↓
React Native
 ↓
Display message
```

---

# 41. Streaming Chat

Untuk membuat AI terasa cepat, gunakan streaming.

Contoh:

```text
AI sedang mengetik...

"Kayaknya hari ini..."
        ↓
"Kayaknya hari ini kamu..."
        ↓
"Kayaknya hari ini kamu lumayan capek..."
```

Frontend menerima token secara bertahap.

Pilihan teknologi:

**Server-Sent Events (SSE)**

atau WebSocket.

Untuk MVP:

**SSE direkomendasikan.**

---

# 42. Authentication

Gunakan:

```text
JWT Access Token
+
Refresh Token
```

Access token disimpan secara aman menggunakan:

```text
Expo SecureStore
```

Jangan menyimpan token authentication di plain AsyncStorage.

---

# 43. Security

Wajib:

* HTTPS
* password hashing
* JWT expiration
* refresh token rotation
* rate limiting
* request validation
* SQL injection protection
* input sanitization
* API key hanya di backend
* secure headers
* CORS configuration
* logging
* error handling

---

# 44. Rate Limiting

Untuk mencegah abuse.

Contoh:

```text
Anonymous:
5 requests/minute

Authenticated:
30 requests/minute

AI:
limit berdasarkan user
```

Angka dapat disesuaikan saat production.

---

# 45. AI Cost Control

Karena AI API memiliki biaya, backend harus mengontrol penggunaan.

Implementasi:

* maximum token
* conversation truncation
* context window management
* memory summarization
* rate limit
* daily usage limit
* caching jika relevan

---

# 46. Conversation Context

Jangan selalu mengirim seluruh chat.

Contoh:

```text
Last 20 messages
+
Conversation summary
+
Relevant memory
```

Jika chat panjang:

```text
Old messages
      ↓
Summary
      ↓
Recent messages
```

Ini mengurangi biaya AI.

---

# 47. Free Plan

Contoh MVP monetization:

```text
FREE

20 AI messages/day
4 personalities
Basic memory
Mood tracking
Reflection
Chat history
```

---

# 48. Premium

Tahap setelah MVP.

```text
TEMENIN PLUS

Unlimited chat
Advanced memory
Premium personalities
Advanced insights
Longer context
Custom personality
Priority AI
```

Jangan membangun subscription terlebih dahulu jika MVP belum mendapatkan traction.

---

# 49. Monetization Strategy

Tahap awal:

**Free**

Tujuan:

* mendapatkan user
* menguji retention
* menguji AI quality
* mengetahui fitur yang paling digunakan

Setelah product-market fit mulai terlihat:

**Freemium**

Kemudian:

**Subscription**

---

# 50. Analytics

Track event:

```text
app_open
onboarding_complete
chat_started
message_sent
conversation_created
personality_changed
mood_logged
reflection_completed
memory_enabled
memory_disabled
notification_enabled
premium_clicked
account_deleted
```

Jangan mengumpulkan data sensitif lebih dari yang diperlukan.

---

# 51. Error Handling

Contoh:

### Network error

> "Koneksi kamu sedang bermasalah."

Button:

**Coba Lagi**

### AI error

> "Kayaknya gue lagi sedikit bermasalah 😅
> Coba kirim lagi ya."

### Server error

> "Terjadi kesalahan. Coba beberapa saat lagi."

Jangan menampilkan error teknis seperti:

```text
500 Internal Server Error
```

kepada user.

---

# 52. Loading State

Gunakan skeleton/loading animation.

Chat:

```text
● ● ●
```

Home:

Skeleton cards.

Mood:

Animated selector.

---

# 53. Empty State

Contoh chat:

> "Belum ada cerita hari ini."

CTA:

**Mulai Ngobrol**

History:

> "Belum ada percakapan."

CTA:

**Temui TEMENIN**

---

# 54. Offline Handling

Jika internet terputus:

```text
No Internet Connection

Kamu sedang offline.

Beberapa fitur mungkin tidak tersedia.
```

Mood/reflection yang dapat dilakukan offline bisa disimpan sementara dan disinkronkan ketika online.

---

# 55. Android Requirements

Target:

```text
Android 8.0+
```

atau menyesuaikan minimum SDK yang dipilih saat setup Expo.

Support:

* portrait
* Android back button
* status bar
* edge-to-edge
* keyboard handling
* notification permission

---

# 56. App Icon

Logo:

**T**

atau simbol dua bentuk percakapan yang membentuk hati secara subtle.

Style:

* simple
* recognizable
* tidak terlalu ramai
* cocok untuk launcher Android

---

# 57. Splash Screen

Background:

```text
#6C5CE7
```

Logo:

```text
TEMENIN
```

Animation:

Fade + scale.

Durasi singkat.

---

# 58. Play Store Assets

Wajib disiapkan:

* App icon
* Feature graphic
* Screenshot Android
* Short description
* Full description
* Privacy policy
* Terms
* Data safety information
* Content rating
* App category

Kategori dapat disesuaikan dengan positioning aplikasi saat publikasi.

---

# 59. Privacy

Karena aplikasi memproses percakapan personal, privacy harus menjadi prioritas.

User harus dapat:

* menghapus akun
* menghapus chat
* menghapus memory
* mematikan memory
* menghapus mood
* menghapus reflection

Data deletion harus benar-benar diproses backend.

---

# 60. Data Retention

Jangan menyimpan data tanpa tujuan.

Contoh:

```text
Chat:
stored until user deletes

Memory:
stored until user deletes

Mood:
stored until user deletes

Reflection:
stored until user deletes
```

Kebijakan retention final harus mengikuti kebutuhan produk dan kebijakan privasi yang diterbitkan.

---

# 61. User Flow

```text
Install
   ↓
Splash
   ↓
Welcome
   ↓
Nickname
   ↓
Purpose
   ↓
Personality
   ↓
Privacy
   ↓
Home
   ↓
Choose Mode
   ↓
Chat
   ↓
AI Response
   ↓
Mood / Reflection
   ↓
History
```

---

# 62. Main User Journey

Contoh:

```text
User membuka aplikasi

↓

"Hi Surya 👋"

↓

"Gimana kabarmu hari ini?"

↓

User memilih:
😔

↓

"Kayaknya lagi nggak terlalu baik."

↓

User:
"Kerjaan gue bikin pusing."

↓

TEMENIN:
"Waduh 😕
Mau cerita bagian mana yang paling bikin
pusing?"

↓

User bercerita

↓

TEMENIN mendengarkan

↓

Conversation selesai

↓

"Thanks udah cerita."

↓

Mood dicatat
```

---

# 63. MVP Scope

## WAJIB

### Authentication

* Register
* Login
* Logout

### Onboarding

* Nickname
* Purpose
* Personality
* Privacy

### Home

* Greeting
* Quick actions
* Continue chat

### Chat

* Create chat
* Send message
* AI response
* Streaming
* History

### Personality

* Pendengar
* Santai
* Bijak
* Motivator

### Memory

* Enable/disable
* Basic memory
* Delete memory

### Mood

* Daily mood
* Mood history
* Basic summary

### Reflection

* Daily question
* Save answer
* History

### Profile

* Settings
* Theme
* Notification
* Privacy
* Delete account

### Safety

* AI safety prompt
* Crisis detection
* Safe response

---

# 64. Phase 2

Setelah MVP stabil:

* Voice chat
* Voice input
* Voice output
* AI-generated summaries
* Advanced memory
* Custom personality
* AI avatar
* More themes
* More companion characters
* Premium subscription
* Advanced mood analytics

---

# 65. Phase 3

Potential future features:

### AI Voice Companion

User dapat berbicara langsung.

### Companion Avatar

Animated character.

### Personalized AI

Personality semakin menyesuaikan preferensi user.

### Journaling AI

AI membantu membuat journal summary.

### Weekly Reflection

AI memberikan ringkasan refleksi mingguan tanpa diagnosis.

---

# 66. Recommended Development Order

## Sprint 1 — Foundation

```text
React Native
Expo
TypeScript
Expo Router
Theme
Navigation
Components
```

## Sprint 2 — Authentication

```text
Register
Login
JWT
Profile
SecureStore
```

## Sprint 3 — Chat

```text
Conversation
Message
AI API
Streaming
History
```

## Sprint 4 — Companion

```text
Personality
Modes
Memory
```

## Sprint 5 — Wellness Features

```text
Mood
Reflection
Dashboard
```

## Sprint 6 — Safety

```text
Safety prompt
Risk detection
Crisis flow
```

## Sprint 7 — Polish

```text
Animation
Dark mode
Loading
Error states
Empty states
```

## Sprint 8 — Production

```text
Testing
Security
Analytics
Privacy
Play Store assets
AAB
Closed testing
Production release
```

---

# 67. Recommended Tech Stack

## Mobile

```text
React Native
Expo
TypeScript
Expo Router
Zustand
TanStack Query
React Native Reanimated
Expo SecureStore
Expo Notifications
```

## Backend

```text
Go
Gin
JWT
PostgreSQL
Redis (optional)
SSE
```

## AI

```text
LLM API
```

AI provider dibuat abstraction layer sehingga provider dapat diganti tanpa mengubah mobile app.

Contoh:

```text
AIService
   │
   ├── Provider A
   ├── Provider B
   └── Provider C
```

---

# 68. Backend AI Interface

Contoh konsep:

```go
type AIProvider interface {
    GenerateResponse(
        ctx context.Context,
        request AIRequest,
    ) (<-chan string, error)
}
```

Dengan demikian AI provider tidak mengikat seluruh backend.

---

# 69. Environment Variables

Backend:

```env
APP_ENV=production

PORT=8080

DATABASE_URL=

JWT_SECRET=

AI_API_KEY=

REDIS_URL=

ALLOWED_ORIGINS=
```

Jangan pernah memasukkan:

```env
AI_API_KEY
JWT_SECRET
DATABASE_PASSWORD
```

ke React Native app.

---

# 70. Deployment

## Mobile

Development:

```text
Expo
```

Build:

```text
EAS Build
```

Production:

```text
Android App Bundle (.aab)
```

Upload ke Google Play Console.

---

## Backend

Contoh deployment:

```text
Internet
   ↓
Reverse Proxy
   ↓
Go API
   ↓
PostgreSQL
```

Bisa menggunakan:

* VPS
* cloud server
* managed PostgreSQL

---

# 71. Testing

## Frontend

Test:

* navigation
* login
* chat
* mood
* memory
* settings
* dark mode

## Backend

Test:

* authentication
* authorization
* chat
* database
* AI service
* rate limit

## Security

Test:

* unauthorized API access
* expired token
* invalid JWT
* SQL injection
* excessive requests
* account deletion
* data isolation

---

# 72. Acceptance Criteria

MVP dianggap selesai apabila:

### Authentication

* User dapat register
* User dapat login
* User dapat logout
* Token aman

### Chat

* User dapat membuat conversation
* User dapat mengirim message
* AI memberikan response
* Response dapat ditampilkan secara streaming
* History tersimpan

### Personality

* User dapat memilih personality
* Personality memengaruhi gaya response

### Memory

* Memory dapat aktif/nonaktif
* Memory dapat dihapus

### Mood

* User dapat mencatat mood
* History mood dapat dilihat

### Reflection

* User dapat menjawab daily reflection
* Jawaban tersimpan

### Privacy

* User dapat menghapus account
* User dapat menghapus data

### Safety

* AI tidak memberikan harmful instructions
* Crisis flow tersedia

### Android

* APK/AAB dapat di-build
* Tidak crash pada device target
* Back navigation bekerja
* Keyboard chat bekerja dengan baik
* Dark mode bekerja

---

# 73. Design Principle

TEMENIN harus terasa seperti:

**Teman ngobrol**

Bukan:

**Customer Service**

Bukan:

**Chatbot perusahaan**

Bukan:

**Aplikasi terapi**

Contoh tone:

❌

> "Berdasarkan informasi yang Anda berikan, saya menyarankan Anda melakukan..."

✅

> "Kalau menurut gue, ada beberapa kemungkinan sih. Tapi sebelum itu, gue pengen tahu dulu..."

---

# 74. Brand Personality

TEMENIN:

```text
Warm
Friendly
Casual
Non-judgmental
Calm
Respectful
Modern
```

Avoid:

```text
Too formal
Too robotic
Too motivational
Too clingy
Too emotional
Manipulative
```

---

# 75. Product Success Metrics

Metrik utama:

### Activation

Persentase user yang menyelesaikan onboarding.

### First Conversation

Persentase user yang mengirim pesan pertama.

### D1 Retention

User kembali setelah 1 hari.

### D7 Retention

User kembali setelah 7 hari.

### Conversations/User

Jumlah conversation per user.

### Messages/Conversation

Kedalaman percakapan.

### Mood Engagement

Berapa banyak user menggunakan mood tracking.

### Reflection Engagement

Berapa banyak user menggunakan reflection.

---

# 76. Important Product Principle

Jangan mengejar:

> "User harus terus menggunakan TEMENIN."

Tetapi:

> "User mendapatkan pengalaman yang berguna setiap kali menggunakan TEMENIN."

Produk tidak boleh sengaja membuat user merasa bersalah karena tidak membuka aplikasi.

---

# 77. Final Product Structure

Secara keseluruhan:

```text
                 TEMENIN
                    │
        ┌───────────┼───────────┐
        │           │           │
       CHAT        MOOD       REFLECTION
        │           │           │
        │           │           │
    AI Companion  Tracking    Journaling
        │
 ┌──────┼─────────────┐
 │      │             │
Curhat Santai     Perspektif
 │
Night Talk
 │
Personality
 │
Memory
```

---

# 78. Final MVP Architecture

```text
                 ANDROID
                    │
          React Native + Expo
                    │
          ┌─────────┴─────────┐
          │                   │
       Zustand          TanStack Query
          │                   │
          └─────────┬─────────┘
                    │
                  HTTPS
                    │
                    ▼
              ┌───────────┐
              │ Go / Gin  │
              └─────┬─────┘
                    │
        ┌───────────┼────────────┐
        │           │            │
        ▼           ▼            ▼
 PostgreSQL       Redis       AI API
```

---

# 79. Final Recommendation

Untuk versi pertama, **jangan membuat terlalu banyak fitur**.

Fokus utama:

```text
ONBOARDING
     ↓
HOME
     ↓
CHAT
     ↓
AI RESPONSE
     ↓
PERSONALITY
     ↓
MEMORY
     ↓
MOOD
     ↓
REFLECTION
```

Yang paling penting bukan jumlah fitur, tetapi kualitas percakapannya.

Kalau chat AI-nya terasa natural, UI-nya nyaman, response cepat, dan memory bekerja dengan baik, TEMENIN sudah memiliki fondasi yang kuat untuk menjadi aplikasi Play Store yang benar-benar usable.

---

# 80. MVP Definition

### TEMENIN V1.0

> **Aplikasi Android AI companion yang memungkinkan pengguna ngobrol, curhat, mendapatkan perspektif, mencatat mood, melakukan refleksi, dan memiliki percakapan yang terasa personal melalui personality serta memory yang dapat dikontrol pengguna.**

**Frontend:** React Native + Expo + TypeScript
**Backend:** Go + Gin
**Database:** PostgreSQL
**AI:** LLM API
**Authentication:** JWT
**Streaming:** SSE
**Build:** Expo EAS
**Platform:** Android / Google Play Store

**Prioritas utama:**

1. UX chat
2. AI response quality
3. Safety
4. Privacy
5. Performance
6. Personalization
7. Visual quality
