package cld2_nlpt

import "testing"

func TestDetectEnglish(t *testing.T) {
	s := "This is an English sentence."
	tag := Language("en")
	lang, err := SimpleDetect(s)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang, "with ERROR:", err)
		t.Fail()
	}
}

func TestEnglishName(t *testing.T) {
	s := "This is an English sentence."
	tag := Language("ENGLISH")
	buffer_length := len(s)
	format := "name"
	lang, err := DetectLanguage(buffer_length, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestEnglishCode(t *testing.T) {
	s := "This is an English sentence."
	tag := Language("en")
	buffer_length := len(s)
	format := "code"
	lang, err := DetectLanguage(buffer_length, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestEnglishDeclaredName(t *testing.T) {
	s := "This is an English sentence."
	tag := Language("ENGLISH")
	buffer_length := len(s)
	format := "declname"
	lang, err := DetectLanguage(buffer_length, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

// This directory seeks to provide a comprehensive overview of non-profit organizations and community groups that are active in and for Myanmar (Burma).
func TestBurmeseName(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := Language("BURMESE")
	buffer_length := len(s)
	format := "name"
	lang, err := DetectLanguage(buffer_length, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

// This directory seeks to provide a comprehensive overview of non-profit organizations and community groups that are active in and for Myanmar (Burma).
func TestBurmeseCode(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := Language("my")
	buffer_length := len(s)
	format := "code"
	lang, err := DetectLanguage(buffer_length, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

// This directory seeks to provide a comprehensive overview of non-profit organizations and community groups that are active in and for Myanmar (Burma).
func TestBurmeseDeclaredName(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := Language("BURMESE")
	buffer_length := len(s)
	format := "declname"
	lang, err := DetectLanguage(buffer_length, s, format)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestExtendedEnglishName(t *testing.T) {
	s := "This is an English sentence."
	tag := Language("ENGLISH")
	buffer_length := len(s)
	format := "name"
	lang, err := DetectExtendedLanguage(s, format, buffer_length, 3, 3, 3)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestExtendedEnglishCode(t *testing.T) {
	s := "This is an English sentence."
	tag := Language("en")
	buffer_length := len(s)
	format := "code"
	lang, err := DetectExtendedLanguage(s, format, buffer_length, 3, 3, 3)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

func TestExtendedEnglishDeclaredName(t *testing.T) {
	s := "This is an English sentence."
	tag := Language("ENGLISH")
	buffer_length := len(s)
	format := "declname"
	lang, err := DetectExtendedLanguage(s, format, buffer_length, 3, 3, 3)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

// This directory seeks to provide a comprehensive overview of non-profit organizations and community groups that are active in and for Myanmar (Burma).
func TestExtendedBurmeseName(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := Language("BURMESE")
	buffer_length := len(s)
	format := "name"
	lang, err := DetectExtendedLanguage(s, format, buffer_length, 3, 3, 3)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

// This directory seeks to provide a comprehensive overview of non-profit organizations and community groups that are active in and for Myanmar (Burma).
func TestExtendedBurmeseCode(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := Language("my")
	buffer_length := len(s)
	format := "code"
	lang, err := DetectExtendedLanguage(s, format, buffer_length, 3, 3, 3)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}

// This directory seeks to provide a comprehensive overview of non-profit organizations and community groups that are active in and for Myanmar (Burma).
func TestExtendedBurmeseDeclaredName(t *testing.T) {
	s := "ဤလမ်းညွှန်သည် မြန်မာနိုင်ငံအတွင်းနှင့် မြန်မာနိုင်ငံအတွက် အကျိူးအမြတ်မယူဘဲ ဆောင်ရွက်ပေးနေသော အဖွဲ့အစည်းများ၏ ပြည့်စုံသော ယေဘုယျအချက်အလက်များကို ရှာဖွေးပေးပါသည်။"
	tag := Language("BURMESE")
	buffer_length := len(s)
	format := "declname"
	lang, err := DetectExtendedLanguage(s, format, buffer_length, 3, 3, 3)
	if tag != lang || err != nil {
		t.Log("Expected: ", tag, "but got... ", lang)
		t.Fail()
	}
}
