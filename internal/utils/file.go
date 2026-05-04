package utils

import (
	"os"
	"path/filepath"
	"strings"

	gitignore "github.com/sabhiram/go-gitignore"
)

func CollectFiles(root string) []string {
	var files []string

	// Initialize gitignore parser
	ignorer := loadGitignore(root)

	// Hardcoded exclusions (fallback if .gitignore doesn't cover them)
	excludeDirs := map[string]bool{
		"node_modules": true,
		"vendor":        true,
		"prisma":        true,
		".next":         true,
		"public":        true,
		"dist":          true,
	}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		// Skip directories that should not be traversed
		if info.IsDir() {
			name := info.Name()
			// Always skip .git directory
			if name == ".git" {
				return filepath.SkipDir
			}
			// Skip hardcoded excluded directories
			if excludeDirs[name] {
				return filepath.SkipDir
			}
		}

		// Get relative path for gitignore check
		relPath, _ := filepath.Rel(root, path)
		if relPath == "" {
			relPath = "."
		}

		// Check if file should be ignored by .gitignore
		if ignorer.MatchesPath(relPath) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		// Additional check: skip if path contains excluded dirs
		parts := strings.Split(path, string(os.PathSeparator))
		for _, part := range parts {
			if excludeDirs[part] {
				if info.IsDir() && part == info.Name() {
					return filepath.SkipDir
				}
				return nil
			}
		}

		// Only collect code files
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		switch ext {
		case ".go", ".js", ".ts", ".jsx", ".tsx":
			files = append(files, path)
		}

		return nil
	})

	return files
}

// loadGitignore loads .gitignore rules from the root directory
func loadGitignore(root string) *gitignore.GitIgnore {
	gitignorePath := filepath.Join(root, ".gitignore")

	// Try to load .gitignore file
	ignorer, err := gitignore.CompileIgnoreFile(gitignorePath)
	if err != nil {
		// If .gitignore doesn't exist or fails, create empty ignorer
		return gitignore.CompileIgnoreLines()
	}

	return ignorer
}
