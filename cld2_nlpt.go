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

type Cld2NlptError struct {
	When time.Time
	Msg  string
}

func (e Cld2NlptError) Error() string {
	return fmt.Sprint("%v, %v", e.When, e.Msg)
}

// Detect returns the language code for detected language
// in the given text.
func Detect(text string) (lang string, err error) {
	cs := C.CString(text)
	res := C.CLD2_DetectSummaryLanguage(cs, -1)
	defer C.free(unsafe.Pointer(cs))
	if res != nil {
		lang = C.GoString(res)
	} else {
		err = Cld2NlptError{
			time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
			"result returned nil: C.CLD2_DetectSummaryLanguage(cs, -1)",
		}
	}
	return
}

// DetectLanguage uses nlpt_wrapper.h and returns the full name of the language in uppercase. See nlpt_wrapper tests TestNlptCLD2EnglishName
// If the buffer_length is less than or equal to zero the C++ code will set the length to the that of the text.
//
func DetectLanguage(buffer_length int, text, format string) (lang string, err error) {
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
	//res := C.CLD2_LanguageName(C.CLD2_DetectLanguage(cs, b_length))
	//res := C.GoString(C.CLD2_LanguageName(C.CLD2_DetectLanguage(cs, b_length)))
	if c_char != nil {
		lang = C.GoString(c_char)
		err = nil
	} else {
		//lang = "corrupt"
		err = Cld2NlptError{
			time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
			"result returned nil: C.CLD2_LanguageName(C.CLD2_DetectLanguage(cs, b_length))",
		}
	}
	return
}
