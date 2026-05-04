package goanalyzer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAnalyzer_Name(t *testing.T) {
	a := Analyzer{}
	if got := a.Name(); got != "Go" {
		t.Errorf("Name() = %q, want %q", got, "Go")
	}
}

func TestAnalyzer_Detect(t *testing.T) {
	a := Analyzer{}

	// Create temp dir with go.mod
	tmpDir := t.TempDir()
	os.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte("module test"), 0644)

	got, err := a.Detect(tmpDir)
	if err != nil {
		t.Errorf("Detect() error = %v", err)
	}
	if !got {
		t.Error("Detect() should return true when go.mod exists")
	}
}

func TestAnalyzer_Parse(t *testing.T) {
	a := Analyzer{}

	// Create temp Go file
	tmpFile := filepath.Join(t.TempDir(), "test.go")
	content := `package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

func greet(name string) {
	if name != "" {
		fmt.Println("Hello", name)
	}
}
`
	os.WriteFile(tmpFile, []byte(content), 0644)

	functions, err := a.Parse(tmpFile)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if len(functions) != 2 {
		t.Errorf("Parse() returned %d functions, want 2", len(functions))
	}

	names := make(map[string]bool)
	for _, fn := range functions {
		names[fn.Name] = true
	}

	for _, want := range []string{"hello", "greet"} {
		if !names[want] {
			t.Errorf("Function %q not found in parsed results", want)
		}
	}
}

func TestAnalyzer_Analyze_Simple(t *testing.T) {
	a := Analyzer{}

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.go")
	content := `package main

func simple() {
	println("hello")
}
`
	os.WriteFile(tmpFile, []byte(content), 0644)

	functions, _ := a.Parse(tmpFile)
	if len(functions) == 0 {
		t.Fatal("No functions parsed")
	}

	result := a.Analyze(functions[0])
	if result.Complexity < 1 {
		t.Errorf("Complexity should be at least 1, got %d", result.Complexity)
	}
	if result.FunctionName != "simple" {
		t.Errorf("FunctionName = %q, want %q", result.FunctionName, "simple")
	}
}

func TestAnalyzer_Analyze_Complex(t *testing.T) {
	a := Analyzer{}

	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "complex.go")
	content := `package main

func complex(x int) int {
	if x > 0 {
		for i := 0; i < x; i++ {
			if i%2 == 0 {
				switch i {
				case 0:
					return 1
				case 2:
					return 2
				}
			}
		}
	}
	return x
}
`
	os.WriteFile(tmpFile, []byte(content), 0644)

	functions, _ := a.Parse(tmpFile)
	if len(functions) == 0 {
		t.Fatal("No functions parsed")
	}

	result := a.Analyze(functions[0])
	// Base(1) + if(1) + for(1) + if(1) + switch with 2 cases(2) = 6+
	if result.Complexity < 5 {
		t.Errorf("Complexity should be >= 5 for complex function, got %d", result.Complexity)
	}
	if len(result.Issues) > 0 {
		t.Logf("Issues found: %v", result.Issues)
	}
}
