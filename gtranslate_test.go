package gtranslate

import (
	"testing"
	"time"
)

func TestTranslateWithFromTo(t *testing.T) {
	for i := 0; i < 1; i++ {
		for _, ta := range testingTable {
			translation, err := TranslateWithParams(ta.inText, TranslationParams{
				From:       ta.langFrom,
				To:         ta.langTo,
				Tries:      5,
				Delay:      time.Second,
				GoogleHost: "google.cn",
			})
			if err != nil {
				t.Error(err, err.Error())
				t.Fail()
			}
			if translation.Text != ta.outText {
				t.Errorf("Translated text is not the expected. Expected: %v. Got: %v\n", ta.outText, translation.Text)
			}
			if translation.SourceLanguage != ta.langFrom {
				t.Errorf("Detected language is not the expected. Expected: %v. Got: %v\n", ta.langFrom, translation.SourceLanguage)
			}
		}
	}
}
