package appipc

import (
	"strings"

	"github.com/jackli/frank/internal/canonicaljson"
)

// MarshalJCS encodes the strict-integer JSON data model used by app IPC.
func MarshalJCS(value any) ([]byte, error) {
	return canonicaljson.Marshal(value)
}

func hasJSONOption(options, wanted string) bool {
	for _, option := range strings.Split(options, ",") {
		if option == wanted {
			return true
		}
	}
	return false
}
