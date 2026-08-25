package frame

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// Counter is the full uint64 trust-bearing counter domain carried over JSON
// as an unpadded canonical decimal string.
type Counter uint64

// ParseCounter parses the canonical counter grammar ^(0|[1-9][0-9]*)$.
func ParseCounter(text string) (Counter, error) {
	if text == "" || (len(text) > 1 && text[0] == '0') {
		return 0, fmt.Errorf("frame: malformed counter %q", text)
	}
	for i := range text {
		if text[i] < '0' || text[i] > '9' {
			return 0, fmt.Errorf("frame: malformed counter %q", text)
		}
	}
	value, err := strconv.ParseUint(text, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("frame: malformed counter %q: %w", text, err)
	}
	return Counter(value), nil
}

func (c Counter) String() string {
	return strconv.FormatUint(uint64(c), 10)
}

func (c Counter) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c *Counter) UnmarshalJSON(data []byte) error {
	if c == nil {
		return fmt.Errorf("frame: nil counter receiver")
	}
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("frame: counter must be a JSON string")
	}
	var text string
	decoder := json.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&text); err != nil {
		return fmt.Errorf("frame: decode counter: %w", err)
	}
	parsed, err := ParseCounter(text)
	if err != nil {
		return err
	}
	*c = parsed
	return nil
}
