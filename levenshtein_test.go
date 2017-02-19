// Copyright 2015 Benjamin BALET. All rights reserved.
// Use of this source code is governed by the BSD license
// license that can be found in the LICENSE file.

package stopwords

import (
	"testing"
)

//French language
func TestLevenshteinFrench(t *testing.T) {
	string1 := []byte("la fin d'un bel après-midi d'été")
	string2 := []byte("cet été, nous avons eu un bel après-midi")
	string3 := []byte("l'hiver, la chaleur n'était pas là à midi")

	if LevenshteinDistance(string2, string3, "fr", false) < LevenshteinDistance(string2, string1, "fr", false) {
		t.Errorf("Comparison failed")
	}

	//Levenshtein distance should be the same
	if LevenshteinDistance(string1, string2, "fr", false) != LevenshteinDistance(string2, string1, "fr", false) {
		t.Errorf("Levenshtein distance should be the same between 1/2 and 2/1")
	}
}

//English language
func TestLevenshteinEnglish(t *testing.T) {
	string1 := []byte("the war is coming, send more troop. This soldier was courageous.")
	string2 := []byte("the troop was composed of veteran soldiers of the last war.")
	string3 := []byte("the is an article of the english language it was created long before the last reform.")

	if LevenshteinDistance(string2, string3, "en", false) < LevenshteinDistance(string2, string1, "en", false) {
		t.Errorf("Comparison failed")
	}

	//Levenshtein distance should be the same
	if LevenshteinDistance(string1, string2, "fr", false) != LevenshteinDistance(string2, string1, "fr", false) {
		t.Errorf("Levenshtein distance should be the same between 1/2 and 2/1")
	}
}
