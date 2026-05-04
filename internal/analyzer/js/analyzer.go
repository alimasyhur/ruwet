package jsanalyzer

import (
	"io/ioutil"
	"path/filepath"

	sitter "github.com/smacker/go-tree-sitter"
	js "github.com/smacker/go-tree-sitter/javascript"
	ts "github.com/smacker/go-tree-sitter/typescript/typescript"

	"github.com/alimasyhur/ruwet/internal/core"
)

type Analyzer struct{}

func (a Analyzer) Name() string {
	return "JavaScript/TypeScript"
}

func (a Analyzer) Detect(path string) (bool, error) {
	return true, nil
}

func isJSFile(file string) bool {
	ext := filepath.Ext(file)
	return ext == ".js" || ext == ".ts" || ext == ".jsx" || ext == ".tsx"
}

func (a Analyzer) Parse(filePath string) ([]core.Function, error) {
	if !isJSFile(filePath) {
		return nil, nil
	}

	src, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	parser := sitter.NewParser()

	// pilih grammar
	if filepath.Ext(filePath) == ".ts" || filepath.Ext(filePath) == ".tsx" {
		parser.SetLanguage(ts.GetLanguage())
	} else {
		parser.SetLanguage(js.GetLanguage())
	}

	tree := parser.Parse(nil, src)
	root := tree.RootNode()

	functions := collectFunctions(root, src, filePath)

	return functions, nil
}

func collectFunctions(node *sitter.Node, src []byte, file string) []core.Function {
	var functions []core.Function
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		switch n.Type() {
		case "function_declaration":
			name := getNodeName(n, src)
			startLine := int(n.StartPoint().Row) + 1
			endLine := int(n.EndPoint().Row) + 1
			functions = append(functions, core.Function{
				Name:      name,
				File:      file,
				StartLine: startLine,
				EndLine:   endLine,
			})
		case "method_definition":
			name := getNodeName(n, src)
			startLine := int(n.StartPoint().Row) + 1
			endLine := int(n.EndPoint().Row) + 1
			functions = append(functions, core.Function{
				Name:      name,
				File:      file,
				StartLine: startLine,
				EndLine:   endLine,
			})
		case "arrow_function":
			name := getNodeName(n, src)
			if name == "anonymous" {
				name = "arrow_function"
			}
			startLine := int(n.StartPoint().Row) + 1
			endLine := int(n.EndPoint().Row) + 1
			functions = append(functions, core.Function{
				Name:      name,
				File:      file,
				StartLine: startLine,
				EndLine:   endLine,
			})
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(node)
	return functions
}

func (a Analyzer) Analyze(fn core.Function) core.Result {
	src, err := ioutil.ReadFile(fn.File)
	if err != nil {
		return core.Result{}
	}

	parser := sitter.NewParser()

	if filepath.Ext(fn.File) == ".ts" || filepath.Ext(fn.File) == ".tsx" {
		parser.SetLanguage(ts.GetLanguage())
	} else {
		parser.SetLanguage(js.GetLanguage())
	}

	tree := parser.Parse(nil, src)
	root := tree.RootNode()

	var result core.Result
	var found bool
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found {
			return
		}
		switch n.Type() {
		case "function_declaration", "method_definition", "arrow_function":
			name := getNodeName(n, src)
			if name == "anonymous" && n.Type() == "arrow_function" {
				name = "arrow_function"
			}
			startLine := int(n.StartPoint().Row) + 1
			if name == fn.Name && startLine == fn.StartLine {
				complexity := calculateComplexity(n, src)
				result = buildResult(fn.Name, fn.File, complexity)
				found = true
				return
			}
		}
		for i := 0; i < int(n.ChildCount()) && !found; i++ {
			walk(n.Child(i))
		}
	}
	walk(root)

	if !found {
		return core.Result{}
	}

	return result
}

func getNodeName(node *sitter.Node, src []byte) string {
	nameNode := node.ChildByFieldName("name")
	if nameNode != nil {
		return nameNode.Content(src)
	}
	return "anonymous"
}

func calculateComplexity(node *sitter.Node, src []byte) int {
	complexity := 1

	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {

		switch n.Type() {

		case "if_statement",
			"for_statement",
			"while_statement",
			"do_statement",
			"catch_clause",
			"ternary_expression":
			complexity++

		case "switch_case":
			complexity++

		case "binary_expression":
			op := n.ChildByFieldName("operator")
			if op != nil {
				start := int(op.StartByte())
				end := int(op.EndByte())
				if start < len(src) && end <= len(src) && start < end {
					val := string(src[start:end])
					if val == "&&" || val == "||" {
						complexity++
					}
				}
			}
		}

		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}

	walk(node)

	return complexity
}

func buildResult(name, file string, complexity int) core.Result {
	issues := []string{}
	suggestions := []string{}

	if complexity > 15 {
		issues = append(issues, "Very high complexity")
		suggestions = append(suggestions, "Break into smaller functions")
	} else if complexity > 10 {
		issues = append(issues, "High complexity")
		suggestions = append(suggestions, "Refactor for readability")
	} else if complexity > 5 {
		issues = append(issues, "Moderate complexity")
	}

	return core.Result{
		FunctionName: name,
		File:         file,
		Complexity:   complexity,
		Issues:       issues,
		Suggestions:  suggestions,
	}
}
