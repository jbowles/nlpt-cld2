package cld2_nlpt

import "testing"

const (
	PRINTALL = false
)

func printAll(t *testing.T, tag, format, lang string) {
	if PRINTALL {
		t.Log("  tag:", tag, "  format:", format, "  correct format returned:", lang)
	}
}

func TestDetectEnglish(t *testing.T) {
	s := "This is an English sentence."
	tag := "en"
	lang, err := Detect(s)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang, "with ERROR:", err)
		t.Fail()
	}
}

func TestEnglishName(t *testing.T) {
	s := "This is an English sentence."
	tag := "ENGLISH"
	buffer_length := len(s)
	format := "name"
	lang, err := DetectLanguage(buffer_length, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
	printAll(t, tag, format, lang)
}

func TestEnglishCode(t *testing.T) {
	s := "This is an English sentence."
	tag := "en"
	buffer_length := len(s)
	format := "code"
	lang, err := DetectLanguage(buffer_length, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
	printAll(t, tag, format, lang)
}

func TestEnglishDeclaredName(t *testing.T) {
	s := "This is an English sentence."
	tag := "ENGLISH"
	buffer_length := len(s)
	format := "declname"
	lang, err := DetectLanguage(buffer_length, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
	printAll(t, tag, format, lang)
}

// This directory seeks to provide a comprehensive overview of non-profit organizations and community groups that are active in and for Myanmar (Burma).
func TestBurmeseName(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := "BURMESE"
	buffer_length := len(s)
	format := "name"
	lang, err := DetectLanguage(buffer_length, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
	printAll(t, tag, format, lang)
}

// This directory seeks to provide a comprehensive overview of non-profit organizations and community groups that are active in and for Myanmar (Burma).
func TestBurmeseCode(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := "my"
	buffer_length := len(s)
	format := "code"
	lang, err := DetectLanguage(buffer_length, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
	printAll(t, tag, format, lang)
}

// This directory seeks to provide a comprehensive overview of non-profit organizations and community groups that are active in and for Myanmar (Burma).
func TestBurmeseDeclaredName(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := "BURMESE"
	buffer_length := len(s)
	format := "declname"
	lang, err := DetectLanguage(buffer_length, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
	printAll(t, tag, format, lang)
}
