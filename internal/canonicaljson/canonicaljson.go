// Package canonicaljson provides Frank's single canonical JSON primitive.
// Numbers are restricted to strict integer tokens and emitted verbatim.
package canonicaljson

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

type NumberError struct{ token string }

func (err *NumberError) Error() string {
	if err.token == "" {
		return "canonicaljson: Go floating-point values are forbidden"
	}
	return fmt.Sprintf("canonicaljson: number %q is not a strict integer token", err.token)
}
func (err *NumberError) NumberToken() string { return err.token }

func Marshal(input any) ([]byte, error) {
	var out bytes.Buffer
	if err := appendReflected(&out, reflect.ValueOf(input)); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func Canonicalize(input []byte) ([]byte, error) {
	if !utf8.Valid(input) {
		return nil, fmt.Errorf("canonicaljson: input is not valid UTF-8")
	}
	if err := validateStringTokens(input); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(bytes.NewReader(input))
	decoder.UseNumber()
	value, err := parseValue(decoder)
	if err != nil {
		return nil, err
	}
	if _, err := decoder.Token(); err != io.EOF {
		if err == nil {
			return nil, fmt.Errorf("canonicaljson: trailing JSON value")
		}
		return nil, fmt.Errorf("canonicaljson: trailing data: %w", err)
	}
	return appendParsed(nil, value)
}

func IsCanonical(input []byte) bool {
	canonical, err := Canonicalize(input)
	return err == nil && bytes.Equal(input, canonical)
}

func Digest(input []byte) (string, error) {
	canonical, err := Canonicalize(input)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(canonical)
	return hex.EncodeToString(sum[:]), nil
}

func strictInteger(token string) bool {
	if token == "" || token == "-0" {
		return false
	}
	start := 0
	if token[0] == '-' {
		start = 1
		if len(token) == 1 {
			return false
		}
	}
	if token[start] == '0' {
		return len(token)-start == 1
	}
	if token[start] < '1' || token[start] > '9' {
		return false
	}
	for i := start + 1; i < len(token); i++ {
		if token[i] < '0' || token[i] > '9' {
			return false
		}
	}
	return true
}

func appendReflected(out *bytes.Buffer, value reflect.Value) error {
	if !value.IsValid() {
		out.WriteString("null")
		return nil
	}
	for value.Kind() == reflect.Interface || value.Kind() == reflect.Pointer {
		if value.IsNil() {
			out.WriteString("null")
			return nil
		}
		value = value.Elem()
	}
	if value.Type() == reflect.TypeOf(json.Number("")) {
		token := value.Interface().(json.Number).String()
		if !strictInteger(token) {
			return &NumberError{token: token}
		}
		out.WriteString(token)
		return nil
	}
	switch value.Kind() {
	case reflect.Bool:
		out.WriteString(strconv.FormatBool(value.Bool()))
	case reflect.String:
		return appendStringBuffer(out, value.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		out.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		out.WriteString(strconv.FormatUint(value.Uint(), 10))
	case reflect.Float32, reflect.Float64:
		return &NumberError{}
	case reflect.Slice:
		if value.IsNil() {
			out.WriteString("null")
			return nil
		}
		return appendReflectedArray(out, value)
	case reflect.Array:
		return appendReflectedArray(out, value)
	case reflect.Map:
		return appendReflectedMap(out, value)
	case reflect.Struct:
		return appendReflectedStruct(out, value)
	default:
		return fmt.Errorf("canonicaljson: unsupported value kind %s", value.Kind())
	}
	return nil
}

func appendReflectedArray(out *bytes.Buffer, value reflect.Value) error {
	out.WriteByte('[')
	for i := 0; i < value.Len(); i++ {
		if i > 0 {
			out.WriteByte(',')
		}
		if err := appendReflected(out, value.Index(i)); err != nil {
			return err
		}
	}
	out.WriteByte(']')
	return nil
}

type reflectedMember struct {
	name  string
	value reflect.Value
}

func appendReflectedMap(out *bytes.Buffer, value reflect.Value) error {
	if value.Type().Key().Kind() != reflect.String {
		return fmt.Errorf("canonicaljson: object keys must be strings")
	}
	if value.IsNil() {
		out.WriteString("null")
		return nil
	}
	members := make([]reflectedMember, 0, value.Len())
	iterator := value.MapRange()
	for iterator.Next() {
		name := iterator.Key().String()
		if !utf8.ValidString(name) {
			return fmt.Errorf("canonicaljson: object key is not valid UTF-8")
		}
		members = append(members, reflectedMember{name: name, value: iterator.Value()})
	}
	return appendReflectedObject(out, members)
}

func appendReflectedStruct(out *bytes.Buffer, value reflect.Value) error {
	typeOf := value.Type()
	members := make([]reflectedMember, 0, value.NumField())
	seen := make(map[string]struct{}, value.NumField())
	for i := 0; i < value.NumField(); i++ {
		field := typeOf.Field(i)
		if field.PkgPath != "" {
			continue
		}
		name, options, _ := strings.Cut(field.Tag.Get("json"), ",")
		if name == "-" {
			continue
		}
		if name == "" {
			name = field.Name
		}
		if !utf8.ValidString(name) {
			return fmt.Errorf("canonicaljson: struct field name is not valid UTF-8")
		}
		fieldValue := value.Field(i)
		if hasOption(options, "omitempty") && fieldValue.IsZero() {
			continue
		}
		if _, duplicate := seen[name]; duplicate {
			return fmt.Errorf("canonicaljson: duplicate object member %q", name)
		}
		seen[name] = struct{}{}
		members = append(members, reflectedMember{name: name, value: fieldValue})
	}
	return appendReflectedObject(out, members)
}

func hasOption(options, wanted string) bool {
	for _, option := range strings.Split(options, ",") {
		if option == wanted {
			return true
		}
	}
	return false
}

func appendReflectedObject(out *bytes.Buffer, members []reflectedMember) error {
	sort.Slice(members, func(i, j int) bool { return utf16Less(members[i].name, members[j].name) })
	out.WriteByte('{')
	for i, member := range members {
		if i > 0 {
			out.WriteByte(',')
		}
		if err := appendStringBuffer(out, member.name); err != nil {
			return err
		}
		out.WriteByte(':')
		if err := appendReflected(out, member.value); err != nil {
			return err
		}
	}
	out.WriteByte('}')
	return nil
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

type parsedValue struct {
	kind    valueKind
	boolean bool
	number  string
	text    string
	array   []parsedValue
	object  []parsedMember
}
type parsedMember struct {
	name  string
	value parsedValue
}

func parseValue(decoder *json.Decoder) (parsedValue, error) {
	token, err := decoder.Token()
	if err != nil {
		return parsedValue{}, fmt.Errorf("canonicaljson: parse JSON: %w", err)
	}
	switch token := token.(type) {
	case nil:
		return parsedValue{kind: nullKind}, nil
	case bool:
		return parsedValue{kind: boolKind, boolean: token}, nil
	case string:
		return parsedValue{kind: stringKind, text: token}, nil
	case json.Number:
		return parsedValue{kind: numberKind, number: token.String()}, nil
	case json.Delim:
		switch token {
		case '[':
			var items []parsedValue
			for decoder.More() {
				item, err := parseValue(decoder)
				if err != nil {
					return parsedValue{}, err
				}
				items = append(items, item)
			}
			end, err := decoder.Token()
			if err != nil || end != json.Delim(']') {
				return parsedValue{}, fmt.Errorf("canonicaljson: malformed array")
			}
			return parsedValue{kind: arrayKind, array: items}, nil
		case '{':
			seen := make(map[string]struct{})
			var members []parsedMember
			for decoder.More() {
				nameToken, err := decoder.Token()
				if err != nil {
					return parsedValue{}, fmt.Errorf("canonicaljson: read object name: %w", err)
				}
				name, ok := nameToken.(string)
				if !ok {
					return parsedValue{}, fmt.Errorf("canonicaljson: object name is not a string")
				}
				if _, duplicate := seen[name]; duplicate {
					return parsedValue{}, fmt.Errorf("canonicaljson: duplicate object name %q", name)
				}
				seen[name] = struct{}{}
				memberValue, err := parseValue(decoder)
				if err != nil {
					return parsedValue{}, err
				}
				members = append(members, parsedMember{name: name, value: memberValue})
			}
			end, err := decoder.Token()
			if err != nil || end != json.Delim('}') {
				return parsedValue{}, fmt.Errorf("canonicaljson: malformed object")
			}
			return parsedValue{kind: objectKind, object: members}, nil
		}
	}
	return parsedValue{}, fmt.Errorf("canonicaljson: unsupported JSON token %T", token)
}

func appendParsed(destination []byte, value parsedValue) ([]byte, error) {
	switch value.kind {
	case nullKind:
		return append(destination, "null"...), nil
	case boolKind:
		return strconv.AppendBool(destination, value.boolean), nil
	case numberKind:
		if !strictInteger(value.number) {
			return nil, &NumberError{token: value.number}
		}
		return append(destination, value.number...), nil
	case stringKind:
		return appendString(destination, value.text), nil
	case arrayKind:
		destination = append(destination, '[')
		for i, item := range value.array {
			if i > 0 {
				destination = append(destination, ',')
			}
			var err error
			destination, err = appendParsed(destination, item)
			if err != nil {
				return nil, err
			}
		}
		return append(destination, ']'), nil
	case objectKind:
		sort.Slice(value.object, func(i, j int) bool { return utf16Less(value.object[i].name, value.object[j].name) })
		destination = append(destination, '{')
		for i, member := range value.object {
			if i > 0 {
				destination = append(destination, ',')
			}
			destination = appendString(destination, member.name)
			destination = append(destination, ':')
			var err error
			destination, err = appendParsed(destination, member.value)
			if err != nil {
				return nil, err
			}
		}
		return append(destination, '}'), nil
	}
	return nil, fmt.Errorf("canonicaljson: invalid internal value kind %d", value.kind)
}

func appendStringBuffer(out *bytes.Buffer, value string) error {
	if !utf8.ValidString(value) {
		return fmt.Errorf("canonicaljson: string is not valid UTF-8")
	}
	out.Write(appendString(nil, value))
	return nil
}

func appendString(destination []byte, value string) []byte {
	const hexadecimal = "0123456789abcdef"
	destination = append(destination, '"')
	for i := 0; i < len(value); i++ {
		character := value[i]
		switch character {
		case '"', '\\':
			destination = append(destination, '\\', character)
		case '\b':
			destination = append(destination, '\\', 'b')
		case '\t':
			destination = append(destination, '\\', 't')
		case '\n':
			destination = append(destination, '\\', 'n')
		case '\f':
			destination = append(destination, '\\', 'f')
		case '\r':
			destination = append(destination, '\\', 'r')
		default:
			if character < 0x20 {
				destination = append(destination, '\\', 'u', '0', '0', hexadecimal[character>>4], hexadecimal[character&0x0f])
			} else {
				destination = append(destination, character)
			}
		}
	}
	return append(destination, '"')
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
		character := input[i]
		if !inString {
			if character == '"' {
				inString = true
			}
			continue
		}
		switch character {
		case '"':
			inString = false
		case '\\':
			if i+1 >= len(input) {
				return fmt.Errorf("canonicaljson: unterminated string escape")
			}
			escape := input[i+1]
			if escape != 'u' {
				if !bytes.ContainsRune([]byte(`\"\\/bfnrt`), rune(escape)) {
					return fmt.Errorf("canonicaljson: invalid string escape \\%c", escape)
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
					return fmt.Errorf("canonicaljson: high surrogate is not followed by a low surrogate")
				}
				second, err := parseHex16(input, i+8)
				if err != nil {
					return err
				}
				if second < 0xdc00 || second > 0xdfff {
					return fmt.Errorf("canonicaljson: high surrogate is not followed by a low surrogate")
				}
				i += 11
			case first >= 0xdc00 && first <= 0xdfff:
				return fmt.Errorf("canonicaljson: lone low surrogate")
			default:
				i += 5
			}
		default:
			if character < 0x20 {
				return fmt.Errorf("canonicaljson: unescaped control byte in string")
			}
		}
	}
	if inString {
		return fmt.Errorf("canonicaljson: unterminated string")
	}
	return nil
}

func parseHex16(input []byte, start int) (uint16, error) {
	if start+4 > len(input) {
		return 0, fmt.Errorf("canonicaljson: short Unicode escape")
	}
	value, err := strconv.ParseUint(string(input[start:start+4]), 16, 16)
	if err != nil {
		return 0, fmt.Errorf("canonicaljson: invalid Unicode escape: %w", err)
	}
	return uint16(value), nil
}
