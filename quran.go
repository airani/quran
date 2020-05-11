package quran

import (
	"encoding/xml"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// Quran struct of xml file
type Quran struct {
	Surahs []Surah `xml:"sura"`
}

// New Quran
func New() (Quran, error) {
	return NewQuranByXmlFile(datasetPath + string(datasetQuranSimple))
}

// NewTranslate Translate of Quran
func NewTranslate(file datasetFile) (Quran, error) {
	return NewQuranByXmlFile(datasetPath + string(file))
}

// NewQuranByXmlFile read xml file of Quran and returns a Quran struct
func NewQuranByXmlFile(file string) (q Quran, err error) {
	xmlFile, err := os.Open(file)
	if err != nil {
		return
	}

	defer func() {
		if ferr := xmlFile.Close(); ferr != nil {
			err = ferr
		}
	}()

	b, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return
	}

	err = xml.Unmarshal(b, &q)
	if err != nil {
		return
	}

	return
}

// Surah Returns a Surah by number
func (q Quran) Surah(n int) Surah {
	if n > len(q.Surahs) || n == 0 {
		return Surah{}
	}
	return q.Surahs[n-1]
}

// RandSurah Returns a Surah by random
func (q Quran) RandSurah() Surah {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(q.Surahs)) + 1
	return q.Surah(n)
}
