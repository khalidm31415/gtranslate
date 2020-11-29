package gtranslate

import (
	"time"

	"golang.org/x/text/language"
)

var GoogleHost = "google.com"

// TranslationParams is a util struct to pass as parameter to indicate how to translate
type TranslationParams struct {
	From       string
	To         string
	Tries      int
	Delay      time.Duration
	GoogleHost string
}

// Translate translate a text using native tags offer by go language
func Translate(text string, from language.Tag, to language.Tag, googleHost ...string) (Translation, error) {
	if len(googleHost) != 0 && googleHost[0] != "" {
		GoogleHost = googleHost[0]
	}
	translation, err := translate(text, from.String(), to.String(), false, 2, 0)
	if err != nil {
		return Translation{}, err
	}

	return translation, nil
}

// TranslateWithParams translate a text with simple params as string
func TranslateWithParams(text string, params TranslationParams) (Translation, error) {
	if params.GoogleHost == "" {
		GoogleHost = "google.com"
	} else {
		GoogleHost = params.GoogleHost
	}
	translation, err := translate(text, params.From, params.To, true, params.Tries, params.Delay)
	if err != nil {
		return Translation{}, err
	}
	return translation, nil
}
