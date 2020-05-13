package quran

import (
	"encoding/xml"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/airani/quran/dataset"
)

// Quran struct of xml file
type Quran struct {
	Surahs []Surah `xml:"sura"`
}

// NewSimple Quran
func NewSimple() (Quran, error) {
	return NewQuranByXml(dataset.Simple())
}

// NewQuranByXml read xml string of Quran and returns a Quran struct
func NewQuranByXml(xmlStr string) (q Quran, err error) {
	r := strings.NewReader(xmlStr)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	err = xml.Unmarshal(b, &q)
	if err != nil {
		return
	}

	return
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
