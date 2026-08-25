package wire

import "testing"

// The m-10 §A.2 canonical-decimal-uint64 STRING rule: grammar ^(0|[1-9][0-9]*)$,
// value < 2^64, numeric decoded comparison; violation ⇒ malformed frame.

func TestParseCounterAccepts(t *testing.T) {
	cases := []struct {
		in   string
		want uint64
	}{
		{"0", 0},
		{"7", 7},
		{"10", 10},
		{"9007199254740993", 9007199254740993}, // > 2^53: must not lose exactness to a float path
		{"18446744073709551615", 18446744073709551615}, // 2^64 - 1, the domain ceiling
	}
	for _, c := range cases {
		got, err := ParseCounter(c.in)
		if err != nil {
			t.Errorf("ParseCounter(%q): unexpected error: %v", c.in, err)
			continue
		}
		if got != c.want {
			t.Errorf("ParseCounter(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

func TestParseCounterRejects(t *testing.T) {
	cases := []string{
		"",                      // empty
		"01",                    // leading zero
		"00",                    // leading zero
		"+1",                    // sign
		"-1",                    // sign
		" 1",                    // leading space
		"1 ",                    // trailing space
		"1\n",                   // trailing newline
		"1.0",                   // not an integer literal
		"1e3",                   // exponent form
		"0x10",                  // hex
		"1_000",                 // go-ism separators
		"١",                     // non-ASCII digit (Arabic-Indic one)
		"abc",                   // not digits
		"18446744073709551616",  // 2^64: out of domain
		"99999999999999999999",  // 20 digits, over 2^64
		"999999999999999999999", // 21 digits, far over
	}
	for _, c := range cases {
		if got, err := ParseCounter(c); err == nil {
			t.Errorf("ParseCounter(%q) = %d, want malformed-counter error", c, got)
		}
	}
}

func TestFormatCounterCanonical(t *testing.T) {
	cases := []struct {
		in   uint64
		want string
	}{
		{0, "0"},
		{7, "7"},
		{9007199254740993, "9007199254740993"},
		{18446744073709551615, "18446744073709551615"},
	}
	for _, c := range cases {
		if got := FormatCounter(c.in); got != c.want {
			t.Errorf("FormatCounter(%d) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestCounterRoundTrip(t *testing.T) {
	for _, v := range []uint64{0, 1, 42, 1 << 53, 18446744073709551615} {
		got, err := ParseCounter(FormatCounter(v))
		if err != nil || got != v {
			t.Errorf("round-trip %d: got %d, err %v", v, got, err)
		}
	}
}
