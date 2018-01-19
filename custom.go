// Copyright 2015 Benjamin BALET. All rights reserved.
// Use of this source code is governed by the BSD license
// license that can be found in the LICENSE file.

// Package stopwords allows you to customize the list of stopwords
package stopwords

import (
	"bytes"
	"io/ioutil"
	"sort"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/unicode/norm"
)

// wordCounters is an Internal map allowing us to count the frequencies of words
// Unless it is reset, the counting is cumulative
var wordCounters map[string]int

// Pair is a data structure to hold a Word/Frequency pair.
type Pair struct {
  Word string
  Frequency int
}

// A slice of Pairs that implements sort.
type PairList []Pair
// Interface to sort by Value in DESCending order.
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Frequency > p[j].Frequency }

// init the map holding the word frequencies
func init() {
  wordCounters = make(map[string]int)
}

// ResetWordsCounter resets the map holding the word frequencies
// because the library accepts cumulative calls to functions calculating the
// word frequencies. So it is the responsability of the user to reset counters.
func ResetWordsCounter() {
	wordCounters = make(map[string]int)
}

// CountWordsFrequencyFromString counts the frequency of words in a string
// content is the string to analyze
func CountWordsFrequencyString(content string) {
	CountWordsFrequency([]byte(content))
}

// CountWordsFrequency counts the frequency of words in a bytes array (string)
// It relies on the same word segmenter used by the other stopwords' algos
// So it means that it can be overwritten and that n option allows you to strip
// digits or not.
// content is the string to analyze (as a a bytes array)
func CountWordsFrequency(content []byte) {
	content = norm.NFC.Bytes(content)
	content = bytes.ToLower(content)
	words := wordSegmenter.FindAll(content, -1)
	for _, w := range words {
		wordCounters[string(w)]++
	}
}

// GetWordsFrequencies returns the words list sorted by frequency in DESC order
func GetWordsFrequencies() PairList {
	p := make(PairList, len(wordCounters))
	i := 0
	for k, v := range wordCounters {
		 p[i] = Pair{k, v}
		 i++
	}
	sort.Sort(p)
	return p
}

// LoadStopWordsFromFreqTable erases the predefined list of stop words
// with the list of words frequency that we've calculated.
// langCode is a BCP 47 or ISO 639-1 language code (e.g. "en" for English).
// max allows to set a limit (the 'max' most frequent words)
func LoadStopWordsFromFreqTable(langCode string, max int) {
	var wordsList string
	frequencies := GetWordsFrequencies()
	if max > 0 {
		if max < len(frequencies) {
			frequencies = frequencies[:max]
		}
	}
	for k, _ := range frequencies {
		 wordsList += frequencies[k].Word + ","
	}
	LoadStopWordsFromString(wordsList, langCode, ",")
}

// LoadStopWordsFromFile loads a list of stop words from a file
// filePath is the full path to the file to be loaded
// langCode is a BCP 47 or ISO 639-1 language code (e.g. "en" for English).
// sep is the string separator (e.g. "\n" for newline)
func LoadStopWordsFromFile(filePath string, langCode string, sep string) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
			panic(err)
	}
	loadStopWordsFromString(string(b), langCode, sep, true)
}

// AppendStopWordsFromFile append a list of stop words from a file
// filePath is the full path to the file to be loaded
// langCode is a BCP 47 or ISO 639-1 language code (e.g. "en" for English).
// sep is the string separator (e.g. "\n" for newline)
func AppendStopWordsFromFile(filePath string, langCode string, sep string) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
			panic(err)
	}
	loadStopWordsFromString(string(b), langCode, sep, false)
}

// LoadStopWordsFromString loads a list of stop words from a string
// filePath is the full path to the file to be loaded
// langCode is a BCP 47 or ISO 639-1 language code (e.g. "en" for English).
// sep is the string separator (e.g. "\n" for newline)
func LoadStopWordsFromString(wordsList string, langCode string, sep string) {
	loadStopWordsFromString(wordsList, langCode, sep, true)
}

// AppendStopWordsFromString append the list of stop words with a string
// filePath is the full path to the file to be loaded
// langCode is a BCP 47 or ISO 639-1 language code (e.g. "en" for English).
// sep is the string separator (e.g. "\n" for newline)
func AppendStopWordsFromString(wordsList string, langCode string, sep string) {
	loadStopWordsFromString(wordsList, langCode, sep, false)
}

// loadStopWordsFromString loads a list of stop words from a string
// filePath is the full path to the file to be loaded
// langCode is a BCP 47 or ISO 639-1 language code (e.g. "en" for English).
// sep is the string separator (e.g. "\n" for newline)
// reset is a flag indicating if the map must be erased before we add the list
func loadStopWordsFromString(wordsList string, langCode string, sep string, reset bool) {

	//Parse language
	tag := language.Make(langCode)
	base, _ := tag.Base()
	langCode = base.String()

	words := strings.Split(strings.ToLower(wordsList), sep)

	switch langCode {
	case "ar":
		if reset {
			arabic = make(map[string]string)
		}
		for _, word := range words {
				arabic[word] = ""
		}
	case "bg":
		if reset {
			bulgarian = make(map[string]string)
		}
		for _, word := range words {
				bulgarian[word] = ""
		}
	case "cs":
		if reset {
			czech = make(map[string]string)
		}
		for _, word := range words {
				czech[word] = ""
		}
	case "da":
		if reset {
			danish = make(map[string]string)
		}
		for _, word := range words {
				danish[word] = ""
		}
	case "de":
		if reset {
			german = make(map[string]string)
		}
		for _, word := range words {
				german[word] = ""
		}
	case "el":
		if reset {
			greek = make(map[string]string)
		}
		for _, word := range words {
				greek[word] = ""
		}
	case "en":
		if reset {
			english = make(map[string]string)
		}
		for _, word := range words {
				english[word] = ""
		}
	case "es":
		if reset {
			spanish = make(map[string]string)
		}
		for _, word := range words {
				spanish[word] = ""
		}
	case "fa":
		if reset {
			persian = make(map[string]string)
		}
		for _, word := range words {
				persian[word] = ""
		}
	case "fr":
		if reset {
			french = make(map[string]string)
		}
		for _, word := range words {
				french[word] = ""
		}
	case "fi":
		if reset {
			finnish = make(map[string]string)
		}
		for _, word := range words {
				finnish[word] = ""
		}
	case "hu":
		if reset {
			hungarian = make(map[string]string)
		}
		for _, word := range words {
				hungarian[word] = ""
		}
	case "id":
		if reset {
			indonesian = make(map[string]string)
		}
		for _, word := range words {
				indonesian[word] = ""
		}
	case "it":
		if reset {
			italian = make(map[string]string)
		}
		for _, word := range words {
				italian[word] = ""
		}
	case "ja":
		if reset {
			japanese = make(map[string]string)
		}
		for _, word := range words {
				japanese[word] = ""
		}
	case "km":
		if reset {
			khmer = make(map[string]string)
		}
		for _, word := range words {
				khmer[word] = ""
		}
	case "lv":
		if reset {
			latvian = make(map[string]string)
		}
		for _, word := range words {
				latvian[word] = ""
		}
	case "nl":
		if reset {
			dutch = make(map[string]string)
		}
		for _, word := range words {
				dutch[word] = ""
		}
	case "no":
		if reset {
			norwegian = make(map[string]string)
		}
		for _, word := range words {
				norwegian[word] = ""
		}
	case "pl":
		if reset {
			polish = make(map[string]string)
		}
		for _, word := range words {
				polish[word] = ""
		}
	case "pt":
		if reset {
			portuguese = make(map[string]string)
		}
		for _, word := range words {
				portuguese[word] = ""
		}
	case "ro":
		if reset {
			romanian = make(map[string]string)
		}
		for _, word := range words {
				romanian[word] = ""
		}
	case "ru":
		if reset {
			russian = make(map[string]string)
		}
		for _, word := range words {
				russian[word] = ""
		}
	case "sk":
		if reset {
			slovak = make(map[string]string)
		}
		for _, word := range words {
				slovak[word] = ""
		}
	case "sv":
		if reset {
			swedish = make(map[string]string)
		}
		for _, word := range words {
				swedish[word] = ""
		}
	case "th":
		if reset {
			thai = make(map[string]string)
		}
		for _, word := range words {
				thai[word] = ""
		}
	case "tr":
		if reset {
			turkish = make(map[string]string)
		}
		for _, word := range words {
				turkish[word] = ""
		}
	}
}
