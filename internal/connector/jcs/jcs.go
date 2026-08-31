// Package jcs exposes the connector compatibility surface over Frank's one
// strict-integer canonical JSON implementation.
package jcs

import "github.com/The-Frank-Organization/frank/internal/canonicaljson"

func Canonicalize(input []byte) ([]byte, error) { return canonicaljson.Canonicalize(input) }
func IsCanonical(input []byte) bool             { return canonicaljson.IsCanonical(input) }
