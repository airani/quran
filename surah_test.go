package quran

import (
	"reflect"
	"testing"
)

func TestSurah_Ayah(t *testing.T) {
	q, _ := New()
	s := q.Surah(1)
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields Surah
		args   args
		wantA  Ayah
	}{
		{
			fields: s,
			args: args{
				n: 1,
			},
			wantA: s.Ayahs[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Surah{
				Index: tt.fields.Index,
				Name:  tt.fields.Name,
				Ayahs: tt.fields.Ayahs,
			}
			if gotA := s.Ayah(tt.args.n); !reflect.DeepEqual(gotA, tt.wantA) {
				t.Errorf("Surah.Ayah() = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}
