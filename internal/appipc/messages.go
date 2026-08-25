package appipc

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func NewProtocolRegistry() (*Registry, error) {
	registry := NewRegistry()
	for _, register := range []func(*Registry) error{registerCtrlW, registerCtrlC, registerBroker} {
		if err := register(registry); err != nil {
			return nil, err
		}
	}
	return registry, nil
}

func registerBody[T any](registry *Registry, channel Channel, messageType string, reply bool, evolution FamilyEvolution, enums map[string][]string, validate func(*T) error) error {
	typeOf := reflect.TypeOf((*T)(nil)).Elem()
	if typeOf.Kind() != reflect.Struct {
		return fmt.Errorf("appipc: body type for %s/%s is not a struct", channel, messageType)
	}
	fields := make([]string, 0, typeOf.NumField())
	required := make([]string, 0, typeOf.NumField())
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		name, options, _ := strings.Cut(field.Tag.Get("json"), ",")
		if name == "-" {
			continue
		}
		if name == "" {
			name = field.Name
		}
		fields = append(fields, name)
		if !hasJSONOption(options, "omitempty") {
			required = append(required, name)
		}
	}
	var bodyValidator func(map[string]any) error
	if validate != nil {
		bodyValidator = func(body map[string]any) error {
			encoded, err := json.Marshal(body)
			if err != nil {
				return err
			}
			var typed T
			if err := json.Unmarshal(encoded, &typed); err != nil {
				return err
			}
			return validate(&typed)
		}
	}
	return registry.Register(messageType, MessageSpec{
		Channel:            channel,
		Reply:              reply,
		Evolution:          evolution,
		BodyFields:         fields,
		RequiredBodyFields: required,
		Enums:              enums,
		NewBody:            func() any { return new(T) },
		ValidateBody:       bodyValidator,
	})
}

func requiredStrings(values ...string) error {
	for _, value := range values {
		if value == "" {
			return fmt.Errorf("required string is empty")
		}
	}
	return nil
}

func validateDigest(value string) error {
	if len(value) != 64 {
		return fmt.Errorf("digest must be 64 lowercase hex characters")
	}
	for i := range value {
		if (value[i] < '0' || value[i] > '9') && (value[i] < 'a' || value[i] > 'f') {
			return fmt.Errorf("digest must be 64 lowercase hex characters")
		}
	}
	return nil
}

func validateHex32(value string) error {
	if len(value) != 32 {
		return fmt.Errorf("value must be 32 lowercase hex characters")
	}
	for i := range value {
		if (value[i] < '0' || value[i] > '9') && (value[i] < 'a' || value[i] > 'f') {
			return fmt.Errorf("value must be 32 lowercase hex characters")
		}
	}
	return nil
}

func validateCounterFields(values ...string) error {
	for _, value := range values {
		if _, err := ParseCounter(value); err != nil {
			return err
		}
	}
	return nil
}

type EmptyBody struct{}
