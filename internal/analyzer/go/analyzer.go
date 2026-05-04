package goanalyzer

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"github.com/alimasyhur/ruwet/internal/core"
)

type Analyzer struct{}

func (a Analyzer) Name() string {
	return "Go"
}

func (a Analyzer) Detect(path string) (bool, error) {
	_, err := os.Stat(filepath.Join(path, "go.mod"))
	return err == nil, nil
}

func (a Analyzer) Parse(filePath string) ([]core.Function, error) {
	var functions []core.Function

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		return nil, err
	}

	for _, decl := range node.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			start := fset.Position(fn.Pos()).Line
			end := fset.Position(fn.End()).Line

			functions = append(functions, core.Function{
				Name:      fn.Name.Name,
				File:      filePath,
				StartLine: start,
				EndLine:   end,
			})
		}
	}

	return functions, nil
}

func (a Analyzer) Analyze(fn core.Function) core.Result {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, fn.File, nil, 0)
	if err != nil {
		return core.Result{}
	}

	complexity := 1

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {

		case *ast.FuncDecl:
			if x.Name.Name != fn.Name {
				return false
			}

		case *ast.IfStmt:
			complexity++

		case *ast.ForStmt:
			complexity++

		case *ast.RangeStmt:
			complexity++

		case *ast.CaseClause:
			complexity++

		case *ast.BinaryExpr:
			if x.Op.String() == "&&" || x.Op.String() == "||" {
				complexity++
			}
		}
		return true
	})

	issues := []string{}
	suggestions := []string{}

	if complexity > 10 {
		issues = append(issues, "High cyclomatic complexity")
		suggestions = append(suggestions, "Split into smaller functions")
	}

	return core.Result{
		FunctionName: fn.Name,
		File:         fn.File,
		Complexity:   complexity,
		Issues:       issues,
		Suggestions:  suggestions,
	}
}
