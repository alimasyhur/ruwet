package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCollectFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Create directory structure
	os.MkdirAll(filepath.Join(tmpDir, "src"), 0755)
	os.MkdirAll(filepath.Join(tmpDir, "node_modules"), 0755)
	os.MkdirAll(filepath.Join(tmpDir, "dist"), 0755)

	// Create test files
	os.WriteFile(filepath.Join(tmpDir, "main.go"), []byte("package main"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "src", "app.js"), []byte("console.log('test')"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "src", "types.ts"), []byte("export const x = 1;"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "node_modules", "lib.js"), []byte("module.exports = {}"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "dist", "bundle.js"), []byte("console.log('bundle')"), 0644)

	files := CollectFiles(tmpDir)

	// Should only include main.go, app.js, types.ts
	if len(files) != 3 {
		t.Errorf("CollectFiles() returned %d files, want 3", len(files))
		for _, f := range files {
			t.Logf("  - %s", f)
		}
	}

	// Verify no node_modules or dist files
	for _, f := range files {
		if filepath.Base(filepath.Dir(f)) == "node_modules" || filepath.Base(filepath.Dir(f)) == "dist" {
			t.Errorf("CollectFiles() should not include %s", f)
		}
	}
}

func TestCollectFilesWithGitignore(t *testing.T) {
	tmpDir := t.TempDir()

	// Create .gitignore
	gitignore := `node_modules/
dist/
*.log
build/
`
	os.WriteFile(filepath.Join(tmpDir, ".gitignore"), []byte(gitignore), 0644)

	// Create test files
	os.WriteFile(filepath.Join(tmpDir, "app.go"), []byte("package main"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "app.js"), []byte("console.log('test')"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "debug.log"), []byte("log data"), 0644)
	os.MkdirAll(filepath.Join(tmpDir, "node_modules"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "node_modules", "lib.js"), []byte("module.exports = {}"), 0644)
	os.MkdirAll(filepath.Join(tmpDir, "build"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "build", "output.js"), []byte("console.log('build')"), 0644)

	files := CollectFiles(tmpDir)

	// Should only include app.go and app.js
	if len(files) != 2 {
		t.Errorf("CollectFiles() returned %d files, want 2", len(files))
		for _, f := range files {
			t.Logf("  - %s", f)
		}
	}

	// Verify no gitignored files
	for _, f := range files {
		base := filepath.Base(f)
		if base == "debug.log" || base == "lib.js" || base == "output.js" {
			t.Errorf("CollectFiles() should not include gitignored file: %s", f)
		}
	}
}

func TestCollectFilesEmptyDir(t *testing.T) {
	tmpDir := t.TempDir()

	files := CollectFiles(tmpDir)
	if len(files) != 0 {
		t.Errorf("CollectFiles() returned %d files, want 0", len(files))
	}
}

func TestCollectFilesNoGitignore(t *testing.T) {
	tmpDir := t.TempDir()

	// No .gitignore file
	os.WriteFile(filepath.Join(tmpDir, "app.go"), []byte("package main"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "app.js"), []byte("console.log('test')"), 0644)

	files := CollectFiles(tmpDir)
	if len(files) != 2 {
		t.Errorf("CollectFiles() returned %d files, want 2", len(files))
	}
}

func TestLoadGitignore(t *testing.T) {
	tmpDir := t.TempDir()

	// Create .gitignore
	gitignoreContent := `node_modules/
*.log
build/
`
	os.WriteFile(filepath.Join(tmpDir, ".gitignore"), []byte(gitignoreContent), 0644)

	ignorer := loadGitignore(tmpDir)

	tests := []struct {
		path string
		want bool
	}{
		{"node_modules/lib.js", true},
		{"src/app.js", false},
		{"debug.log", true},
		{"build/output.js", true},
		{"src/main.go", false},
	}

	for _, tt := range tests {
		got := ignorer.MatchesPath(tt.path)
		if got != tt.want {
			t.Errorf("MatchesPath(%q) = %v, want %v", tt.path, got, tt.want)
		}
	}
}
