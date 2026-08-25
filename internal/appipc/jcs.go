package appipc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode/utf16"
	"unicode/utf8"
)

// MarshalJCS encodes the JSON data model used by the app IPC protocol. Trust
// paths deliberately exclude floating-point values.
func MarshalJCS(value any) ([]byte, error) {
	var out bytes.Buffer
	if err := appendJCS(&out, reflect.ValueOf(value)); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func appendJCS(out *bytes.Buffer, value reflect.Value) error {
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
		number := value.Interface().(json.Number).String()
		if !isJSONInteger(number) {
			return fmt.Errorf("appipc: floating-point values are forbidden on JCS trust paths")
		}
		out.WriteString(number)
		return nil
	}

	switch value.Kind() {
	case reflect.Bool:
		out.WriteString(strconv.FormatBool(value.Bool()))
	case reflect.String:
		return appendJCSString(out, value.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		out.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		out.WriteString(strconv.FormatUint(value.Uint(), 10))
	case reflect.Float32, reflect.Float64:
		return fmt.Errorf("appipc: floating-point values are forbidden on JCS trust paths")
	case reflect.Slice:
		if value.IsNil() {
			out.WriteString("null")
			return nil
		}
		return appendJCSArray(out, value)
	case reflect.Array:
		return appendJCSArray(out, value)
	case reflect.Map:
		return appendJCSMap(out, value)
	case reflect.Struct:
		return appendJCSStruct(out, value)
	default:
		return fmt.Errorf("appipc: unsupported JCS value kind %s", value.Kind())
	}
	return nil
}

func isJSONInteger(value string) bool {
	if value == "" {
		return false
	}
	start := 0
	if value[0] == '-' {
		start = 1
		if len(value) == 1 {
			return false
		}
	}
	if value[start] == '0' && len(value)-start > 1 {
		return false
	}
	for i := start; i < len(value); i++ {
		if value[i] < '0' || value[i] > '9' {
			return false
		}
	}
	return true
}

func appendJCSArray(out *bytes.Buffer, value reflect.Value) error {
	out.WriteByte('[')
	for i := 0; i < value.Len(); i++ {
		if i > 0 {
			out.WriteByte(',')
		}
		if err := appendJCS(out, value.Index(i)); err != nil {
			return err
		}
	}
	out.WriteByte(']')
	return nil
}

type jcsMember struct {
	name  string
	value reflect.Value
}

func appendJCSMap(out *bytes.Buffer, value reflect.Value) error {
	if value.Type().Key().Kind() != reflect.String {
		return fmt.Errorf("appipc: JCS object keys must be strings")
	}
	if value.IsNil() {
		out.WriteString("null")
		return nil
	}
	members := make([]jcsMember, 0, value.Len())
	iter := value.MapRange()
	for iter.Next() {
		key := iter.Key().String()
		if !utf8.ValidString(key) {
			return fmt.Errorf("appipc: JCS object key is not valid UTF-8")
		}
		members = append(members, jcsMember{name: key, value: iter.Value()})
	}
	return appendJCSObject(out, members)
}

func appendJCSStruct(out *bytes.Buffer, value reflect.Value) error {
	typeOf := value.Type()
	members := make([]jcsMember, 0, value.NumField())
	seen := make(map[string]struct{}, value.NumField())
	for i := 0; i < value.NumField(); i++ {
		field := typeOf.Field(i)
		if field.PkgPath != "" { // unexported
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
			return fmt.Errorf("appipc: JCS struct field name is not valid UTF-8")
		}
		fieldValue := value.Field(i)
		if hasJSONOption(options, "omitempty") && fieldValue.IsZero() {
			continue
		}
		if _, duplicate := seen[name]; duplicate {
			return fmt.Errorf("appipc: duplicate JCS object member %q", name)
		}
		seen[name] = struct{}{}
		members = append(members, jcsMember{name: name, value: fieldValue})
	}
	return appendJCSObject(out, members)
}

func hasJSONOption(options, wanted string) bool {
	for _, option := range strings.Split(options, ",") {
		if option == wanted {
			return true
		}
	}
	return false
}

func appendJCSObject(out *bytes.Buffer, members []jcsMember) error {
	sort.Slice(members, func(i, j int) bool {
		return lessUTF16(members[i].name, members[j].name)
	})
	out.WriteByte('{')
	for i, member := range members {
		if i > 0 {
			out.WriteByte(',')
		}
		if err := appendJCSString(out, member.name); err != nil {
			return err
		}
		out.WriteByte(':')
		if err := appendJCS(out, member.value); err != nil {
			return err
		}
	}
	out.WriteByte('}')
	return nil
}

func lessUTF16(left, right string) bool {
	a := utf16.Encode([]rune(left))
	b := utf16.Encode([]rune(right))
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func appendJCSString(out *bytes.Buffer, value string) error {
	if !utf8.ValidString(value) {
		return fmt.Errorf("appipc: JCS string is not valid UTF-8")
	}
	out.WriteByte('"')
	for _, r := range value {
		switch r {
		case '"', '\\':
			out.WriteByte('\\')
			out.WriteRune(r)
		case '\b':
			out.WriteString(`\b`)
		case '\t':
			out.WriteString(`\t`)
		case '\n':
			out.WriteString(`\n`)
		case '\f':
			out.WriteString(`\f`)
		case '\r':
			out.WriteString(`\r`)
		default:
			if r < 0x20 {
				out.WriteString(`\u00`)
				const hex = "0123456789abcdef"
				out.WriteByte(hex[byte(r)>>4])
				out.WriteByte(hex[byte(r)&0x0f])
			} else {
				out.WriteRune(r)
			}
		}
	}
	out.WriteByte('"')
	return nil
}
