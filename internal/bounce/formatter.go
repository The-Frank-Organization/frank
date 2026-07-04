package bounce

import (
	"fmt"
	"strings"

	"github.com/jackli/frank/internal/fieldspec"
	"github.com/jackli/frank/internal/lineage"
)

func Format(values ...any) string {
	var parts []string
	for _, value := range values {
		switch v := value.(type) {
		case fieldspec.Violation:
			parts = append(parts, fmt.Sprintf("%s:%s", v.Field, v.Class))
		case lineage.Bounce:
			parts = append(parts, fmt.Sprintf("%s:%s", v.Edge, v.Kind))
		case *lineage.Bounce:
			if v != nil {
				parts = append(parts, fmt.Sprintf("%s:%s", v.Edge, v.Kind))
			}
		}
	}
	return strings.Join(parts, "; ")
}
