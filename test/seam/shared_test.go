//go:build seam

package seam_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jackli/frank/internal/appipc"
	connectorjcs "github.com/jackli/frank/internal/connector/jcs"
	workerjcs "github.com/jackli/frank/internal/worker/jcs"
)

// TestCT_D01 binds CT-D01; selectors r9 D01 and extraction seam; flags MISMATCH, EXTRACT, latent.
func TestCT_D01(t *testing.T) {
	integers := []string{
		"0", "1", "-1", "9007199254740993", "-9007199254740993",
		"12345678901234567890123456789012345678901234567890",
	}
	for _, integer := range integers {
		input := []byte(`{"n":` + integer + `}`)
		a, ea := appCanonical(input)
		c, ec := connectorjcs.Canonicalize(input)
		w, ew := workerjcs.Canonicalize(input)
		digest, ed := workerjcs.Digest(input)
		sum := sha256.Sum256(input)
		wantDigest := hex.EncodeToString(sum[:])
		contract(t, ea == nil && ec == nil && ew == nil && ed == nil && canonicalEqual(input, a, c, w) && digest == wantDigest,
			explain("integer was not carried verbatim through encode/digest: integer=%s app=%s/%v connector=%s/%v worker=%s/%v digest=%s/%v", integer, a, ea, c, ec, w, ew, digest, ed))
	}

	for _, number := range []string{"-0", "1.5", "1e2", "1E2"} {
		input := []byte(`{"n":` + number + `}`)
		_, ea := appCanonical(input)
		_, ec := connectorjcs.Canonicalize(input)
		_, ew := workerjcs.Canonicalize(input)
		_, ed := workerjcs.Digest(input)
		contract(t, typedNumberRefusal(ea, number) && typedNumberRefusal(ec, number) && typedNumberRefusal(ew, number) && typedNumberRefusal(ed, number),
			explain("number refusal was not typed everywhere: number=%s app=%v connector=%v worker=%v digest=%v", number, ea, ec, ew, ed))
	}
	for _, value := range []any{float32(1), float64(1)} {
		_, err := appipc.MarshalJCS(value)
		contract(t, typedNumberRefusal(err, ""), explain("Go float %T did not receive a typed refusal: %v", value, err))
	}
}

// TestCT_D02 binds CT-D02; selectors r9 D02 and extraction seam; flags MISMATCH, EXTRACT, latent.
func TestCT_D02(t *testing.T) {
	for _, vector := range sharedCanonicalVectors {
		a, ea := appCanonical(vector.input)
		c, ec := connectorjcs.Canonicalize(vector.input)
		w, ew := workerjcs.Canonicalize(vector.input)
		digest, ed := workerjcs.Digest(vector.input)
		sum := sha256.Sum256(vector.want)
		contract(t, ea == nil && ec == nil && ew == nil && ed == nil && canonicalEqual(vector.want, a, c, w) && digest == hex.EncodeToString(sum[:]),
			explain("shared canonical vector %s disagreed: app=%s/%v connector=%s/%v worker=%s/%v digest=%s/%v", vector.name, a, ea, c, ec, w, ew, digest, ed))
	}
}

type canonicalVector struct {
	name        string
	input, want []byte
}

var sharedCanonicalVectors = []canonicalVector{
	{name: "null", input: []byte(`null`), want: []byte(`null`)},
	{name: "booleans", input: []byte(`[true,false]`), want: []byte(`[true,false]`)},
	{name: "integer zero", input: []byte(`0`), want: []byte(`0`)},
	{name: "negative integer", input: []byte(`-17`), want: []byte(`-17`)},
	{name: "above binary64 exact range", input: []byte(`9007199254740993`), want: []byte(`9007199254740993`)},
	{name: "arbitrary precision", input: []byte(`12345678901234567890123456789012345678901234567890`), want: []byte(`12345678901234567890123456789012345678901234567890`)},
	{name: "empty containers", input: []byte(` { "z" : [ ], "a" : { } } `), want: []byte(`{"a":{},"z":[]}`)},
	{name: "nested order", input: []byte(`{"z":[{"b":1,"a":2}],"a":{"d":4,"c":3}}`), want: []byte(`{"a":{"c":3,"d":4},"z":[{"a":2,"b":1}]}`)},
	{name: "UTF-16 member order", input: []byte(`{"\ue000":1,"\ud83d\ude00":2}`), want: []byte(`{"😀":2,"":1}`)},
	{name: "minimal escapes", input: []byte(`{"escaped":"\u003c\u003e\u0026\/","control":"\u000f\n"}`), want: []byte(`{"control":"\u000f\n","escaped":"<>&/"}`)},
	{name: "Unicode preserved", input: []byte(`{"composed":"é","decomposed":"e\u0301"}`), want: []byte("{\"composed\":\"é\",\"decomposed\":\"é\"}")},
}

type numberRefusal interface {
	NumberToken() string
}

func typedNumberRefusal(err error, token string) bool {
	var refusal numberRefusal
	return err != nil && errors.As(err, &refusal) && (token == "" || refusal.NumberToken() == token)
}

// TestCT_D04 binds CT-D04; selectors r9 D04 and A-M10-LIMITS; flags MISMATCH, CP, latent.
func TestCT_D04(t *testing.T) {
	peerReferences := reducedLimitPeerReferences(t)
	wireMembers := registeredWireLimitMembers(t)
	contract(t, len(peerReferences) == 0 && len(wireMembers) == 0,
		explain("test-reduced limits escape appipc: peer_references=%v registered_wire_members=%v", peerReferences, wireMembers))
}

// TestCT_D05 binds CT-D05; selectors r9 D05 and shared frame extraction; flags MISMATCH, CP, benign.
func TestCT_D05(t *testing.T) {
	var output bytes.Buffer
	contract(t, appipc.WriteFrame(&output, nil) != nil, "appipc.WriteFrame accepts an empty payload")
}

func reducedLimitPeerReferences(t *testing.T) []string {
	t.Helper()
	root := repoRoot(t)
	var references []string
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			relative, _ := filepath.Rel(root, path)
			if relative == ".git" || relative == "cmd" || relative == "internal/appctl" || relative == "internal/appipc" || relative == "test/seam" {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		firstLine := string(raw)
		if newline := bytes.IndexByte(raw, '\n'); newline >= 0 {
			firstLine = string(raw[:newline])
		}
		found := strings.HasPrefix(firstLine, "//go:build ") && strings.Contains(firstLine, "frank_test_reduced_limits")
		file, err := parser.ParseFile(token.NewFileSet(), path, raw, 0)
		if err != nil {
			return err
		}
		ast.Inspect(file, func(node ast.Node) bool {
			if identifier, ok := node.(*ast.Ident); ok && identifier.Name == "ReducedLimitsArtifact" {
				found = true
			}
			return true
		})
		if found {
			relative, _ := filepath.Rel(root, path)
			references = append(references, filepath.ToSlash(relative))
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return references
}

func registeredWireLimitMembers(t *testing.T) []string {
	t.Helper()
	paths, err := filepath.Glob(filepath.Join(repoRoot(t), "internal/appipc/msgs_*.go"))
	if err != nil {
		t.Fatal(err)
	}
	var members []string
	for _, path := range paths {
		file, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		for _, declaration := range file.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, specification := range general.Specs {
				typeSpec, ok := specification.(*ast.TypeSpec)
				if !ok {
					continue
				}
				structure, ok := typeSpec.Type.(*ast.StructType)
				if !ok || !strings.HasSuffix(typeSpec.Name.Name, "Body") {
					continue
				}
				for _, field := range structure.Fields.List {
					for _, name := range field.Names {
						lower := strings.ToLower(name.Name)
						if strings.Contains(lower, "limit") || strings.HasPrefix(lower, "max") {
							members = append(members, typeSpec.Name.Name+"."+name.Name)
						}
					}
				}
			}
		}
	}
	return members
}
