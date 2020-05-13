# Quran Go Module

[![Build Status](https://travis-ci.com/airani/quran.svg?branch=master)](https://travis-ci.com/airani/quran)

Quran module for Go lang to working with Quran ayahs and surahs

### New Quran
```go
q, err := quran.New()
if err != nil {
    // ...
}
```

### Get Surah
```go
s := q.Surah(1) // Surah with number 1
```

### Get Random Surah
```go
s.RandSurah()
```

### Get Ayah
```go
s.Ayah(1) // Ayah number 1 of surah
```

### Get Random Ayah Of Surah
```go
s.RandAyah()
```
