package seatclient_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestImportGraphAndMappingHome(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	formschema := filepath.Join(root, "internal", "seatclient", "formschema")
	conduct := filepath.Join(root, "internal", "seatclient", "conduct")
	formschemaImports := packageImports(t, formschema)
	for imported := range formschemaImports {
		if strings.HasPrefix(imported, "github.com/jackli/frank/internal/") && imported != "github.com/jackli/frank/internal/fieldspec" && imported != "github.com/jackli/frank/internal/record" {
			t.Fatalf("formschema imports forbidden internal package %s", imported)
		}
	}
	if packageImports(t, conduct)["github.com/jackli/frank/internal/seatclient/formschema"] {
		t.Fatal("conduct imports formschema")
	}
	mcpSchema := filepath.Join(root, "cmd", "frank-mcp", "schema.go")
	contents, err := os.ReadFile(mcpSchema)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(contents), "internal/seatclient/formschema") {
		t.Fatal("MCP frontend does not consume shared formschema module")
	}
	parsed, err := parser.ParseFile(token.NewFileSet(), mcpSchema, contents, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, declaration := range parsed.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if ok && function.Name.Name == "propertyForField" {
			t.Fatal("mapping implementation remains trapped in cmd/frank-mcp")
		}
	}
}

func packageImports(t *testing.T, directory string) map[string]bool {
	t.Helper()
	packages, err := parser.ParseDir(token.NewFileSet(), directory, func(info os.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go")
	}, parser.ImportsOnly)
	if err != nil {
		t.Fatal(err)
	}
	imports := make(map[string]bool)
	for _, pkg := range packages {
		for _, file := range pkg.Files {
			for _, imported := range file.Imports {
				path, err := strconv.Unquote(imported.Path.Value)
				if err != nil {
					t.Fatal(err)
				}
				imports[path] = true
			}
		}
	}
	return imports
}
