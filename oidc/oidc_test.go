package oidc

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestPublicMethodsHaveDocumentation keeps the public API intentional.  It
// also exercises the source-level contract without requiring an OIDC provider
// (and network access) in the test suite.
func TestPublicMethodsHaveDocumentation(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}

	file, err := parser.ParseFile(token.NewFileSet(), filepath.Join(filepath.Dir(filename), "oidc.go"), nil, 0)
	if err != nil {
		t.Fatalf("parse oidc.go: %v", err)
	}

	for _, declaration := range file.Decls {
		function, ok := declaration.(*ast.FuncDecl)
		if !ok || function.Recv == nil || function.Name == nil || !ast.IsExported(function.Name.Name) {
			continue
		}

		if function.Doc == nil || strings.TrimSpace(function.Doc.Text()) == "" {
			t.Errorf("exported method %s should have a documentation comment", function.Name.Name)
		}
	}
}
