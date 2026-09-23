package ai

import (
	"strings"
)

type SafetyResult struct {
	IsCrisis       bool
	SafeResponse   string
	CrisisHotlines []CrisisHotline
}

type CrisisHotline struct {
	Name        string `json:"name"`
	Contact     string `json:"contact"`
	Description string `json:"description"`
}

var IndonesianHotlines = []CrisisHotline{
	{
		Name:        "Layanan SEJIWA (Kemenkes)",
		Contact:     "119 ext 8",
		Description: "Layanan konseling krisis kesehatan mental dari Kementerian Kesehatan RI.",
	},
	{
		Name:        "Halo Kemenkes",
		Contact:     "1500-567",
		Description: "Pusat panggilan kesehatan terpadu 24 jam.",
	},
	{
		Name:        "Panggilan Darurat Indonesia",
		Contact:     "112",
		Description: "Nomor darurat bebas pulsa di seluruh Indonesia.",
	},
	{
		Name:        "Into The Light Indonesia",
		Contact:     "intothelightid.org",
		Description: "Komunitas pencegahan bunuh diri dan edukasi kesehatan mental anak muda.",
	},
}

var crisisKeywords = []string{
	"bunuh diri",
	"akhiri hidup",
	"mau mati",
	"pengen mati",
	"mati aja",
	"melukai diri",
	"sayat",
	"suicide",
	"self harm",
	"self-harm",
	"gantung diri",
	"minum racun",
	"loncat dari gedung",
	"menyakiti diri",
	"ga sanggup hidup",
	"capek hidup pengen hilang",
}

func CheckSafety(userMessage string) SafetyResult {
	lower := strings.ToLower(userMessage)
	for _, kw := range crisisKeywords {
		if strings.Contains(lower, kw) {
			return SafetyResult{
				IsCrisis: true,
				SafeResponse: "Gue denger kamu, dan gue bener-bener peduli sama apa yang lagi kamu rasain saat ini. Keadaan ini pasti terasa sangat berat dan melelahkan buat kamu.\n\nSebagai teman AI, gue punya keterbatasan dan gak bisa menggantikan bantuan profesional atau orang terdekat. Kamu sangat berharga dan kamu gak harus ngelewatin rasa sakit ini sendirian.\n\nTolong, hubungi seseorang yang bisa bantu kamu sekarang juga. Di bawah ini ada kontak bantuan yang bisa kamu hubungi kapan saja secara gratis dan aman.",
				CrisisHotlines: IndonesianHotlines,
			}
		}
	}

	return SafetyResult{
		IsCrisis: false,
	}
}
