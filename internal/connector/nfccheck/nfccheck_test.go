package nfccheck

import "testing"

func TestIsNormalStringDistinguishesComposedAndDecomposedUnicode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		text string
		want bool
	}{
		{name: "ASCII", text: "lane-codex-1", want: true},
		{name: "composed non-ASCII", text: "caf\u00e9", want: true},
		{name: "decomposed non-ASCII", text: "cafe\u0301", want: false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := IsNormalString(test.text); got != test.want {
				t.Fatalf("IsNormalString(%q) = %v, want %v", test.text, got, test.want)
			}
		})
	}
}
