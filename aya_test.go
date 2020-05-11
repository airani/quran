package quran

import "testing"

func TestAyah_IsSajdaObligatory(t *testing.T) {
	q, _ := New()
	s := q.Surah(32)
	tests := []struct {
		name string
		ayah Ayah
		want bool
	}{
		{
			name: "Ayah is obligatory",
			ayah: s.Ayah(15),
			want: true,
		},
		{
			name: "Ayah is not obligatory",
			ayah: s.Ayah(14),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Ayah{
				Index:     tt.ayah.Index,
				Text:      tt.ayah.Text,
				Bismillah: tt.ayah.Bismillah,
				Sajda:     tt.ayah.Sajda,
			}
			if got := a.IsSajdaObligatory(); got != tt.want {
				t.Errorf("Ayah.IsSajdaObligatory() = %v, want %v", got, tt.want)
			}
		})
	}
}
