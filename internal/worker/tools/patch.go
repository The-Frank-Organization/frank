package tools

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type patchKind uint8

const (
	patchAdd patchKind = iota
	patchUpdate
	patchDelete
)

type patchEdit struct {
	kind     patchKind
	path     string
	contents []byte
}

type parsedPatch struct {
	kind  patchKind
	path  string
	lines []string
}

func (backend *LocalBackend) preparePatch(encoded string) ([]patchEdit, error) {
	parsed, err := parsePatch(encoded)
	if err != nil {
		return nil, err
	}
	edits := make([]patchEdit, 0, len(parsed))
	seen := make(map[string]struct{}, len(parsed))
	for _, operation := range parsed {
		var resolved string
		switch operation.kind {
		case patchAdd:
			resolved, err = backend.resolveWritable(operation.path)
			if err == nil {
				if _, statErr := os.Lstat(resolved); statErr == nil {
					err = errors.New("add target already exists")
				} else if !errors.Is(statErr, os.ErrNotExist) {
					err = statErr
				}
			}
		case patchUpdate, patchDelete:
			resolved, err = backend.resolveExisting(operation.path)
		}
		if err != nil {
			if errors.Is(err, ErrPathOutsideWorkspace) {
				return nil, err
			}
			return nil, fmt.Errorf("%w: %s: %v", ErrMalformedPatch, operation.path, err)
		}
		if _, duplicate := seen[resolved]; duplicate {
			return nil, fmt.Errorf("%w: duplicate target %s", ErrMalformedPatch, operation.path)
		}
		seen[resolved] = struct{}{}
		edit := patchEdit{kind: operation.kind, path: resolved}
		switch operation.kind {
		case patchAdd:
			contents, addErr := addedContents(operation.lines)
			if addErr != nil {
				return nil, addErr
			}
			edit.contents = contents
		case patchUpdate:
			contents, readErr := os.ReadFile(resolved)
			if readErr != nil {
				return nil, readErr
			}
			updated, updateErr := applyHunks(string(contents), operation.lines)
			if updateErr != nil {
				return nil, updateErr
			}
			edit.contents = []byte(updated)
		case patchDelete:
			if len(operation.lines) != 0 {
				return nil, fmt.Errorf("%w: delete carries body", ErrMalformedPatch)
			}
		}
		edits = append(edits, edit)
	}
	return edits, nil
}

func parsePatch(encoded string) ([]parsedPatch, error) {
	lines := strings.Split(strings.TrimSuffix(encoded, "\n"), "\n")
	if len(lines) < 2 || lines[0] != "*** Begin Patch" || lines[len(lines)-1] != "*** End Patch" {
		return nil, fmt.Errorf("%w: missing patch envelope", ErrMalformedPatch)
	}
	var operations []parsedPatch
	for index := 1; index < len(lines)-1; {
		line := lines[index]
		var operation parsedPatch
		switch {
		case strings.HasPrefix(line, "*** Add File: "):
			operation.kind = patchAdd
			operation.path = strings.TrimPrefix(line, "*** Add File: ")
		case strings.HasPrefix(line, "*** Update File: "):
			operation.kind = patchUpdate
			operation.path = strings.TrimPrefix(line, "*** Update File: ")
		case strings.HasPrefix(line, "*** Delete File: "):
			operation.kind = patchDelete
			operation.path = strings.TrimPrefix(line, "*** Delete File: ")
		default:
			return nil, fmt.Errorf("%w: expected file header at line %d", ErrMalformedPatch, index+1)
		}
		if operation.path == "" {
			return nil, fmt.Errorf("%w: empty target", ErrMalformedPatch)
		}
		index++
		start := index
		for index < len(lines)-1 && !strings.HasPrefix(lines[index], "*** ") {
			index++
		}
		operation.lines = append([]string(nil), lines[start:index]...)
		operations = append(operations, operation)
	}
	if len(operations) == 0 {
		return nil, fmt.Errorf("%w: no operations", ErrMalformedPatch)
	}
	return operations, nil
}

func addedContents(lines []string) ([]byte, error) {
	result := make([]string, len(lines))
	for index, line := range lines {
		if !strings.HasPrefix(line, "+") {
			return nil, fmt.Errorf("%w: add line lacks + prefix", ErrMalformedPatch)
		}
		result[index] = strings.TrimPrefix(line, "+")
	}
	if len(result) == 0 {
		return nil, nil
	}
	return []byte(strings.Join(result, "\n") + "\n"), nil
}

func applyHunks(original string, lines []string) (string, error) {
	if len(lines) == 0 {
		return "", fmt.Errorf("%w: update lacks hunk", ErrMalformedPatch)
	}
	updated := original
	for index := 0; index < len(lines); {
		if lines[index] != "@@" {
			return "", fmt.Errorf("%w: malformed hunk header", ErrMalformedPatch)
		}
		index++
		var oldLines, newLines []string
		for index < len(lines) && lines[index] != "@@" {
			line := lines[index]
			if line == "" {
				return "", fmt.Errorf("%w: unprefixed hunk line", ErrMalformedPatch)
			}
			switch line[0] {
			case ' ':
				oldLines = append(oldLines, line[1:])
				newLines = append(newLines, line[1:])
			case '-':
				oldLines = append(oldLines, line[1:])
			case '+':
				newLines = append(newLines, line[1:])
			default:
				return "", fmt.Errorf("%w: invalid hunk line", ErrMalformedPatch)
			}
			index++
		}
		if len(oldLines) == 0 {
			return "", fmt.Errorf("%w: hunk has no old context", ErrMalformedPatch)
		}
		oldText := strings.Join(oldLines, "\n")
		newText := strings.Join(newLines, "\n")
		if strings.HasSuffix(updated, oldText+"\n") || strings.Contains(updated, oldText+"\n") {
			oldText += "\n"
			newText += "\n"
		}
		if count := strings.Count(updated, oldText); count != 1 {
			return "", fmt.Errorf("%w: hunk match count %d", ErrMalformedPatch, count)
		}
		updated = strings.Replace(updated, oldText, newText, 1)
	}
	return updated, nil
}
