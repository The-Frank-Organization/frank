package appipc

import (
	"errors"
	"fmt"
	"strconv"
)

const StoreCounterWidth = 20

var ErrInvalidCounter = errors.New("appipc: invalid canonical uint64 counter")

// ParseCounter validates the wire grammar ^(0|[1-9][0-9]*)$ and the uint64
// bound before returning the numeric comparison value.
func ParseCounter(wire string) (uint64, error) {
	if wire == "" || (len(wire) > 1 && wire[0] == '0') {
		return 0, ErrInvalidCounter
	}
	for i := 0; i < len(wire); i++ {
		if wire[i] < '0' || wire[i] > '9' {
			return 0, ErrInvalidCounter
		}
	}
	value, err := strconv.ParseUint(wire, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %q", ErrInvalidCounter, wire)
	}
	return value, nil
}

func FormatCounter(value uint64) string {
	return strconv.FormatUint(value, 10)
}

// PadCounter converts a canonical wire counter into the fixed-width sortable
// store representation.
func PadCounter(wire string) (string, error) {
	value, err := ParseCounter(wire)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%0*d", StoreCounterWidth, value), nil
}

// UnpadCounter validates and converts the fixed-width store representation to
// the canonical wire representation.
func UnpadCounter(stored string) (string, error) {
	if len(stored) != StoreCounterWidth {
		return "", ErrInvalidCounter
	}
	for i := 0; i < len(stored); i++ {
		if stored[i] < '0' || stored[i] > '9' {
			return "", ErrInvalidCounter
		}
	}
	value, err := strconv.ParseUint(stored, 10, 64)
	if err != nil {
		return "", fmt.Errorf("%w: %q", ErrInvalidCounter, stored)
	}
	return FormatCounter(value), nil
}
