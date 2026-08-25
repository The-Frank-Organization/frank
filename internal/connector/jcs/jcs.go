// Package jcs implements the JSON Canonicalization Scheme defined by RFC 8785.
package jcs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"unicode/utf16"
	"unicode/utf8"
)

// Canonicalize parses one JSON value and returns its RFC 8785 encoding.
// It rejects duplicate object names, invalid UTF-8, and lone UTF-16
// surrogates instead of accepting encoding/json's replacement behavior.
func Canonicalize(input []byte) ([]byte, error) {
	if !utf8.Valid(input) {
		return nil, fmt.Errorf("jcs: input is not valid UTF-8")
	}
	if err := validateStringTokens(input); err != nil {
		return nil, err
	}

	dec := json.NewDecoder(bytes.NewReader(input))
	dec.UseNumber()
	v, err := parseValue(dec)
	if err != nil {
		return nil, err
	}
	if _, err := dec.Token(); err != io.EOF {
		if err == nil {
			return nil, fmt.Errorf("jcs: trailing JSON value")
		}
		return nil, fmt.Errorf("jcs: trailing data: %w", err)
	}

	out, err := appendValue(nil, v)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IsCanonical reports whether input is exactly one valid RFC 8785 value.
func IsCanonical(input []byte) bool {
	canonical, err := Canonicalize(input)
	return err == nil && bytes.Equal(input, canonical)
}

type valueKind uint8

const (
	nullKind valueKind = iota
	boolKind
	numberKind
	stringKind
	arrayKind
	objectKind
)

type value struct {
	kind    valueKind
	boolean bool
	number  string
	text    string
	array   []value
	object  []member
}

type member struct {
	name  string
	value value
}

func parseValue(dec *json.Decoder) (value, error) {
	tok, err := dec.Token()
	if err != nil {
		return value{}, fmt.Errorf("jcs: parse JSON: %w", err)
	}

	switch tok := tok.(type) {
	case nil:
		return value{kind: nullKind}, nil
	case bool:
		return value{kind: boolKind, boolean: tok}, nil
	case string:
		return value{kind: stringKind, text: tok}, nil
	case json.Number:
		return value{kind: numberKind, number: tok.String()}, nil
	case json.Delim:
		switch tok {
		case '[':
			var items []value
			for dec.More() {
				item, err := parseValue(dec)
				if err != nil {
					return value{}, err
				}
				items = append(items, item)
			}
			end, err := dec.Token()
			if err != nil || end != json.Delim(']') {
				return value{}, fmt.Errorf("jcs: malformed array")
			}
			return value{kind: arrayKind, array: items}, nil
		case '{':
			seen := make(map[string]struct{})
			var members []member
			for dec.More() {
				nameToken, err := dec.Token()
				if err != nil {
					return value{}, fmt.Errorf("jcs: read object name: %w", err)
				}
				name, ok := nameToken.(string)
				if !ok {
					return value{}, fmt.Errorf("jcs: object name is not a string")
				}
				if _, exists := seen[name]; exists {
					return value{}, fmt.Errorf("jcs: duplicate object name %q", name)
				}
				seen[name] = struct{}{}
				memberValue, err := parseValue(dec)
				if err != nil {
					return value{}, err
				}
				members = append(members, member{name: name, value: memberValue})
			}
			end, err := dec.Token()
			if err != nil || end != json.Delim('}') {
				return value{}, fmt.Errorf("jcs: malformed object")
			}
			return value{kind: objectKind, object: members}, nil
		default:
			return value{}, fmt.Errorf("jcs: unexpected delimiter %q", tok)
		}
	default:
		return value{}, fmt.Errorf("jcs: unsupported JSON token %T", tok)
	}
}

func appendValue(dst []byte, v value) ([]byte, error) {
	switch v.kind {
	case nullKind:
		return append(dst, "null"...), nil
	case boolKind:
		return strconv.AppendBool(dst, v.boolean), nil
	case numberKind:
		number, err := canonicalNumber(v.number)
		if err != nil {
			return nil, err
		}
		return append(dst, number...), nil
	case stringKind:
		return appendString(dst, v.text), nil
	case arrayKind:
		dst = append(dst, '[')
		for i, item := range v.array {
			if i > 0 {
				dst = append(dst, ',')
			}
			var err error
			dst, err = appendValue(dst, item)
			if err != nil {
				return nil, err
			}
		}
		return append(dst, ']'), nil
	case objectKind:
		sort.Slice(v.object, func(i, j int) bool {
			return utf16Less(v.object[i].name, v.object[j].name)
		})
		dst = append(dst, '{')
		for i, item := range v.object {
			if i > 0 {
				dst = append(dst, ',')
			}
			dst = appendString(dst, item.name)
			dst = append(dst, ':')
			var err error
			dst, err = appendValue(dst, item.value)
			if err != nil {
				return nil, err
			}
		}
		return append(dst, '}'), nil
	default:
		return nil, fmt.Errorf("jcs: invalid internal value kind %d", v.kind)
	}
}

func canonicalNumber(raw string) (string, error) {
	f, err := strconv.ParseFloat(raw, 64)
	if err != nil || math.IsInf(f, 0) || math.IsNaN(f) {
		return "", fmt.Errorf("jcs: number %q is not an IEEE 754 finite value", raw)
	}
	if f == 0 {
		return "0", nil
	}

	format := byte('e')
	abs := math.Abs(f)
	if abs >= 1e-6 && abs < 1e21 {
		format = 'f'
	}
	formatted := strconv.FormatFloat(f, format, -1, 64)
	if format == 'e' {
		formatted = trimExponentZeroes(formatted)
	}
	return formatted, nil
}

func trimExponentZeroes(number string) string {
	e := bytes.IndexByte([]byte(number), 'e')
	if e < 0 || e+2 >= len(number) {
		return number
	}
	digits := e + 2
	for digits+1 < len(number) && number[digits] == '0' {
		digits++
	}
	if digits == e+2 {
		return number
	}
	return number[:e+2] + number[digits:]
}

func appendString(dst []byte, s string) []byte {
	const hex = "0123456789abcdef"
	dst = append(dst, '"')
	for i := 0; i < len(s); i++ {
		b := s[i]
		switch b {
		case '"', '\\':
			dst = append(dst, '\\', b)
		case '\b':
			dst = append(dst, '\\', 'b')
		case '\t':
			dst = append(dst, '\\', 't')
		case '\n':
			dst = append(dst, '\\', 'n')
		case '\f':
			dst = append(dst, '\\', 'f')
		case '\r':
			dst = append(dst, '\\', 'r')
		default:
			if b < 0x20 {
				dst = append(dst, '\\', 'u', '0', '0', hex[b>>4], hex[b&0x0f])
			} else {
				dst = append(dst, b)
			}
		}
	}
	return append(dst, '"')
}

func utf16Less(left, right string) bool {
	a := utf16.Encode([]rune(left))
	b := utf16.Encode([]rune(right))
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func validateStringTokens(input []byte) error {
	inString := false
	for i := 0; i < len(input); i++ {
		b := input[i]
		if !inString {
			if b == '"' {
				inString = true
			}
			continue
		}

		switch b {
		case '"':
			inString = false
		case '\\':
			if i+1 >= len(input) {
				return fmt.Errorf("jcs: unterminated string escape")
			}
			escape := input[i+1]
			if escape != 'u' {
				if !bytes.ContainsRune([]byte(`\"\\/bfnrt`), rune(escape)) {
					return fmt.Errorf("jcs: invalid string escape \\%c", escape)
				}
				i++
				continue
			}

			first, err := parseHex16(input, i+2)
			if err != nil {
				return err
			}
			switch {
			case first >= 0xd800 && first <= 0xdbff:
				if i+12 > len(input) || input[i+6] != '\\' || input[i+7] != 'u' {
					return fmt.Errorf("jcs: high surrogate is not followed by a low surrogate")
				}
				second, err := parseHex16(input, i+8)
				if err != nil {
					return err
				}
				if second < 0xdc00 || second > 0xdfff {
					return fmt.Errorf("jcs: high surrogate is not followed by a low surrogate")
				}
				i += 11
			case first >= 0xdc00 && first <= 0xdfff:
				return fmt.Errorf("jcs: lone low surrogate")
			default:
				i += 5
			}
		default:
			if b < 0x20 {
				return fmt.Errorf("jcs: unescaped control byte in string")
			}
		}
	}
	if inString {
		return fmt.Errorf("jcs: unterminated string")
	}
	return nil
}

func parseHex16(input []byte, start int) (uint16, error) {
	if start+4 > len(input) {
		return 0, fmt.Errorf("jcs: short Unicode escape")
	}
	value, err := strconv.ParseUint(string(input[start:start+4]), 16, 16)
	if err != nil {
		return 0, fmt.Errorf("jcs: invalid Unicode escape: %w", err)
	}
	return uint16(value), nil
}
