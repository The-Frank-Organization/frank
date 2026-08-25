package connector_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

// The provider protocol is app-internal. This structural fixture prevents the
// connector from acquiring a conductor persistence path and prevents the
// conductor from acquiring a serializer for an m8.* payload.
func TestConnectorTypesHaveNoConductorPayloadRoute(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve fixture path")
	}
	repository := filepath.Clean(filepath.Join(filepath.Dir(filename), "..", ".."))
	connectorRoot := filepath.Join(repository, "internal", "connector")
	forbiddenImports := map[string]bool{
		"github.com/jackli/frank/internal/egress": true,
		"github.com/jackli/frank/internal/intake": true,
		"github.com/jackli/frank/internal/record": true,
		"github.com/jackli/frank/internal/store":  true,
	}
	walkGo(t, connectorRoot, func(path string, file *ast.File) {
		for _, imported := range file.Imports {
			value, err := strconv.Unquote(imported.Path.Value)
			if err != nil {
				t.Fatalf("%s: invalid import: %v", path, err)
			}
			if forbiddenImports[value] {
				t.Fatalf("connector acquired conductor route %q in %s", value, path)
			}
		}
	})

	conductorRoots := []string{"egress", "intake", "record", "store"}
	for _, name := range conductorRoots {
		walkGo(t, filepath.Join(repository, "internal", name), func(path string, file *ast.File) {
			for _, imported := range file.Imports {
				value, err := strconv.Unquote(imported.Path.Value)
				if err != nil {
					t.Fatalf("%s: invalid import: %v", path, err)
				}
				if strings.HasPrefix(value, "github.com/jackli/frank/internal/connector") {
					t.Fatalf("conductor package acquired m8 serializer route %q in %s", value, path)
				}
			}
		})
	}
}

func TestOpaqueLaneIdentifiersAreNeverParsedForMeaning(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve fixture path")
	}
	connectorRoot := filepath.Dir(filename)
	walkGo(t, connectorRoot, func(path string, file *ast.File) {
		if filepath.Base(filepath.Dir(path)) == "catalog" {
			return // The owning validator may enforce the identifier grammar.
		}
		ast.Inspect(file, func(node ast.Node) bool {
			switch expression := node.(type) {
			case *ast.IndexExpr:
				if opaqueSelector(expression.X) {
					t.Fatalf("opaque lane identifier indexed in %s", path)
				}
			case *ast.SliceExpr:
				if opaqueSelector(expression.X) {
					t.Fatalf("opaque lane identifier sliced in %s", path)
				}
			case *ast.CallExpr:
				selector, ok := expression.Fun.(*ast.SelectorExpr)
				if !ok || selector == nil {
					return true
				}
				packageName, packageOK := selectorPackage(selector)
				if !packageOK || (packageName != "strings" && packageName != "regexp") {
					return true
				}
				for _, argument := range expression.Args {
					if containsOpaqueSelector(argument) {
						t.Fatalf("opaque lane identifier passed to %s.%s in %s", packageName, selector.Sel.Name, path)
					}
				}
			}
			return true
		})
	})
}

func opaqueSelector(expression ast.Expr) bool {
	selector, ok := expression.(*ast.SelectorExpr)
	return ok && (selector.Sel.Name == "LaneID" || selector.Sel.Name == "ServingProfileID")
}

func containsOpaqueSelector(expression ast.Expr) bool {
	found := false
	ast.Inspect(expression, func(node ast.Node) bool {
		selector, ok := node.(*ast.SelectorExpr)
		if ok && (selector.Sel.Name == "LaneID" || selector.Sel.Name == "ServingProfileID") {
			found = true
			return false
		}
		return !found
	})
	return found
}

func selectorPackage(selector *ast.SelectorExpr) (string, bool) {
	if selector == nil {
		return "", false
	}
	identifier, ok := selector.X.(*ast.Ident)
	if !ok || identifier == nil {
		return "", false
	}
	return identifier.Name, true
}

func walkGo(t *testing.T, root string, visit func(string, *ast.File)) {
	t.Helper()
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || filepath.Ext(path) != ".go" || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		parsed, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if err != nil {
			return err
		}
		visit(path, parsed)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
