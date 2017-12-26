// Copyright 2015 Benjamin BALET. All rights reserved.
// Use of this source code is governed by the BSD license
// license that can be found in the LICENSE file.

package stopwords

import (
	"testing"
)

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
