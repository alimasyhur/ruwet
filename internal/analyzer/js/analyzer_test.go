package jsanalyzer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAnalyzer_Name(t *testing.T) {
	a := Analyzer{}
	if got := a.Name(); got != "JavaScript/TypeScript" {
		t.Errorf("Name() = %q, want %q", got, "JavaScript/TypeScript")
	}
}

func TestIsJSFile(t *testing.T) {
	tests := []struct {
		file string
		want bool
	}{
		{"test.js", true},
		{"test.ts", true},
		{"test.jsx", true},
		{"test.tsx", true},
		{"test.go", false},
		{"test.py", false},
	}

	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			if got := isJSFile(tt.file); got != tt.want {
				t.Errorf("isJSFile(%q) = %v, want %v", tt.file, got, tt.want)
			}
		})
	}
}

func TestAnalyzer_Parse(t *testing.T) {
	a := Analyzer{}

	tests := []struct {
		name     string
		content  string
		wantFunc int
	}{
		{
			name:     "simple function",
			content:  `function hello() { console.log("hi"); }`,
			wantFunc: 1,
		},
		{
			name: "arrow function",
			content: `const add = (a, b) => {
	return a + b;
};`,
			wantFunc: 1,
		},
		{
			name: "multiple functions",
			content: `function a() {}
function b() {}
const c = () => {};`,
			wantFunc: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile := filepath.Join(t.TempDir(), "test.js")
			os.WriteFile(tmpFile, []byte(tt.content), 0644)

			functions, err := a.Parse(tmpFile)
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}
			if len(functions) != tt.wantFunc {
				t.Errorf("Parse() returned %d functions, want %d", len(functions), tt.wantFunc)
			}
		})
	}
}

func TestAnalyzer_Analyze_Simple(t *testing.T) {
	a := Analyzer{}

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "simple.js")
	content := `function hello() {
	console.log("hello");
}`
	os.WriteFile(tmpFile, []byte(content), 0644)

	functions, _ := a.Parse(tmpFile)
	if len(functions) == 0 {
		t.Fatal("No functions parsed")
	}

	result := a.Analyze(functions[0])
	if result.Complexity < 1 {
		t.Errorf("Complexity should be at least 1, got %d", result.Complexity)
	}
}

func TestAnalyzer_Analyze_Complex(t *testing.T) {
	a := Analyzer{}

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "complex.js")
	content := `function process(x) {
	if (x > 0) {
		for (let i = 0; i < x; i++) {
			if (i % 2 == 0) {
				while (x > 0) {
					x--;
				}
			}
		}
	}
	return x;
}`
	os.WriteFile(tmpFile, []byte(content), 0644)

	functions, _ := a.Parse(tmpFile)
	if len(functions) == 0 {
		t.Fatal("No functions parsed")
	}

	result := a.Analyze(functions[0])
	// Base(1) + if(1) + for(1) + if(1) + while(1) = 5+
	if result.Complexity < 4 {
		t.Errorf("Complexity should be >= 4 for complex function, got %d", result.Complexity)
	}
}

func TestCalculateComplexity(t *testing.T) {
	a := Analyzer{}

	tests := []struct {
		name string
		code string
		want int
	}{
		{
			name: "simple",
			code: `function test() { return 1; }`,
			want: 1,
		},
		{
			name: "with if",
			code: `function test(a) { if (a > 0) { return a; } }`,
			want: 2,
		},
		{
			name: "with ternary",
			code: `function test(a) { return a > 0 ? "y" : "n"; }`,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpFile := filepath.Join(t.TempDir(), "test.js")
			os.WriteFile(tmpFile, []byte(tt.code), 0644)

			functions, _ := a.Parse(tmpFile)
			if len(functions) == 0 {
				t.Fatal("No functions")
			}

			result := a.Analyze(functions[0])
			if result.Complexity != tt.want {
				t.Errorf("Complexity = %d, want %d", result.Complexity, tt.want)
			}
		})
	}
}
