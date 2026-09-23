package ai

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type LocalSimulatorProvider struct{}

func NewLocalSimulatorProvider() *LocalSimulatorProvider {
	return &LocalSimulatorProvider{}
}

func (s *LocalSimulatorProvider) GenerateStream(ctx context.Context, req AIRequest) (<-chan string, <-chan error) {
	out := make(chan string)
	errCh := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errCh)

		// Determine last user message
		var lastUserMsg string
		for i := len(req.Messages) - 1; i >= 0; i-- {
			if req.Messages[i].Role == "user" {
				lastUserMsg = req.Messages[i].Content
				break
			}
		}

		response := generateEmpatheticResponse(req.Personality, req.Mode, req.UserNickname, lastUserMsg, req.Memories)

		// Simulate natural thinking/typing pause before first token
		time.Sleep(1200 * time.Millisecond)

		// Simulate token streaming with natural cadence
		words := strings.Fields(response)
		for i, word := range words {
			select {
			case <-ctx.Done():
				errCh <- ctx.Err()
				return
			default:
				token := word
				if i < len(words)-1 {
					token += " "
				}
				out <- token
				time.Sleep(35 * time.Millisecond)
			}
		}
	}()

	return out, errCh
}

func generateEmpatheticResponse(personality, mode, nickname, message string, memories []string) string {
	name := nickname
	if name == "" {
		name = "kamu"
	}
	lowerMsg := strings.ToLower(message)

	var memoryNote string
	if len(memories) > 0 {
		memoryNote = fmt.Sprintf(" (Gue juga masih inget lho, %s)", memories[0])
	}

	// Mode-based responses
	switch mode {
	case "night":
		switch personality {
		case "bijak":
			return fmt.Sprintf("Malam yang hening gini emang sering bikin pikiran kita mengembara ya, %s. Hari ini udah kamu lewatin dengan sebaik-baiknya. Hal apa yang paling menyita pikiranmu sebelum tidur ini?", name)
		case "santai":
			return fmt.Sprintf("Hei %s, masih melek aja jam segini hehe. Kenapa nih, ada yang bikin susah merem atau cuma pengen ngobrol santai aja? Sini cerita santai.", name)
		default:
			return fmt.Sprintf("Malam %s... Hari yang panjang ya? Kalau masih ada uneg-uneg atau hal yang ngeganjel di hati sebelum tidur, keluarin aja di sini. Gue bakal dengerin dengan tenang.", name)
		}

	case "perspective":
		if strings.Contains(lowerMsg, "bingung") || strings.Contains(lowerMsg, "pilihan") || strings.Contains(lowerMsg, "harus") {
			return fmt.Sprintf("Wajar banget kalau %s ngerasa dilema di posisi ini. Kalau kita liat dari satu sisi, keputusan ini punya nilai positifnya. Tapi di sisi lain, konsekuensinya juga perlu kamu antisipasi. Kira-kira hal apa yang paling kamu takuti kalau ngambil keputusan itu?", name)
		}
		return fmt.Sprintf("Menarik sudut pandangmu, %s. Coba kita bedah pelan-pelan ya. Terkadang apa yang kelihatan buntu itu cuma butuh kita liat dari sudut yang agak berbeda. Menurutmu opsi apa yang paling bikin hatimu lega?", name)

	case "casual":
		if strings.Contains(lowerMsg, "halo") || strings.Contains(lowerMsg, "hai") {
			return fmt.Sprintf("Halo %s! Lagi santai atau lagi ngerjain apa nih hari ini? Mau ngobrolin topik apa yang seru?", name)
		}
		if strings.Contains(lowerMsg, "game") || strings.Contains(lowerMsg, "main") {
			return fmt.Sprintf("Wah seru tuh! Kalau lagi jenuh emang paling pas rehat sejenak sambil main game atau refreshing. %s lagi sering main apa belakangan ini?", name)
		}
		return fmt.Sprintf("Haha iya juga ya, %s! Seru denger cerita kamu. Lanjutin dong, abis itu gimana ceritanya?", name)

	default: // "curhat" / "vent"
		if strings.Contains(lowerMsg, "capek") || strings.Contains(lowerMsg, "lelah") || strings.Contains(lowerMsg, "berat") {
			if personality == "motivator" {
				return fmt.Sprintf("Wajar banget kamu ngerasa capek hari ini, %s. Tarik napas dulu pelan-pelan ya. Kamu udah berjuang hebat sejauh ini! Ceritain ke gue bagian mana yang paling nguras energimu, kita urai bareng-bareng.%s", name, memoryNote)
			}
			if personality == "bijak" {
				return fmt.Sprintf("Rasa capek itu sinyal wajar dari tubuh dan pikiran kita, %s. Jangan dipaksa untuk selalu kuat setiap saat. Kalau boleh tahu, apa yang paling memberatkan perasaanmu saat ini?", name)
			}
			return fmt.Sprintf("Waduh, kedengerannya hari ini beneran berat banget buat kamu ya 🥺. Istirahat sejenak, jangan terlalu membebani diri sendiri. Mau cerita bagian mana yang paling bikin kamu capek, %s? Gue di sini siap dengerin.", name)
		}

		if strings.Contains(lowerMsg, "sedih") || strings.Contains(lowerMsg, "kecewa") || strings.Contains(lowerMsg, "sakit") {
			return fmt.Sprintf("Gue ikut ngerasain apa yang kamu rasain, %s. Perasaan kecewa dan sedih itu valid banget kok. Gak apa-apa kalau mau nangis atau meluapkan semuanya di sini. Gue ada buat kamu.", name)
		}

		return fmt.Sprintf("Gue dengerin baik-baik apa yang kamu ceritain barusan, %s. Kadang ngeluarin apa yang ada di kepala itu langkah pertama yang paling melegakan. Terus, gimana kelanjutannya?", name)
	}
}
