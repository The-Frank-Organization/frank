package manifest

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

const PathMaxM10 = 4096

type WorkspaceRootReason string

const (
	RootNotAbsolute    WorkspaceRootReason = "not_absolute"
	RootFilesystemRoot WorkspaceRootReason = "filesystem_root"
	RootUnresolvable   WorkspaceRootReason = "unresolvable"
	RootOutOfGrammar   WorkspaceRootReason = "out_of_grammar"
	RootTooLong        WorkspaceRootReason = "too_long"
)

type WorkspaceRootError struct {
	Reason WorkspaceRootReason
}

func (err *WorkspaceRootError) Error() string {
	return "workspace_root_invalid:" + string(err.Reason)
}

type WorkspaceRoot struct {
	Path string
	ID   string
}

func ResolveWorkspaceRoot(input string) (WorkspaceRoot, error) {
	return resolveWorkspaceRoot(input, filepath.EvalSymlinks)
}

func resolveWorkspaceRoot(input string, realpath func(string) (string, error)) (WorkspaceRoot, error) {
	if input == "" || !strings.HasPrefix(input, "/") || !filepath.IsAbs(input) {
		return WorkspaceRoot{}, &WorkspaceRootError{Reason: RootNotAbsolute}
	}
	resolved, err := realpath(input)
	if err != nil || resolved == "" || !filepath.IsAbs(resolved) {
		return WorkspaceRoot{}, &WorkspaceRootError{Reason: RootUnresolvable}
	}
	if resolved == string(filepath.Separator) {
		return WorkspaceRoot{}, &WorkspaceRootError{Reason: RootFilesystemRoot}
	}
	normalized := norm.NFC.String(resolved)
	if !validRootPath(normalized) {
		return WorkspaceRoot{}, &WorkspaceRootError{Reason: RootOutOfGrammar}
	}
	if len([]byte(normalized)) > PathMaxM10 {
		return WorkspaceRoot{}, &WorkspaceRootError{Reason: RootTooLong}
	}
	digest := sha256.Sum256([]byte(normalized))
	return WorkspaceRoot{Path: normalized, ID: hex.EncodeToString(digest[:])}, nil
}

func validRootPath(path string) bool {
	if !utf8.ValidString(path) || !strings.HasPrefix(path, "/") || strings.HasSuffix(path, "/") {
		return false
	}
	for _, r := range path {
		if r == 0 || unicode.IsControl(r) {
			return false
		}
	}
	return true
}

func IsWorkspaceRootError(err error, reason WorkspaceRootReason) bool {
	var typed *WorkspaceRootError
	return errors.As(err, &typed) && typed.Reason == reason
}
