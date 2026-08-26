// Package jcs exposes the worker compatibility surface over Frank's one
// strict-integer canonical JSON implementation.
package jcs

import "github.com/jackli/frank/internal/canonicaljson"

func Canonicalize(input []byte) ([]byte, error) { return canonicaljson.Canonicalize(input) }
func Digest(input []byte) (string, error)       { return canonicaljson.Digest(input) }
