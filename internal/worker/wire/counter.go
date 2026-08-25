package wire

import (
	"errors"
	"strconv"
)

var errMalformedCounter = errors.New("malformed counter")

// ParseCounter accepts only the canonical decimal uint64 string form used by
// CTRL-W and DATA-P frames.
func ParseCounter(text string) (uint64, error) {
	if text == "" || (len(text) > 1 && text[0] == '0') {
		return 0, errMalformedCounter
	}
	for index := 0; index < len(text); index++ {
		if text[index] < '0' || text[index] > '9' {
			return 0, errMalformedCounter
		}
	}
	value, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		return 0, errMalformedCounter
	}
	return value, nil
}

// FormatCounter emits the unique canonical decimal representation of value.
func FormatCounter(value uint64) string {
	return strconv.FormatUint(value, 10)
}
