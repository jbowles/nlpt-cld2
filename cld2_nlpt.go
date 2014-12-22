// Package cld2 implements language detection using the
// Compact Language Detector.
//
// This package includes the relevant sources from the cld2
// project, so it doesn't require any external dependencies.
// For more information about CLD2, see https://code.google.com/p/cld2/.
package cld2_nlpt

// #include <stdlib.h>
// #include "cld2_nlpt.h"
import "C"

import (
	"fmt"
	"time"
	"unsafe"
)

/*
type LanguageDialect string

type LanguageInfo struct {
	Language Language         // language code; "en"
	Dialect  LanguageDialect  // languace code + dialect; "en-uk"
	Scores   [3]LanguageScore // the 3 most likely languages
	Reliable bool             // is the result reliable?
}

type LanguageScore struct {
	Dialect LanguageDialect
	Percent int // probability/"confidence"
}

func DetectLanguageInfo() LanguageInfo {
	...
}
*/

type Language string
type Cld2NlptError struct {
	When time.Time
	Msg  string
}

func (e Cld2NlptError) Error() string {
	return fmt.Sprint("%v, %v", e.When, e.Msg)
}

// Detect returns the language code for detected language in the given text.
// It uses nlpt_wrapper.h and returns the language code, eg. 'en'.
// C++ function sets the length to the that of the text.
// See nlpt_wrapper tests.

func SimpleDetect(text string) (lang Language, err error) {
	cs := C.CString(text)
	res := C.CLD2_Static_ExtDetectLanguageSummary(cs)
	defer C.free(unsafe.Pointer(cs))
	if res != nil {
		lang = Language(C.GoString(res))
		return lang, err
	} else {
		err = Cld2NlptError{
			time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
			"result returned nil: C.CLD2_DetectSummaryLanguage(cs, -1)",
		}
		return lang, err
	}
	return
}

// DetectLanguage uses nlpt_wrapper.h and returns a format of the output. Choices are 'name' ('ENGLISH'), 'code' ('en'), 'declname' ('ENGLISH').
// See nlpt_wrapper tests.
// If the buffer_length is less than or equal to zero the C++ code will set the length to the that of the text.
//
func DetectLanguage(buffer_length int, text, format string) (lang Language, err error) {
	b_length := C.int(buffer_length)
	var c_char = C.CString("")
	defer C.free(unsafe.Pointer(c_char))
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))
	switch {
	case format == "name":
		c_char = C.CLD2_LanguageName(C.CLD2_DetectLanguage(cs, b_length))
	case format == "code":
		c_char = C.CLD2_LanguageCode(C.CLD2_DetectLanguage(cs, b_length))
	case format == "declname":
		c_char = C.CLD2_LanguageDeclaredName(C.CLD2_DetectLanguage(cs, b_length))
	default:
		c_char = C.CLD2_LanguageCode(C.CLD2_DetectLanguage(cs, b_length))
	}
	if c_char != nil {
		lang = Language(C.GoString(c_char))
		return lang, err
	} else {
		//lang = "corrupt"
		err = Cld2NlptError{
			time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
			"result returned nil: C.CLD2_LanguageName(C.CLD2_DetectLanguage(cs, b_length))",
		}
		return lang, err
	}
	return
}
