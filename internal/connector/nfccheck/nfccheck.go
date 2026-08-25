// Package nfccheck provides check-only Unicode NFC validation for schema
// fields whose contracts explicitly require NFC. It never rewrites text.
package nfccheck

import "golang.org/x/text/unicode/norm"

// IsNormalString reports whether text is already in Unicode NFC form.
func IsNormalString(text string) bool {
	return norm.NFC.IsNormalString(text)
}
