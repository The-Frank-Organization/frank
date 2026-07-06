package egress

import "strings"

type Classifier func(Item) []Finding

type Rules struct {
	Classifiers []Classifier
}

func DefaultRules() Rules {
	return Rules{Classifiers: []Classifier{defaultClassifier}}
}

func (r Rules) withDefaults() Rules {
	if len(r.Classifiers) == 0 {
		return DefaultRules()
	}
	return r
}

func defaultClassifier(item Item) []Finding {
	value := strings.ToLower(item.Field.Value)
	var findings []Finding
	if containsAny(value, []string{
		"api_key",
		"secret",
		"password=",
		"authorization:",
		"auth_url=",
		"://user:pass@",
	}) {
		findings = append(findings, Finding{Field: item.Field.Name, Class: ClassSafety})
	}
	if containsAny(value, []string{
		"gpt-",
		"claude-",
		"llama",
		"model_name=",
	}) {
		findings = append(findings, Finding{Field: item.Field.Name, Class: ClassConfidentiality})
	}
	return findings
}

func containsAny(value string, needles []string) bool {
	for _, needle := range needles {
		if strings.Contains(value, needle) {
			return true
		}
	}
	return false
}
