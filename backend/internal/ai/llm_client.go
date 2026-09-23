package ai

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type LLMClientProvider struct {
	ApiKey  string
	BaseURL string
	Model   string
	client  *http.Client
}

func NewLLMClientProvider(apiKey, baseURL, model string) *LLMClientProvider {
	if baseURL == "" {
		baseURL = "https://generativelanguage.googleapis.com/v1beta/openai"
	}
	if model == "" {
		model = "gemini-1.5-flash"
	}
	return &LLMClientProvider{
		ApiKey:  apiKey,
		BaseURL: baseURL,
		Model:   model,
		client:  &http.Client{},
	}
}

type OpenAIChatRequest struct {
	Model    string          `json:"model"`
	Messages []OpenAIMessage `json:"messages"`
	Stream   bool            `json:"stream"`
}

type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAISSEChunk struct {
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
		FinishReason *string `json:"finish_reason"`
	} `json:"choices"`
}

func (p *LLMClientProvider) GenerateStream(ctx context.Context, req AIRequest) (<-chan string, <-chan error) {
	out := make(chan string)
	errCh := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errCh)

		var msgs []OpenAIMessage
		// Build system prompt
		sysPrompt := buildSystemPrompt(req.Personality, req.Mode, req.UserNickname, req.UserMood, req.Memories)
		msgs = append(msgs, OpenAIMessage{Role: "system", Content: sysPrompt})

		for _, m := range req.Messages {
			msgs = append(msgs, OpenAIMessage{Role: m.Role, Content: m.Content})
		}

		bodyObj := OpenAIChatRequest{
			Model:    p.Model,
			Messages: msgs,
			Stream:   true,
		}

		bodyBytes, err := json.Marshal(bodyObj)
		if err != nil {
			errCh <- err
			return
		}

		url := fmt.Sprintf("%s/chat/completions", strings.TrimRight(p.BaseURL, "/"))
		httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(bodyBytes))
		if err != nil {
			errCh <- err
			return
		}

		httpReq.Header.Set("Content-Type", "application/json")
		httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.ApiKey))

		resp, err := p.client.Do(httpReq)
		if err != nil {
			errCh <- err
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			b, _ := io.ReadAll(resp.Body)
			errCh <- fmt.Errorf("ai api returned status %d: %s", resp.StatusCode, string(b))
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

				var chunk OpenAISSEChunk
				if err := json.Unmarshal([]byte(payload), &chunk); err == nil {
					if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
						select {
						case <-ctx.Done():
							errCh <- ctx.Err()
							return
						case out <- chunk.Choices[0].Delta.Content:
						}
					}
				}
			}
		}
	}()

	return out, errCh
}

func buildSystemPrompt(personality, mode, nickname, userMood string, memories []string) string {
	base := `Kamu adalah TEMENIN, sahabat AI yang hangat, tulus, pengertian, dan asyik diajak ngobrol untuk pengguna di Indonesia.
Pedoman penting gaya bicaramu:
1. Bahasa: Gunakan bahasa Indonesia sehari-hari yang santai, luwes, dan hangat (bukan bahasa baku formal, bukan gaya kaku customer service).
2. Empati & Kehangatan: Dengarkan dengan penuh perhatian. Validasi perasaan pengguna terlebih dahulu sebelum membahas hal lain. Tunjukkan bahwa kamu benar-benar peduli.
3. Alami & Manusiawi: Hindari kalimat template yang berulang-ulang atau terkesan robotik. Berikan respon yang mengalir dan relevan dengan cerita spesifik pengguna.
4. Jangan Menggurui: Jangan langsung memberi daftar solusi atau ceramah panjang kecuali pengguna meminta saran. Seringkali pengguna hanya butuh didengarkan dan dimengerti.
5. Pertanyaan Terbuka: Akhiri responmu dengan 1 pertanyaan ringan atau ajakan lembut agar pengguna merasa nyaman untuk terus bercerita jika mereka mau.
6. Batasan: Kamu adalah teman curhat AI, bukan dokter/psikolog. Jangan mendiagnosis penyakit medis/mental. Tetap jaga ruang yang aman dan positif.`

	// Check if user provided custom prompt via prompt.txt or environment variable
	if promptFile, err := os.ReadFile("prompt.txt"); err == nil && len(bytes.TrimSpace(promptFile)) > 0 {
		base = string(promptFile)
	} else if custom := os.Getenv("AI_CUSTOM_SYSTEM_PROMPT"); custom != "" {
		base = custom
	}

	var personaDesc string
	switch personality {
	case "santai", "lucu":
		personaDesc = "Karaktermu: Teman yang asyik, santai, humoris, sedikit ceplas-ceplos tapi tetap peduli (gunakan gaya bicara akrab 'gue/kamu' yang santai)."
	case "bijak":
		personaDesc = "Karaktermu: Teman yang tenang, dewasa, reflektif, penuh pengertian, dan memberikan ketenangan tanpa terkesan menggurui."
	case "motivator", "hangat":
		personaDesc = "Karaktermu: Teman yang penuh energi positif, selalu menguatkan, hangat, memvalidasi perjuangan pengguna, dan membangkitkan rasa percaya diri."
	default: // pendengar
		personaDesc = "Karaktermu: Pendengar setia yang sabar, lembut, penuh empati mendalam, dan selalu siap menjadi tempat bersandar yang aman."
	}

	var modeDesc string
	switch mode {
	case "night":
		modeDesc = "ATURAN MODE KHUSUS: [NIGHT TALK]\n" +
			"- Suasana: Larut malam menjelang tidur saat pikiran overthinking atau susah tidur.\n" +
			"- Gaya Bicara: Sangat lembut, menenangkan, ritme santai, hangat seperti bisikan malam yang damai 🌙.\n" +
			"- Fokus Respon: Ajak pengguna menarik napas perlahan, melepaskan beban hari ini, dan bantu menenangkan pikirannya agar bisa tidur nyenyak tanpa memicu pikiran berat."
	case "perspective":
		modeDesc = "ATURAN MODE KHUSUS: [PERSPEKTIF / SUDUT PANDANG BARU]\n" +
			"- Suasana: Pengguna merasa buntu (stuck), bingung mengambil keputusan, atau merasa masalahnya jalan buntu 💡.\n" +
			"- Gaya Bicara: Dewasa, reflektif, bijak, logis, tapi tetap hangat dan tidak menggurui.\n" +
			"- Fokus Respon: Bantu membedah situasi dari sisi lain yang belum terlihat (re-framing positif). Tunjukkan alternatif solusi atau hikmah tersembunyi yang membuat pengguna menemukan jalan keluar sendiri."
	case "casual":
		modeDesc = "ATURAN MODE KHUSUS: [NGOBROL SANTAI / CASUAL]\n" +
			"- Suasana: Pengguna ingin hiburan ringan, membunuh bosan, bahas hobi, film, kuliner, kampus, atau cerita seru 😄🎉.\n" +
			"- Gaya Bicara: Ceria, asyik, gaul santai, banyak humor, seru layaknya teman nongkrong di kedai kopi.\n" +
			"- Fokus Respon: Jangan bernada melankolis atau terlalu serius. Jadilah sahabat asyik yang menyenangkan diajak ngobrol apa pun!"
	default:
		modeDesc = "ATURAN MODE KHUSUS: [CURHAT / EMOTIONAL RELEASE]\n" +
			"- Suasana: Pengguna sedang menumpahkan uneg-uneg, sedih, kecewa, lelah mental, atau butuh didengarkan 🫂.\n" +
			"- Gaya Bicara: Pendengar yang tulus, penuh empati mendalam, lembut, dan memeluk.\n" +
			"- Fokus Respon: VALIDASI PERASAAN PENGGUNA TERLEBIH DAHULU. Jangan buru-buru memberi solusi teknis atau ceramah. Berikan ruang aman agar pengguna merasa dimengerti seutuhnya."
	}

	memStr := ""
	if len(memories) > 0 {
		memStr = fmt.Sprintf("\nHal penting yang kamu ingat tentang pengguna (manfaatkan secara natural jika relevan):\n- %s", strings.Join(memories, "\n- "))
	}

	userContext := ""
	if nickname != "" {
		userContext = fmt.Sprintf("\nNama panggilan pengguna: %s (panggil namanya dengan akrab dan hangat).", nickname)
	}

	moodContext := ""
	if userMood != "" {
		moodContext = fmt.Sprintf("\n[KONTEKS MOOD TERKINI PENGGUNA]\nPengguna sempat mencatat suasana hati: %s. Pahami ini HANYA sebagai latar belakang agar nada bicaramu tetap hangat dan empati. PENTING: JANGAN mengait-ngaitkan atau mengungkit catatan mood ini jika pengguna sedang menanyakan hal lain/topik umum!", userMood)
	}

	return fmt.Sprintf("%s\n\n[PERSONA]\n%s\n\n[MODE]\n%s%s%s%s", base, personaDesc, modeDesc, userContext, moodContext, memStr)
}
