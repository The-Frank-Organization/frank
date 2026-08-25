package appipc

import "testing"

func TestCounterWireGrammarAndUint64Boundary(t *testing.T) {
	for _, valid := range []string{"0", "1", "42", "18446744073709551615"} {
		value, err := ParseCounter(valid)
		if err != nil {
			t.Fatalf("ParseCounter(%q): %v", valid, err)
		}
		if got := FormatCounter(value); got != valid {
			t.Fatalf("FormatCounter(ParseCounter(%q)) = %q", valid, got)
		}
	}

	for _, invalid := range []string{
		"", "00", "01", "+1", "-1", " 1", "1 ", "1.0", "1a",
		"18446744073709551616",
	} {
		if _, err := ParseCounter(invalid); err == nil {
			t.Fatalf("ParseCounter(%q) succeeded", invalid)
		}
	}
}

func TestCounterStoreCodecUsesFixedTwentyDigitPadding(t *testing.T) {
	cases := map[string]string{
		"0":                    "00000000000000000000",
		"42":                   "00000000000000000042",
		"18446744073709551615": "18446744073709551615",
	}
	for wire, stored := range cases {
		got, err := PadCounter(wire)
		if err != nil {
			t.Fatalf("PadCounter(%q): %v", wire, err)
		}
		if got != stored {
			t.Fatalf("PadCounter(%q) = %q, want %q", wire, got, stored)
		}
		roundTrip, err := UnpadCounter(got)
		if err != nil {
			t.Fatalf("UnpadCounter(%q): %v", got, err)
		}
		if roundTrip != wire {
			t.Fatalf("UnpadCounter(PadCounter(%q)) = %q", wire, roundTrip)
		}
	}

	for _, invalid := range []string{
		"0000000000000000000",
		"000000000000000000000",
		"0000000000000000000a",
		"18446744073709551616",
	} {
		if _, err := UnpadCounter(invalid); err == nil {
			t.Fatalf("UnpadCounter(%q) succeeded", invalid)
		}
	}
}
