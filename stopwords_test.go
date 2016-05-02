// Copyright 2015 Benjamin BALET. All rights reserved.
// Use of this source code is governed by the BSD license
// license that can be found in the LICENSE file.

package stopwords

import (
	"reflect"
	"testing"
)

var expected = " abcdefghijk lmnopqrstuvwxyz "

func TestLangCodesHTML(t *testing.T) {
	var langCodeTests = []struct {
		langCode string
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
		actual := CleanString(source, tt.langCode, false)
		if actual != expected2 {
			t.Errorf("string test failed: expected %q, got %q", expected2, actual)
		}

		actualB := Clean([]byte(source), tt.langCode, false)
		if !reflect.DeepEqual(actualB, []byte(expected2)) {
			t.Errorf("bytes test failed: expected %q, got: %q", []byte(expected2), actualB)
		}
	}
}

func TestRemoveHTML(t *testing.T) {
	source := "<a href='resource.htm'>&quot;Fran &amp; Freddie&#39;s Diner&quot;</a>&nbsp;<b>&lt;tasty@example.com&gt;</b>"
	expected2 := "fran freddie's diner tasty example com "
	actual := CleanString(source, "en", true)
	if actual != expected2 {
		t.Errorf("string test failed: expected %q, got %q", expected2, actual)
	}

	actualB := Clean([]byte(source), "en", true)
	if !reflect.DeepEqual(actualB, []byte(expected2)) {
		t.Errorf("bytes test failed: expected %q, got: %q", []byte(expected2), actualB)
	}
}

func TestRemoveSpaces(t *testing.T) {
	source := "ab cd      ef  gh ij     kl         mn      op qr   st uv    wx yz"
	expected2 := "ab cd ef gh ij kl mn op qr st uv wx yz "
	actual := CleanString(source, "en", true)
	if actual != expected2 {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestToLower(t *testing.T) {
	source := "All Problematic ALmoST Always"
	expected2 := " problematic "
	actual := CleanString(source, "en", true)
	if actual != expected2 {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestArabicStopWords(t *testing.T) {
	source := "كلم abcdefghijk واضافت lmnopqrstuvwxyz اليوم"
	actual := CleanString(source, "ar", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestBulgarianStopWords(t *testing.T) {
	source := "беше abcdefghijk въпреки lmnopqrstuvwxyz юмрук"
	actual := CleanString(source, "bg", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestCzechStopWords(t *testing.T) {
	source := "ačkoli abcdefghijk chceš lmnopqrstuvwxyz neděláš"
	actual := CleanString(source, "cs", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestDanishStopWords(t *testing.T) {
	source := "før abcdefghijk få lmnopqrstuvwxyz hvornår næste"
	actual := CleanString(source, "da", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestDutchStopWords(t *testing.T) {
	source := "aangezien abcdefghijk hierbeneden lmnopqrstuvwxyz ofschoon uitgezonderd"
	actual := CleanString(source, "nl", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestEnglishStopWords(t *testing.T) {
	source := "along abcdefghijk another lmnopqrstuvwxyz yet"
	actual := CleanString(source, "en", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestFrenchStopWords(t *testing.T) {
	source := "assez abcdefghijk certaines lmnopqrstuvwxyz vous-mêmes"
	actual := CleanString(source, "fr", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestFinnishStopWords(t *testing.T) {
	source := "ylös abcdefghijk vähintään lmnopqrstuvwxyz täytyvät kyllä"
	actual := CleanString(source, "fi", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestGermanStopWords(t *testing.T) {
	source := "dafür abcdefghijk dementsprechend lmnopqrstuvwxyz großen zwölf"
	actual := CleanString(source, "de", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestGreekStopWords(t *testing.T) {
	source := "δαίσ abcdefghijk τοιοῦτοσ lmnopqrstuvwxyz ειστε εκεινουσ"
	actual := CleanString(source, "el", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestHungarianStopWords(t *testing.T) {
	source := "alapján abcdefghijk általában lmnopqrstuvwxyz belőle különbözőbb"
	actual := CleanString(source, "hu", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestItalianStopWords(t *testing.T) {
	source := "ahimè abcdefghijk quantunque lmnopqrstuvwxyz perchè tuttavia"
	actual := CleanString(source, "it", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestLatvianStopWords(t *testing.T) {
	source := "būsiet abcdefghijk kļūsim lmnopqrstuvwxyz līdzko tiklīdz"
	actual := CleanString(source, "lv", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestNorwegianStopWords(t *testing.T) {
	source := "fÅ abcdefghijk tilstand lmnopqrstuvwxyz vÖre gjÛre"
	actual := CleanString(source, "no", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestPersianStopWords(t *testing.T) {
	source := "شده abcdefghijk شوند lmnopqrstuvwxyz خواهند بنابراين"
	actual := CleanString(source, "fa", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestPolishStopWords(t *testing.T) {
	source := "daleko abcdefghijk dziś lmnopqrstuvwxyz natychmiast każdy"
	actual := CleanString(source, "pl", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestPortugueseStopWords(t *testing.T) {
	source := "aí abcdefghijk área lmnopqrstuvwxyz vocês vão"
	actual := CleanString(source, "pt", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestRomanianStopWords(t *testing.T) {
	source := "voastră abcdefghijk ţi lmnopqrstuvwxyz aibă"
	actual := CleanString(source, "ro", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestRussianStopWords(t *testing.T) {
	source := "многочисленное abcdefghijk меньше lmnopqrstuvwxyz тринадцать третий"
	actual := CleanString(source, "ru", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestSpanishStopWords(t *testing.T) {
	source := "aquél abcdefghijk cómo lmnopqrstuvwxyz día mucho"
	actual := CleanString(source, "es", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestSlovakStopWords(t *testing.T) {
	source := "budú abcdefghijk každý lmnopqrstuvwxyz môže tvojími"
	actual := CleanString(source, "sk", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestSwedishStopWords(t *testing.T) {
	source := "alltså abcdefghijk åttonde lmnopqrstuvwxyz verkligen likställda"
	actual := CleanString(source, "sv", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestTurkishStopWords(t *testing.T) {
	source := "altmış abcdefghijk yaptığı lmnopqrstuvwxyz ancak beş"
	actual := CleanString(source, "tr", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestIndonesianStopWords(t *testing.T) {
	source := "dia  abcdefghijk lmnopqrstuvwxyz adalah seorang"
	actual := CleanString(source, "id", false)
	if actual != expected {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestUnicodeWordBreakStopWords(t *testing.T) {
	//If the text has been edited with a modern text processor, the words are broken using ZWSP unicode character, e.g.
	//住宅​地域​に​おける​本機​の​使用​は​有害な​電波妨害​を​​引き起こす​こと​が​あり、​その​場合​ユーザー​は​自己負担​で​電波妨害​の​問題​を​解決​しなければなりません​。
	/* This sentence should be broken into the following words (see http://www.atilika.org/):
	   (N) 住宅地域		-
	   (P) に			- stop word
	   (V) おける		- stop word
	   (N) 本機
	   (P) の			- stop word
	   (N) 使用
	   (P) は			- stop word
	   (AN) 有害な
	   (N) 電波妨害
	   (P) を			- stop word
	   (V) 引き起こす
	   (N) こと			- stop word
	   (P) が			- stop word
	   (V) あり、
	   (D) その			- stop word
	   (N) 場合
	   (N) ユーザー
	   (P) は			- stop word
	   (N) 自己負担
	   (P) で			- stop word
	   (N) 電波妨害
	   (P) の			- stop word
	   (N) 問題
	   (P) を			- stop word
	   (N) 解決
	   (V) しなければなりません。
	*/
	source := "住宅​地域​に​おける​本機​の​使用​は​有害な​電波妨害​を​​引き起こす​こと​が​あり、​その​場合​ユーザー​は​自己負担​で​電波妨害​の​問題​を​解決​しなければなりません​。"
	expected2 := "住宅 地域 おける 本機 使用 有害な 電波妨害 引き起こす 場合 ユーザー 自己負担 電波妨害 問題 解決 しなければなりません "
	actual := CleanString(source, "ja", false)
	if actual != expected2 {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}

func TestThaiStopWords(t *testing.T) {
	//As for TestUnicodeWordBreakStopWords, we assume that the words are properly
	//tokenized using ZWSP unicode character or space as a word boundary.
	source := "การ​ที่​ได้​ต้อง​แสดง​ว่า​งาน​ดี"
	expected2 := " แสดง งาน ดี "
	actual := CleanString(source, "th", false)
	if actual != expected2 {
		t.Errorf("Test failed, got: '%s'", actual)
	}
}
