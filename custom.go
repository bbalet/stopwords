// Copyright 2015 Benjamin BALET. All rights reserved.
// Use of this source code is governed by the BSD license
// license that can be found in the LICENSE file.

// Package stopwords allows you to customize the list of stopwords
package stopwords

import (
	"io/ioutil"
	"strings"

	"golang.org/x/text/language"
)

// LoadStopWordsFromFile loads a list of stop words from a file
// filePath is the full path to the file to be loaded
// langCode is a BCP 47 or ISO 639-1 language code (e.g. "en" for English).
// sep is the string separator (e.g. "\n" for newline)
func LoadStopWordsFromFile(filePath string, langCode string, sep string) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	LoadStopWordsFromString(string(b), langCode, sep)
}

// LoadStopWordsFromString loads a list of stop words from a string
// filePath is the full path to the file to be loaded
// langCode is a BCP 47 or ISO 639-1 language code (e.g. "en" for English).
// sep is the string separator (e.g. "\n" for newline)
func LoadStopWordsFromString(wordsList string, langCode string, sep string) {

	//Parse language
	tag := language.Make(langCode)
	base, _ := tag.Base()
	langCode = base.String()

	words := strings.Split(strings.ToLower(wordsList), sep)
	stopWordList := make(map[string]string)
	for _, word := range words {
		stopWordList[word] = ""
	}

	switch langCode {
	case "ar":
		arabic = stopWordList
	case "bg":
		bulgarian = stopWordList
	case "cs":
		czech = stopWordList
	case "da":
		danish = stopWordList
	case "de":
		german = stopWordList
	case "el":
		greek = stopWordList
	case "en":
		english = stopWordList
	case "es":
		spanish = stopWordList
	case "fa":
		persian = stopWordList
	case "fr":
		french = stopWordList
	case "fi":
		finnish = stopWordList
	case "hu":
		hungarian = stopWordList
	case "id":
		indonesian = stopWordList
	case "it":
		italian = stopWordList
	case "ja":
		japanese = stopWordList
	case "km":
		khmer = stopWordList
	case "lv":
		latvian = stopWordList
	case "nl":
		dutch = stopWordList
	case "no":
		norwegian = stopWordList
	case "pl":
		polish = stopWordList
	case "pt":
		portuguese = stopWordList
	case "ro":
		romanian = stopWordList
	case "ru":
		russian = stopWordList
	case "sk":
		slovak = stopWordList
	case "sv":
		swedish = stopWordList
	case "th":
		thai = stopWordList
	case "tr":
		turkish = stopWordList
	}
}
