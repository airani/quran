package quran

import (
	"reflect"
	"testing"
)

func TestQuran_Surah(t *testing.T) {
	q, _ := NewQuran()
	type fields struct {
		Surahs []Surah
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantS  Surah
	}{
		{
			name: "Get first Surah",
			fields: fields{
				Surahs: q.Surahs,
			},
			args: args{
				n: 1,
			},
			wantS: q.Surahs[0],
		},
		{
			name: "Get Surah number 0",
			fields: fields{
				Surahs: q.Surahs,
			},
			args: args{
				n: 0,
			},
			wantS: Surah{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Quran{
				Surahs: tt.fields.Surahs,
			}
			if gotS := q.Surah(tt.args.n); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("Quran.Surah() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestQuranTranslate_Surah(t *testing.T) {
	q, _ := NewQuranTranslate(datasetTranslateFaFooladvand)
	type fields struct {
		Surahs []Surah
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantS  Surah
	}{
		{
			name: "Get first Surah",
			fields: fields{
				Surahs: q.Surahs,
			},
			args: args{
				n: 1,
			},
			wantS: q.Surahs[0],
		},
		{
			name: "Get Surah number 0",
			fields: fields{
				Surahs: q.Surahs,
			},
			args: args{
				n: 0,
			},
			wantS: Surah{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Quran{
				Surahs: tt.fields.Surahs,
			}
			if gotS := q.Surah(tt.args.n); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("Quran.Surah() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
