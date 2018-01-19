// Copyright 2015 Benjamin BALET. All rights reserved.
// Use of this source code is governed by the BSD license
// license that can be found in the LICENSE file.

package stopwords

import (
	"testing"
)

func TestAppendToStopWords(t *testing.T) {
	source3 := "Après son enfance à Vinci, Léonard est élève auprès du célèbre peintre et sculpteur florentin Andrea del Verrocchio."
	additionalStopWords := "vinci léonard élève célèbre peintre sculpteur Andrea del Verrocchio"
	expected3 := " enfance auprès florentin "

	//Append additionnal words
	AppendStopWordsFromString(additionalStopWords, "fr", " ")
	actual3 := CleanString(source3, "fr", false)
	if actual3 != expected3 {
		t.Errorf("Test failed, got: '%s'", actual3)
	}
	//Restore the package state
	LoadStopWordsFromFile("_french.txt", "fr", "\r\n")
}

func TestCountWordsFrequency(t *testing.T) {
	source4 := "AAAA AAAA BBBB BB B BBB AAAA CCCC CC C BBBB BB A B CCCC DDDD AAAA"
	expected4 := " bbb cc c a dddd "

	CountWordsFrequencyString(source4)
	LoadStopWordsFromFreqTable("fr", 100)
	//It would remove all words..
	actual4 := CleanString(source4, "fr", false)
	if actual4 != " " {
		t.Errorf("Test failed, got: '%s'", actual4)
	}
	//It would remove some words
	LoadStopWordsFromFreqTable("fr", 5)
	actual5 := CleanString(source4, "fr", false)
	if actual5 != expected4 {
		t.Errorf("Test failed, got: '%s'", actual5)
	}
	//Restore the package state
	LoadStopWordsFromFile("_french.txt", "fr", "\r\n")
}


// TestLoadStopWordsFromFile tries to load a list of stop words from a file
func TestLoadStopWordsFromFile(t *testing.T) {
	source := "I will find you across the seas, above the sky long after this life. Indeed, everybody can do that"
	expected2 := " seas sky life "
	LoadStopWordsFromFile("_stopwords.txt", "en", "\r\n")
	actual := CleanString(source, "en", false)
	if actual != expected2 {
		t.Errorf("Test failed, got: '%s'", actual)
	}
	//Restore the default stop words of the package
	LoadStopWordsFromFile("_english.txt", "en", "\r\n")
	source = "She saw a secret little clearing, and a secret little hot made of rustic poles."
	expected2 = " saw secret little clearing secret little hot rustic poles "
	actual = CleanString(source, "en", false)
	if actual != expected2 {
		t.Errorf("Test failed, got: '%s'", actual)
	}

}
