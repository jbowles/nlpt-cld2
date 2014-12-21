package cld2_nlpt

import "testing"

func TestNlptCLD2EnglishName(t *testing.T) {
	s := "This is an English sentence."
	tag := "ENGLISH"
	buffer_length := len(s)
	det := DetectLanguage(s, buffer_length)
	if tag != det {
		t.Log("Expected: ", tag, "but got... ", det)
		t.Fail()
	}
}

func TestBasicEnglishName(t *testing.T) {
	s := "This is an English sentence."
	tag := "en"
	det := Detect(s)
	if tag != det {
		t.Log("Expected: ", tag, "but got... ", det)
		t.Fail()
	}
}
