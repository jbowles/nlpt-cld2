package cld2

import (
	"testing"

	"github.com/jbowles/cld2_nlpt"
)

func TestDetectEnglish(t *testing.T) {
	s := "This is an English sentence."
	tag := cld2.Language("en")
	lang, err := cld2.SimpleDetect(s)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang, "with ERROR:", err)
		t.Fail()
	}
}

/*
func TestEnglishName(t *testing.T) {
	s := "This is an English sentence."
	tag := Language("ENGLISH")
	bufferLength := len(s)
	format := "name"
	lang, err := DetectLanguage(bufferLength, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestEnglishCode(t *testing.T) {
	s := "This is an English sentence."
	tag := Language("en")
	bufferLength := len(s)
	format := "code"
	lang, err := DetectLanguage(bufferLength, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestEnglishDeclaredName(t *testing.T) {
	s := "This is an English sentence."
	tag := Language("ENGLISH")
	bufferLength := len(s)
	format := "declname"
	lang, err := DetectLanguage(bufferLength, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestBurmeseName(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := Language("BURMESE")
	bufferLength := len(s)
	format := "name"
	lang, err := DetectLanguage(bufferLength, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestBurmeseCode(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := Language("my")
	bufferLength := len(s)
	format := "code"
	lang, err := DetectLanguage(bufferLength, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestBurmeseDeclaredName(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := Language("BURMESE")
	bufferLength := len(s)
	format := "declname"
	lang, err := DetectLanguage(bufferLength, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestExtendedEnglishName(t *testing.T) {
	s := "This is an English sentence."
	tag := Language("ENGLISH")
	bufferLength := len(s)
	format := "name"
	lang, err := DetectExtendedLanguage(s, format, bufferLength, 3, 3, 3)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestExtendedEnglishCode(t *testing.T) {
	s := "This is an English sentence."
	tag := Language("en")
	bufferLength := len(s)
	format := "code"
	lang, err := DetectExtendedLanguage(s, format, bufferLength, 3, 3, 3)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestExtendedEnglishDeclaredName(t *testing.T) {
	s := "This is an English sentence."
	tag := Language("ENGLISH")
	bufferLength := len(s)
	format := "declname"
	lang, err := DetectExtendedLanguage(s, format, bufferLength, 3, 3, 3)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestExtendedBurmeseName(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := Language("BURMESE")
	bufferLength := len(s)
	format := "name"
	lang, err := DetectExtendedLanguage(s, format, bufferLength, 3, 3, 3)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestExtendedBurmeseCode(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := Language("my")
	bufferLength := len(s)
	format := "code"
	lang, err := DetectExtendedLanguage(s, format, bufferLength, 3, 3, 3)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestExtendedBurmeseDeclaredName(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := Language("BURMESE")
	bufferLength := len(s)
	format := "declname"
	lang, err := DetectExtendedLanguage(s, format, bufferLength, 3, 3, 3)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}
*/
