package tools

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/The-Frank-Organization/frank/internal/worker/executor"
)

type Registry struct {
	backend Backend
}

func NewRegistry(backend Backend) *Registry { return &Registry{backend: backend} }

func (registry *Registry) Invoke(ctx context.Context, invocation executor.Invocation) (any, error) {
	if registry == nil || registry.backend == nil {
		return nil, errors.New("tools: backend is absent")
	}
	switch invocation.Identity.CanonicalToolName {
	case "read":
		var arguments struct {
			Path   string `json:"path"`
			Offset int64  `json:"offset,omitempty"`
			Limit  int64  `json:"limit,omitempty"`
		}
		if err := decodeArguments(invocation.Arguments, &arguments); err != nil || arguments.Path == "" {
			return nil, invalidArguments(err)
		}
		return registry.backend.Read(ctx, arguments.Path, arguments.Offset, arguments.Limit)
	case "write":
		var arguments struct {
			Path    string  `json:"path"`
			Content *string `json:"content"`
		}
		if err := decodeArguments(invocation.Arguments, &arguments); err != nil || arguments.Path == "" || arguments.Content == nil {
			return nil, invalidArguments(err)
		}
		if err := registry.backend.Write(ctx, arguments.Path, *arguments.Content); err != nil {
			return nil, err
		}
		return registry.backend.BoundOutput("wrote " + arguments.Path), nil
	case "edit":
		var arguments struct {
			Path       string  `json:"path"`
			OldString  *string `json:"old_string"`
			NewString  *string `json:"new_string"`
			ReplaceAll bool    `json:"replace_all,omitempty"`
		}
		if err := decodeArguments(invocation.Arguments, &arguments); err != nil || arguments.Path == "" || arguments.OldString == nil || arguments.NewString == nil {
			return nil, invalidArguments(err)
		}
		if err := registry.backend.Edit(ctx, arguments.Path, *arguments.OldString, *arguments.NewString, arguments.ReplaceAll); err != nil {
			return nil, err
		}
		return registry.backend.BoundOutput("edited " + arguments.Path), nil
	case "apply_patch":
		var arguments struct {
			Patch *string `json:"patch"`
		}
		if err := decodeArguments(invocation.Arguments, &arguments); err != nil || arguments.Patch == nil {
			return nil, invalidArguments(err)
		}
		if err := registry.backend.ApplyPatch(ctx, *arguments.Patch); err != nil {
			return nil, err
		}
		return registry.backend.BoundOutput("patch applied"), nil
	case "bash":
		var arguments struct {
			Command   string `json:"command"`
			TimeoutMS int64  `json:"timeout_ms,omitempty"`
		}
		if err := decodeArguments(invocation.Arguments, &arguments); err != nil || arguments.Command == "" || arguments.TimeoutMS < 0 {
			return nil, invalidArguments(err)
		}
		return registry.backend.Bash(ctx, arguments.Command, time.Duration(arguments.TimeoutMS)*time.Millisecond)
	default:
		return nil, fmt.Errorf("tools: unsupported local tool %q", invocation.Identity.CanonicalToolName)
	}
}

func decodeArguments(encoded []byte, target any) error {
	decoder := json.NewDecoder(bytes.NewReader(encoded))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return err
	}
	var trailing any
	if err := decoder.Decode(&trailing); !errors.Is(err, io.EOF) {
		if err == nil {
			return errors.New("trailing JSON value")
		}
		return err
	}
	return nil
}

func invalidArguments(cause error) error {
	if cause == nil {
		return ErrInvalidArguments
	}
	return fmt.Errorf("%w: %v", ErrInvalidArguments, cause)
}
