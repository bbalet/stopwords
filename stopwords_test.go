// Copyright 2015 Benjamin BALET. All rights reserved.
// Use of this source code is governed by the BSD license
// license that can be found in the LICENSE file.

package stopwords

import (
	"testing"
)

var expected = " abcdefghijk lmnopqrstuvwxyz "

func TestLangCodesHTML(t *testing.T) {
	var langCodeTests = []struct {
	  langCode        string
	}{
	  {"en-US"},
	  {"en_US"},
	  {"en_GB"},
	  {"en-GB"},
	  {"en"},
		{"eng"},
		{"zz"},
	}

	source := "afterwards"
	expected2 := " "
	for _, tt := range langCodeTests {
		actual := CleanContent(source, tt.langCode, false)
		if actual != expected2 {
			t.Errorf("Test failed, got: '%s'", actual)
		}
	}
}

func TestRemoveHTML(t *testing.T) {
	source := "<a href='resource.htm'>&quot;Fran &amp; Freddie&#39;s Diner&quot;</a>&nbsp;<b>&lt;tasty@example.com&gt;</b>"
	expected2 := "fran freddie's diner tasty example com "
	actual := CleanContent(source, "en", true)
	if actual != expected2 {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestRemoveSpaces(t *testing.T) {
	source := "ab cd      ef  gh ij     kl         mn      op qr   st uv    wx yz"
	expected2 := "ab cd ef gh ij kl mn op qr st uv wx yz "
	actual := CleanContent(source, "en", true)
	if actual != expected2 {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestToLower(t *testing.T) {
	source := "All Problematic ALmoST Always"
	expected2 := " problematic "
	actual := CleanContent(source, "en", true)
	if actual != expected2 {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestArabicStopWords(t *testing.T) {
	source := "كلم abcdefghijk واضافت lmnopqrstuvwxyz اليوم"
	actual := CleanContent(source, "ar", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestBulgarianStopWords(t *testing.T) {
	source := "беше abcdefghijk въпреки lmnopqrstuvwxyz юмрук"
	actual := CleanContent(source, "bg", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestCzechStopWords(t *testing.T) {
	source := "ačkoli abcdefghijk chceš lmnopqrstuvwxyz neděláš"
	actual := CleanContent(source, "cs", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestDanishStopWords(t *testing.T) {
	source := "før abcdefghijk få lmnopqrstuvwxyz hvornår næste"
	actual := CleanContent(source, "da", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestDutchStopWords(t *testing.T) {
	source := "aangezien abcdefghijk hierbeneden lmnopqrstuvwxyz ofschoon uitgezonderd"
	actual := CleanContent(source, "nl", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestEnglishStopWords(t *testing.T) {
	source := "along abcdefghijk another lmnopqrstuvwxyz yet"
	actual := CleanContent(source, "en", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestFrenchStopWords(t *testing.T) {
	source := "assez abcdefghijk certaines lmnopqrstuvwxyz vous-mêmes"
	actual := CleanContent(source, "fr", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestFinnishStopWords(t *testing.T) {
	source := "ylös abcdefghijk vähintään lmnopqrstuvwxyz täytyvät kyllä"
	actual := CleanContent(source, "fi", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestGermanStopWords(t *testing.T) {
	source := "dafür abcdefghijk dementsprechend lmnopqrstuvwxyz großen zwölf"
	actual := CleanContent(source, "de", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestGreekStopWords(t *testing.T) {
	source := "δαίσ abcdefghijk τοιοῦτοσ lmnopqrstuvwxyz ειστε εκεινουσ"
	actual := CleanContent(source, "el", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestHungarianStopWords(t *testing.T) {
	source := "alapján abcdefghijk általában lmnopqrstuvwxyz belőle különbözőbb"
	actual := CleanContent(source, "hu", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestItalianStopWords(t *testing.T) {
	source := "ahimè abcdefghijk quantunque lmnopqrstuvwxyz perchè tuttavia"
	actual := CleanContent(source, "it", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestLatvianStopWords(t *testing.T) {
	source := "būsiet abcdefghijk kļūsim lmnopqrstuvwxyz līdzko tiklīdz"
	actual := CleanContent(source, "lv", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestNorwegianStopWords(t *testing.T) {
	source := "fÅ abcdefghijk tilstand lmnopqrstuvwxyz vÖre gjÛre"
	actual := CleanContent(source, "no", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestPersianStopWords(t *testing.T) {
	source := "شده abcdefghijk شوند lmnopqrstuvwxyz خواهند بنابراين"
	actual := CleanContent(source, "fa", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestPolishStopWords(t *testing.T) {
	source := "daleko abcdefghijk dziś lmnopqrstuvwxyz natychmiast każdy"
	actual := CleanContent(source, "pl", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestPortugueseStopWords(t *testing.T) {
	source := "aí abcdefghijk área lmnopqrstuvwxyz vocês vão"
	actual := CleanContent(source, "pt", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestRomanianStopWords(t *testing.T) {
	source := "voastră abcdefghijk ţi lmnopqrstuvwxyz aibă"
	actual := CleanContent(source, "ro", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestRussianStopWords(t *testing.T) {
	source := "многочисленное abcdefghijk меньше lmnopqrstuvwxyz тринадцать третий"
	actual := CleanContent(source, "ru", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestSpanishStopWords(t *testing.T) {
	source := "aquél abcdefghijk cómo lmnopqrstuvwxyz día mucho"
	actual := CleanContent(source, "es", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestSlovakStopWords(t *testing.T) {
	source := "budú abcdefghijk každý lmnopqrstuvwxyz môže tvojími"
	actual := CleanContent(source, "sk", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestSwedishStopWords(t *testing.T) {
	source := "alltså abcdefghijk åttonde lmnopqrstuvwxyz verkligen likställda"
	actual := CleanContent(source, "sv", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestTurkishStopWords(t *testing.T) {
	source := "altmış abcdefghijk yaptığı lmnopqrstuvwxyz ancak beş"
	actual := CleanContent(source, "tr", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}
