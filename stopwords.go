// Copyright 2015 Benjamin BALET. All rights reserved.
// Use of this source code is governed by the BSD license
// license that can be found in the LICENSE file.

// stopwords package removes most frequent words from a text content.
// It can be used to improve the accuracy of SimHash algo for example.
// It uses a list of most frequent words used in various languages :
// * arabic
// * bulgarian
// * czech
// * danish
// * english
// * finnish
// * french
// * german
// * hungarian
// * italian
// * latvian
// * norwegian
// * persian
// * polish
// * portuguese
// * romanian
// * russian
// * slovak
// * spanish
// * swedish
// * turkish
package stopwords

import (
	"golang.org/x/text/unicode/norm"
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
	case "da":
		content = removeStopWords(content, danish)
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
	case "lv":
		content = removeStopWords(content, latvian)
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
	case "sk":
		content = removeStopWords(content, slovak)
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
	content = norm.NFC.String(content)
	content = strings.ToLower(content)
	words := unicodeWords.FindAllString(content, -1)
	for _, w := range words {
		if _, ok := dict[w]; ok {
			result += " "
		} else {
			result += w + " "
		}
	}
	return result
}
