package quran

import (
	"math/rand"
	"time"
)

// Surah struct of a surah of Quran
type Surah struct {
	Index int    `xml:"index,attr"`
	Name  string `xml:"name,attr"`
	Ayahs []Ayah `xml:"aya"`
}

// Ayah Returns a Ayah by number
func (s Surah) Ayah(n int) (a Ayah) {
	if n > len(s.Ayahs) || n == 0 {
		return
	}
	return s.Ayahs[n-1]
}

// RandAyah Returns a Ayah by random
func (s Surah) RandAyah() Ayah {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(s.Ayahs)) + 1
	return s.Ayah(n)
}
