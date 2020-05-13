package quran

import (
	"errors"
	"github.com/airani/quran/dataset"
)

type Translate string

const (
	TranslateFaFooladvand Translate = "fa_fooladvand"
)

// NewTranslate Translate of Quran
func NewTranslate(translate Translate) (Quran, error) {
	switch translate {
	case TranslateFaFooladvand:
		return NewQuranByXml(dataset.TranslateFaFooladvand())
	}
	return Quran{}, errors.New("translate is not valid")
}
