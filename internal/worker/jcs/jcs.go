// Package jcs implements the RFC 8785 JSON Canonicalization Scheme.
package jcs

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
	"unicode/utf16"
	"unicode/utf8"
)

// Canonicalize parses an I-JSON value and returns its RFC 8785 canonical UTF-8
// encoding. It rejects duplicate object names, invalid Unicode, non-finite
// numbers, and trailing content.
func Canonicalize(input []byte) ([]byte, error) {
	parser := jsonParser{input: input}
	parsed, err := parser.parseValue()
	if err != nil {
		return nil, err
	}
	parser.skipWhitespace()
	if parser.offset != len(parser.input) {
		return nil, parser.errorf("trailing content")
	}

	var output bytes.Buffer
	if err := parsed.appendCanonical(&output); err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}

// Digest returns the lowercase hexadecimal SHA-256 digest of the RFC 8785
// canonical encoding of input.
func Digest(input []byte) (string, error) {
	canonical, err := Canonicalize(input)
	if err != nil {
		return "", err
	}
	digest := sha256.Sum256(canonical)
	return hex.EncodeToString(digest[:]), nil
}

type valueKind uint8

const (
	kindNull valueKind = iota
	kindFalse
	kindTrue
	kindNumber
	kindString
	kindArray
	kindObject
)

type jsonValue struct {
	kind       valueKind
	number     float64
	text       string
	array      []jsonValue
	properties []jsonProperty
}

type jsonProperty struct {
	name  string
	value jsonValue
}

func (value jsonValue) appendCanonical(output *bytes.Buffer) error {
	switch value.kind {
	case kindNull:
		output.WriteString("null")
	case kindFalse:
		output.WriteString("false")
	case kindTrue:
		output.WriteString("true")
	case kindNumber:
		if value.number == 0 {
			output.WriteByte('0')
			return nil
		}
		encoded, err := json.Marshal(value.number)
		if err != nil {
			return fmt.Errorf("serialize number: %w", err)
		}
		output.Write(encoded)
	case kindString:
		appendString(output, value.text)
	case kindArray:
		output.WriteByte('[')
		for index, element := range value.array {
			if index > 0 {
				output.WriteByte(',')
			}
			if err := element.appendCanonical(output); err != nil {
				return err
			}
		}
		output.WriteByte(']')
	case kindObject:
		properties := append([]jsonProperty(nil), value.properties...)
		sort.Slice(properties, func(left, right int) bool {
			return utf16Less(properties[left].name, properties[right].name)
		})
		output.WriteByte('{')
		for index, property := range properties {
			if index > 0 {
				output.WriteByte(',')
			}
			appendString(output, property.name)
			output.WriteByte(':')
			if err := property.value.appendCanonical(output); err != nil {
				return err
			}
		}
		output.WriteByte('}')
	default:
		return errors.New("serialize JSON: unknown value kind")
	}
	return nil
}

func utf16Less(left, right string) bool {
	leftUnits := utf16.Encode([]rune(left))
	rightUnits := utf16.Encode([]rune(right))
	limit := len(leftUnits)
	if len(rightUnits) < limit {
		limit = len(rightUnits)
	}
	for index := 0; index < limit; index++ {
		if leftUnits[index] != rightUnits[index] {
			return leftUnits[index] < rightUnits[index]
		}
	}
	return len(leftUnits) < len(rightUnits)
}

func appendString(output *bytes.Buffer, text string) {
	const hexadecimal = "0123456789abcdef"
	output.WriteByte('"')
	for _, character := range text {
		switch character {
		case '"':
			output.WriteString(`\"`)
		case '\\':
			output.WriteString(`\\`)
		case '\b':
			output.WriteString(`\b`)
		case '\t':
			output.WriteString(`\t`)
		case '\n':
			output.WriteString(`\n`)
		case '\f':
			output.WriteString(`\f`)
		case '\r':
			output.WriteString(`\r`)
		default:
			if character < 0x20 {
				output.WriteString(`\u00`)
				output.WriteByte(hexadecimal[byte(character)>>4])
				output.WriteByte(hexadecimal[byte(character)&0x0f])
			} else {
				output.WriteRune(character)
			}
		}
	}
	output.WriteByte('"')
}

type jsonParser struct {
	input  []byte
	offset int
}

func (parser *jsonParser) parseValue() (jsonValue, error) {
	parser.skipWhitespace()
	if parser.offset >= len(parser.input) {
		return jsonValue{}, parser.errorf("expected JSON value")
	}
	switch parser.input[parser.offset] {
	case 'n':
		if !parser.consumeLiteral("null") {
			return jsonValue{}, parser.errorf("invalid literal")
		}
		return jsonValue{kind: kindNull}, nil
	case 'f':
		if !parser.consumeLiteral("false") {
			return jsonValue{}, parser.errorf("invalid literal")
		}
		return jsonValue{kind: kindFalse}, nil
	case 't':
		if !parser.consumeLiteral("true") {
			return jsonValue{}, parser.errorf("invalid literal")
		}
		return jsonValue{kind: kindTrue}, nil
	case '"':
		text, err := parser.parseString()
		return jsonValue{kind: kindString, text: text}, err
	case '[':
		return parser.parseArray()
	case '{':
		return parser.parseObject()
	default:
		return parser.parseNumber()
	}
}

func (parser *jsonParser) parseArray() (jsonValue, error) {
	parser.offset++
	array := make([]jsonValue, 0)
	parser.skipWhitespace()
	if parser.consumeByte(']') {
		return jsonValue{kind: kindArray, array: array}, nil
	}
	for {
		element, err := parser.parseValue()
		if err != nil {
			return jsonValue{}, err
		}
		array = append(array, element)
		parser.skipWhitespace()
		if parser.consumeByte(']') {
			return jsonValue{kind: kindArray, array: array}, nil
		}
		if !parser.consumeByte(',') {
			return jsonValue{}, parser.errorf("expected ',' or ']'")
		}
	}
}

func (parser *jsonParser) parseObject() (jsonValue, error) {
	parser.offset++
	properties := make([]jsonProperty, 0)
	seen := make(map[string]struct{})
	parser.skipWhitespace()
	if parser.consumeByte('}') {
		return jsonValue{kind: kindObject, properties: properties}, nil
	}
	for {
		parser.skipWhitespace()
		if parser.offset >= len(parser.input) || parser.input[parser.offset] != '"' {
			return jsonValue{}, parser.errorf("expected object property name")
		}
		name, err := parser.parseString()
		if err != nil {
			return jsonValue{}, err
		}
		if _, duplicate := seen[name]; duplicate {
			return jsonValue{}, parser.errorf("duplicate object property %q", name)
		}
		seen[name] = struct{}{}
		parser.skipWhitespace()
		if !parser.consumeByte(':') {
			return jsonValue{}, parser.errorf("expected ':'")
		}
		propertyValue, err := parser.parseValue()
		if err != nil {
			return jsonValue{}, err
		}
		properties = append(properties, jsonProperty{name: name, value: propertyValue})
		parser.skipWhitespace()
		if parser.consumeByte('}') {
			return jsonValue{kind: kindObject, properties: properties}, nil
		}
		if !parser.consumeByte(',') {
			return jsonValue{}, parser.errorf("expected ',' or '}'")
		}
	}
}

func (parser *jsonParser) parseString() (string, error) {
	parser.offset++
	var text bytes.Buffer
	for parser.offset < len(parser.input) {
		current := parser.input[parser.offset]
		switch current {
		case '"':
			parser.offset++
			return text.String(), nil
		case '\\':
			parser.offset++
			if err := parser.parseEscape(&text); err != nil {
				return "", err
			}
		default:
			if current < 0x20 {
				return "", parser.errorf("unescaped control character")
			}
			character, size := utf8.DecodeRune(parser.input[parser.offset:])
			if character == utf8.RuneError && size == 1 {
				return "", parser.errorf("invalid UTF-8")
			}
			text.WriteRune(character)
			parser.offset += size
		}
	}
	return "", parser.errorf("unterminated string")
}

func (parser *jsonParser) parseEscape(text *bytes.Buffer) error {
	if parser.offset >= len(parser.input) {
		return parser.errorf("unterminated escape")
	}
	escape := parser.input[parser.offset]
	parser.offset++
	switch escape {
	case '"', '\\', '/':
		text.WriteByte(escape)
	case 'b':
		text.WriteByte('\b')
	case 'f':
		text.WriteByte('\f')
	case 'n':
		text.WriteByte('\n')
	case 'r':
		text.WriteByte('\r')
	case 't':
		text.WriteByte('\t')
	case 'u':
		first, err := parser.parseHexCodeUnit()
		if err != nil {
			return err
		}
		switch {
		case first >= 0xd800 && first <= 0xdbff:
			if parser.offset+2 > len(parser.input) || parser.input[parser.offset] != '\\' || parser.input[parser.offset+1] != 'u' {
				return parser.errorf("unpaired high surrogate")
			}
			parser.offset += 2
			second, secondErr := parser.parseHexCodeUnit()
			if secondErr != nil {
				return secondErr
			}
			if second < 0xdc00 || second > 0xdfff {
				return parser.errorf("unpaired high surrogate")
			}
			text.WriteRune(utf16.DecodeRune(rune(first), rune(second)))
		case first >= 0xdc00 && first <= 0xdfff:
			return parser.errorf("unpaired low surrogate")
		default:
			text.WriteRune(rune(first))
		}
	default:
		return parser.errorf("invalid escape")
	}
	return nil
}

func (parser *jsonParser) parseHexCodeUnit() (uint16, error) {
	if parser.offset+4 > len(parser.input) {
		return 0, parser.errorf("short Unicode escape")
	}
	var value uint16
	for count := 0; count < 4; count++ {
		digit, ok := hexDigit(parser.input[parser.offset])
		if !ok {
			return 0, parser.errorf("invalid Unicode escape")
		}
		value = value<<4 | uint16(digit)
		parser.offset++
	}
	return value, nil
}

func hexDigit(value byte) (byte, bool) {
	switch {
	case value >= '0' && value <= '9':
		return value - '0', true
	case value >= 'a' && value <= 'f':
		return value - 'a' + 10, true
	case value >= 'A' && value <= 'F':
		return value - 'A' + 10, true
	default:
		return 0, false
	}
}

func (parser *jsonParser) parseNumber() (jsonValue, error) {
	start := parser.offset
	parser.consumeByte('-')
	if parser.offset >= len(parser.input) {
		return jsonValue{}, parser.errorf("invalid number")
	}
	if parser.consumeByte('0') {
		if parser.offset < len(parser.input) && parser.input[parser.offset] >= '0' && parser.input[parser.offset] <= '9' {
			return jsonValue{}, parser.errorf("leading zero in number")
		}
	} else {
		if parser.input[parser.offset] < '1' || parser.input[parser.offset] > '9' {
			return jsonValue{}, parser.errorf("invalid number")
		}
		for parser.offset < len(parser.input) && parser.input[parser.offset] >= '0' && parser.input[parser.offset] <= '9' {
			parser.offset++
		}
	}
	if parser.consumeByte('.') {
		fractionStart := parser.offset
		for parser.offset < len(parser.input) && parser.input[parser.offset] >= '0' && parser.input[parser.offset] <= '9' {
			parser.offset++
		}
		if parser.offset == fractionStart {
			return jsonValue{}, parser.errorf("empty number fraction")
		}
	}
	if parser.offset < len(parser.input) && (parser.input[parser.offset] == 'e' || parser.input[parser.offset] == 'E') {
		parser.offset++
		if parser.offset < len(parser.input) && (parser.input[parser.offset] == '+' || parser.input[parser.offset] == '-') {
			parser.offset++
		}
		exponentStart := parser.offset
		for parser.offset < len(parser.input) && parser.input[parser.offset] >= '0' && parser.input[parser.offset] <= '9' {
			parser.offset++
		}
		if parser.offset == exponentStart {
			return jsonValue{}, parser.errorf("empty number exponent")
		}
	}

	number, err := strconv.ParseFloat(string(parser.input[start:parser.offset]), 64)
	if err != nil || math.IsInf(number, 0) || math.IsNaN(number) {
		return jsonValue{}, parser.errorf("number is not finite IEEE 754 binary64")
	}
	return jsonValue{kind: kindNumber, number: number}, nil
}

func (parser *jsonParser) consumeLiteral(literal string) bool {
	if !bytes.HasPrefix(parser.input[parser.offset:], []byte(literal)) {
		return false
	}
	parser.offset += len(literal)
	return true
}

func (parser *jsonParser) consumeByte(want byte) bool {
	if parser.offset >= len(parser.input) || parser.input[parser.offset] != want {
		return false
	}
	parser.offset++
	return true
}

func (parser *jsonParser) skipWhitespace() {
	for parser.offset < len(parser.input) {
		switch parser.input[parser.offset] {
		case ' ', '\t', '\n', '\r':
			parser.offset++
		default:
			return
		}
	}
}

func (parser *jsonParser) errorf(format string, arguments ...any) error {
	return fmt.Errorf("parse JSON at byte %d: %s", parser.offset, fmt.Sprintf(format, arguments...))
}
