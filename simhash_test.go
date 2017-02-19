// Copyright 2015 Benjamin BALET. All rights reserved.
// Use of this source code is governed by the BSD license
// license that can be found in the LICENSE file.

package stopwords

import (
	"testing"
)

//French language
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

//English language
func TestSimhashEnglish(t *testing.T) {
	res1 := Simhash([]byte("the war is coming, send more troop. This soldier was courageous."), "en", false)
	res2 := Simhash([]byte("the troop was composed of veteran soldiers of the last war."), "en", false)
	res3 := Simhash([]byte("the is an article of the english language it was created long before the last reform."), "en", false)

	if CompareSimhash(res2, res3) < CompareSimhash(res2, res1) {
		t.Errorf("Comparison failed")
	}

	//Hamming distance should be the same
	if CompareSimhash(res1, res2) != CompareSimhash(res2, res1) {
		t.Errorf("Hamming distance should be the same between 1/2 and 2/1")
	}
}

//Khmer language
func TestSimhashKhmer(t *testing.T) {
	//Articles about election
	res1 := Simhash([]byte("គណៈកម្មាធិការជាតិរៀបចំការបោះឆ្នោត​(គ.ជ.ប) ប្រកាស​ទទួល​ពាក្យ​ស្នើ​​សុំ​សង្កេត​ការណ៍​បោះឆ្នោត​ពី​អ្នក​សង្កេត​ការណ៍​ជាតិ អន្តរជាតិ និង​គណបក្ស​នយោបាយ ខណៈ​ការ​បោះឆ្នោត​ក្រុម​ប្រឹក្សា​ឃុំ​សង្កាត់​អាណត្តិ​ទី ៤ កាន់តែ​ជិត​មក​ដល់​។"), "km", false)
	res2 := Simhash([]byte("អ្នក​ដាក់​ពាក្យ​ស្នើ​សុំ​សង្កេត​ការណ៍​ជាតិ និង​អន្តរ​ជាតិ​អាច​យក​ពាក្យ​ស្នើ​សុំ​នៅ​ការិយា​ល័យ គ.ជ.ប ឬ​ទាញ​យក​ពាក្យ​តាម​គេហទំព័រ​​របស់​គ.ជ.ប និង​ដាក់​ពាក្យ​វិញ​នៅ​ការិយាល័យ​គ.ជ.ប​ដដែល​។"), "km", false)
	//Description of Cambodia
	res3 := Simhash([]byte("​កម្ពុជា​ឬ​ឈ្មោះ​ផ្លូវការ​ព្រះរាជាណាចក្រកម្ពុជា​ជា​ប្រទេស​មួយ​ស្ថិតនៅ​ផ្នែក​ខាងត្បូង​នៃ​ទៀបកោះ​ឥណ្ឌូចិន​នៅ​អាស៊ី​អាគ្នេយ៍​។"), "km", false)

	if CompareSimhash(res2, res3) < CompareSimhash(res2, res1) {
		t.Errorf("Comparison failed")
	}

	//Hamming distance should be the same
	if CompareSimhash(res1, res2) != CompareSimhash(res2, res1) {
		t.Errorf("Hamming distance should be the same between 1/2 and 2/1")
	}
}
