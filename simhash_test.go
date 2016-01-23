// Copyright 2015 Benjamin BALET. All rights reserved.
// Use of this source code is governed by the BSD license
// license that can be found in the LICENSE file.

package stopwords

import (
	"testing"
)

func TestSimhashFrench(t *testing.T) {
	res1 := Simhash([]byte("la fin d'un bel après-midi d'été"), "fr", false)
	res2 := Simhash([]byte("cet été, nous avons eu un bel après-midi"), "fr", false)
	res3 := Simhash([]byte("l'hiver, la chaleur n'était pas là à midi"), "fr", false)

	if CompareSimhash(res2, res3) < CompareSimhash(res2, res1) {
		t.Errorf("Comparison failed")
	}

	//Hamming distance should be the same
	if CompareSimhash(res1, res2) != CompareSimhash(res2, res1) {
		t.Errorf("Hamming distance should be the same between 1/2 and 2/1")
	}
}
