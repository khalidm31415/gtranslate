package gtranslate

import (
	"fmt"
	"testing"
	"time"
)

type testTable struct {
	inText   string
	langFrom string
	langTo   string
	outText  string
}

var testingTable = []testTable{
	{"Hello", "en", "es", "Hola"},
	{"Bye", "en", "es", "Adiós"},
	{"Hola", "es", "en", "Hi"},
	{"Adios", "es", "en", "Goodbye"},
	{"World", "en", "es", "Mundo"},
}

func TestTranslate(t *testing.T) {
	N := 5
	var totalDur time.Duration
	for i := 0; i < N; i++ {
		for _, ta := range testingTable {
			start := time.Now()
			translation, err := translate(ta.inText, ta.langFrom, ta.langTo, true, 5, time.Second)
			if err != nil {
				t.Error(err.Error())
			}
			if len(translation.Text) < 2 {
				t.Fail()
			}
			dur := time.Since(start)
			totalDur += dur
			if translation.Text != ta.outText {
				t.Errorf("Translated text is not the expected. Expected: %v. Got: %v\n", ta.outText, translation.Text)
			}
			if translation.SourceLanguage != ta.langFrom {
				t.Errorf("Detected language is not the expected. Expected: %v. Got: %v\n", ta.langFrom, translation.SourceLanguage)
			}
		}
	}
	fmt.Println("\nMean time:", time.Duration(int(totalDur)/(len(testingTable)*N)))
}

// TestGetGoogleTranslate is for testing propouse
// func TestGetGoogleTranslate(t *testing.T) {
// 	testText := "Some test text"
// 	for i := 0; i < 4; i++ {
// 		for _, ta := range testingTable {

// 			r, err := getGoogleTranslate(ta.inText, ta.langFrom, ta.langTo)
// 			if err != nil {
// 				t.Error(err.Error())
// 				t.Fail()
// 			}
// 			if r.StatusCode != http.StatusOK {

// 				t.Error("[" + strconv.Itoa(r.StatusCode) + "] failed request with text: '" + testText + "'")
// 			}
// 			if r.Body == nil {
// 				t.Fail()
// 			}
// 		}
// 	}

// }

// // TestRawTranslate testing rawTranslate function
// func TestRawTranslate(t *testing.T) {
// 	for i := 0; i < 4; i++ {
// 		for _, ta := range testingTable {
// 			data, err := rawTranslate(ta.inText, ta.langFrom, ta.langTo)
// 			if err != nil {
// 				t.Error(err.Error())
// 			}
// 			if len(data) < 10 {
// 				t.Fail()
// 			}
// 		}
// 	}

// }
