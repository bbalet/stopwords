// Copyright 2015 Benjamin BALET. All rights reserved.
// Use of this source code is governed by the BSD license
// license that can be found in the LICENSE file.

// stopwords package removes most frequent words from a text content
// It uses a list of most frequent words used in various languages :
// * arabic
// * bulgarian
// * czech
// * german
// * english
// * spanish
// * persian
// * french
// * finnish
// * hungarian
// * italian
// * norwegian
// * polish
// * portuguese
// * romanian
// * russian
// * swedish
// * turkish

package stopwords

import (
	"html"
	"regexp"
	"strings"
)

var (
	remTags      = regexp.MustCompile(`<[^>]*>`)
	oneSpace     = regexp.MustCompile(`\s{2,}`)
	unicodeWords = regexp.MustCompile(`[\pL-_']+`)
)

// CleanContent removes useless spaces and stop words from content.
// ISO 639-1 language code (if not supported, nothing is removed form source).
// If removeHTML is TRUE, remove HTML tags from content and unescape entities.
func CleanContent(content string, langCode string, removeHTML bool) string {

	//Remove HTML tags
	if removeHTML {
		content = remTags.ReplaceAllString(content, " ")
		content = html.UnescapeString(content)
	}

	//Remove stop words by using a list of most frequent words
	switch langCode {
	case "ar":
		content = removeStopWords(content, arabic)
	case "bg":
		content = removeStopWords(content, bulgarian)
	case "cs":
		content = removeStopWords(content, czech)
	case "de":
		content = removeStopWords(content, german)
	case "el":
		content = removeStopWords(content, greek)
	case "en":
		content = removeStopWords(content, english)
	case "es":
		content = removeStopWords(content, spanish)
	case "fa":
		content = removeStopWords(content, persian)
	case "fr":
		content = removeStopWords(content, french)
	case "fi":
		content = removeStopWords(content, finnish)
	case "hu":
		content = removeStopWords(content, hungarian)
	case "it":
		content = removeStopWords(content, italian)
	case "nl":
		content = removeStopWords(content, dutch)
	case "no":
		content = removeStopWords(content, norwegian)
	case "pl":
		content = removeStopWords(content, polish)
	case "pt":
		content = removeStopWords(content, portuguese)
	case "ro":
		content = removeStopWords(content, romanian)
	case "ru":
		content = removeStopWords(content, russian)
	case "sv":
		content = removeStopWords(content, swedish)
	case "tr":
		content = removeStopWords(content, turkish)
	}

	//Remove duplicated space characters
	content = oneSpace.ReplaceAllString(content, " ")

	return content
}

// removeStopWords Iterates through a list of words and removes stop words
func removeStopWords(content string, dict map[string]string) string {
	var result string
	words := unicodeWords.FindAllString(content, -1)
	for _, w := range words {
		if _, ok := dict[strings.ToLower(w)]; ok {
			result += " "
		} else {
			result += strings.ToLower(w) + " "
		}
	}
	return result
}
