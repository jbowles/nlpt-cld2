// Package cld2 implements language detection using the
// Compact Language Detector.
//
// This package includes the relevant sources from the cld2
// project, so it doesn't require any external dependencies.
// For more information about CLD2, see https://code.google.com/p/cld2/.
package cld2_nlpt

// #include <stdlib.h>
// #include "cld2_min/cld2_nlpt.h"
import "C"
import "unsafe"

// DetectLanguage uses nlpt_wrapper.h and returns the full name of the language in uppercase. See nlpt_wrapper tests TestNlptCLD2EnglishName
// If the buffer_length is less than or equal to zero the C++ code will set the length to the that of the text.
//
func DetectLanguage(text string, buffer_length int) string {
	b_length := C.int(buffer_length)
	cs := C.CString(text)
	res := C.CLD2_DetectLanguage(cs, b_length)
	defer C.free(unsafe.Pointer(cs))
	//var lang string
	//if res != nil {
	//	lang = C.GoString(res)
	//}
	return C.GoString(C.CLD2_LanguageName(res))
}
