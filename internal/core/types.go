package core

type Function struct {
	Name       string
	File       string
	StartLine  int
	EndLine    int
	Complexity int
}

type Result struct {
	FunctionName string
	File         string
	Complexity   int
	Issues       []string
	Suggestions  []string
}
