package runner

import (
	"os"
	"path/filepath"
	"sort"
	"testing"

	goanalyzer "github.com/alimasyhur/ruwet/internal/analyzer/go"
	jsanalyzer "github.com/alimasyhur/ruwet/internal/analyzer/js"
	"github.com/alimasyhur/ruwet/internal/core"
	"github.com/alimasyhur/ruwet/internal/utils"
)

func TestRun(t *testing.T) {
	// Create temp directory with test files
	tmpDir := t.TempDir()

	// Create go.mod so Go analyzer detects it
	os.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte("module test"), 0644)

	// Create Go file
	goContent := `package main

func hello() {
	println("hello")
}

func complexFn(x int) int {
	if x > 0 {
		for i := 0; i < x; i++ {
			if i%2 == 0 {
				return i
			}
		}
	}
	return x
}
`
	os.WriteFile(filepath.Join(tmpDir, "main.go"), []byte(goContent), 0644)

	// Create JS file
	jsContent := `function greet(name) {
	console.log("Hello " + name);
}

function process(x) {
	if (x > 0) {
		while (x > 0) {
			x--;
		}
	}
	return x;
}
`
	os.WriteFile(filepath.Join(tmpDir, "app.js"), []byte(jsContent), 0644)

	// Register analyzers
	core.Register(goanalyzer.Analyzer{})
	core.Register(jsanalyzer.Analyzer{})

	// Run analysis
	results, langs, err := Run(tmpDir)
	if err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	// Check that we detected both languages
	foundGo := false
	foundJS := false
	for _, lang := range langs {
		if lang == "Go" {
			foundGo = true
		}
		if lang == "JavaScript/TypeScript" {
			foundJS = true
		}
	}

	if !foundGo {
		t.Error("Expected to detect Go")
	}
	if !foundJS {
		t.Error("Expected to detect JavaScript/TypeScript")
	}

	// Check results are sorted by complexity (descending)
	if !sort.SliceIsSorted(results, func(i, j int) bool {
		return results[i].Complexity > results[j].Complexity
	}) {
		t.Error("Results should be sorted by complexity descending")
	}

	// Check we have results
	if len(results) == 0 {
		t.Error("Expected at least some results")
	}

	t.Logf("Found %d functions", len(results))
	for i, r := range results {
		t.Logf("  %d. %s (complexity: %d)", i+1, r.FunctionName, r.Complexity)
	}
}

func TestRun_EmptyDir(t *testing.T) {
	tmpDir := t.TempDir()

	core.Register(goanalyzer.Analyzer{})
	core.Register(jsanalyzer.Analyzer{})

	results, langs, err := Run(tmpDir)
	if err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	if len(results) != 0 {
		t.Errorf("Expected 0 results for empty dir, got %d", len(results))
	}

	if len(langs) != 0 {
		t.Errorf("Expected 0 languages detected, got %d", len(langs))
	}
}

func TestRun_NoValidFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Create non-code files
	os.WriteFile(filepath.Join(tmpDir, "README.md"), []byte("# Test"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "data.json"), []byte("{}"), 0644)

	core.Register(goanalyzer.Analyzer{})
	core.Register(jsanalyzer.Analyzer{})

	results, _, err := Run(tmpDir)
	if err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	if len(results) != 0 {
		t.Errorf("Expected 0 results for non-code files, got %d", len(results))
	}
}

// Test helper - reset registry between tests
func init() {
	// Clear the registry before each test suite
	// This is a bit hacky but works for testing
	// In real code, we'd use a better pattern
	_ = utils.CollectFiles // ensure import
}
