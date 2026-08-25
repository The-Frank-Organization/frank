package jcs

import (
	"bytes"
	"regexp"
	"testing"
)

func TestCanonicalizeRFC8785PrimitiveVector(t *testing.T) {
	input := []byte(`{
  "numbers": [333333333.33333329, 1E30, 4.50,
              2e-3, 0.000000000000000000000000001],
  "string": "\u20ac$\u000F\u000aA'\u0042\u0022\u005c\\\"\/",
  "literals": [null, true, false]
}`)
	want := []byte(`{"literals":[null,true,false],"numbers":[333333333.3333333,1e+30,4.5,0.002,1e-27],"string":"€$\u000f\nA'B\"\\\\\"/"}`)

	got, err := Canonicalize(input)
	if err != nil {
		t.Fatalf("Canonicalize: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("canonical bytes\n got: %s\nwant: %s", got, want)
	}
}

func TestCanonicalizeSortsObjectNamesByUTF16CodeUnits(t *testing.T) {
	input := []byte(`{
  "\u20ac": "Euro Sign",
  "\r": "Carriage Return",
  "\ufb33": "Hebrew Letter Dalet With Dagesh",
  "1": "One",
  "\ud83d\ude00": "Emoji: Grinning Face",
  "\u0080": "Control",
  "\u00f6": "Latin Small Letter O With Diaeresis"
}`)
	want := []byte(`{"\r":"Carriage Return","1":"One","":"Control","ö":"Latin Small Letter O With Diaeresis","€":"Euro Sign","😀":"Emoji: Grinning Face","דּ":"Hebrew Letter Dalet With Dagesh"}`)

	got, err := Canonicalize(input)
	if err != nil {
		t.Fatalf("Canonicalize: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("UTF-16 order\n got: %s\nwant: %s", got, want)
	}
}

func TestCanonicalizeSortsSupplementaryNameBeforeLaterBMPName(t *testing.T) {
	input := []byte(`{"\ue000":1,"\ud83d\ude00":2}`)
	want := []byte(`{"😀":2,"":1}`)

	got, err := Canonicalize(input)
	if err != nil {
		t.Fatalf("Canonicalize: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("surrogate-pair order: got %s, want %s", got, want)
	}
}

func TestCanonicalizeRecursesThroughObjectsInArrays(t *testing.T) {
	input := []byte(`{"z":[{"b":1,"a":2}],"a":{"d":4,"c":3}}`)
	want := []byte(`{"a":{"c":3,"d":4},"z":[{"a":2,"b":1}]}`)

	got, err := Canonicalize(input)
	if err != nil {
		t.Fatalf("Canonicalize: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("recursive order: got %s, want %s", got, want)
	}
}

func TestCanonicalizeUsesECMAScriptNumberThresholds(t *testing.T) {
	input := []byte(`[1e21,1e20,1e-6,1e-7,-0,4.50]`)
	want := []byte(`[1e+21,100000000000000000000,0.000001,1e-7,0,4.5]`)

	got, err := Canonicalize(input)
	if err != nil {
		t.Fatalf("Canonicalize: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("number serialization: got %s, want %s", got, want)
	}
}

func TestCanonicalizeEscapesOnlyRequiredCharacters(t *testing.T) {
	input := []byte(`{"x":"\b\t\n\f\r\u0000\u001f\"\\\/é"}`)
	want := []byte(`{"x":"\b\t\n\f\r\u0000\u001f\"\\/é"}`)

	got, err := Canonicalize(input)
	if err != nil {
		t.Fatalf("Canonicalize: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("string serialization: got %s, want %s", got, want)
	}
}

func TestCanonicalizePreservesUnicodeWithoutNormalization(t *testing.T) {
	input := []byte(`{"composed":"é","decomposed":"e\u0301"}`)
	want := []byte(`{"composed":"é","decomposed":"é"}`)

	got, err := Canonicalize(input)
	if err != nil {
		t.Fatalf("Canonicalize: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("Unicode preservation: got %q, want %q", got, want)
	}
}

func TestCanonicalizeRejectsNonIJSONInput(t *testing.T) {
	cases := map[string][]byte{
		"duplicate object member": []byte(`{"x":1,"x":2}`),
		"nested duplicate":        []byte(`{"x":{"y":1,"y":2}}`),
		"trailing content":        []byte(`{}[]`),
		"raw control":             []byte{'{', '"', 'x', '"', ':', '"', 0x01, '"', '}'},
		"invalid UTF-8":           []byte{'[', '"', 0xff, '"', ']'},
		"lone high surrogate":     []byte(`["\ud800"]`),
		"lone low surrogate":      []byte(`["\udead"]`),
		"number overflow":         []byte(`[1e400]`),
		"leading zero":            []byte(`[01]`),
	}

	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			if got, err := Canonicalize(input); err == nil {
				t.Fatalf("Canonicalize(%q) = %s, want error", input, got)
			}
		})
	}
}

func TestDigestIsLowercaseSHA256AndInputOrderInvariant(t *testing.T) {
	left, err := Digest([]byte(`{"b":2,"a":1}`))
	if err != nil {
		t.Fatalf("Digest left: %v", err)
	}
	right, err := Digest([]byte(`{ "a": 1.0, "b": 2e0 }`))
	if err != nil {
		t.Fatalf("Digest right: %v", err)
	}
	if left != right {
		t.Fatalf("equivalent objects have different digests: %s != %s", left, right)
	}
	if !regexp.MustCompile(`^[0-9a-f]{64}$`).MatchString(left) {
		t.Fatalf("digest %q is not 64 lowercase hexadecimal characters", left)
	}
}
