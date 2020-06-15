package quran

const (
	// SajdaRecommended Sajda for ayah recommended
	SajdaRecommended = "recommended"
	// SajdaObligatory Sajda for ayah obligatory
	SajdaObligatory = "obligatory"
)

// Ayah returns Ayah struct of Surah
type Ayah struct {
	Index     int    `xml:"index,attr"`
	Text      string `xml:"text,attr"`
	Bismillah string `xml:"bismillah,attr"`
	Sajda     string `xml:"sajda,attr,omitempty"`
}

// IsSajdaObligatory returns Ayah is sajda obligatory
func (a Ayah) IsSajdaObligatory() bool {
	if a.Sajda == SajdaObligatory {
		return true
	}
	return false
}

func (a Ayah) HasBismillah() bool {
	if a.Bismillah == "" {
		return false
	}
	return true
}
