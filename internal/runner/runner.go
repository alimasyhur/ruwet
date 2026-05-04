package runner

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/alimasyhur/ruwet/internal/core"
	"github.com/alimasyhur/ruwet/internal/utils"
)

func Run(path string) ([]core.Result, []string, error) {
	analyzers := core.GetAnalyzers()

	files := utils.CollectFiles(path)
	total := len(files)

	fmt.Println("🔍 Scanning files...")

	bar := utils.NewProgressBar(total)

	var allResults []core.Result
	var usedAnalyzers []string

	for i, file := range files {
		// Tampilkan hanya nama file saja, bukan full path
		// Ini agar progress bar tidak wrap ke baris baru
		bar.Render(i+1, "Processing: "+filepath.Base(file))

		for _, analyzer := range analyzers {
			ok, _ := analyzer.Detect(path)
			if !ok {
				continue
			}

			functions, err := analyzer.Parse(file)
			if err != nil || len(functions) == 0 {
				continue
			}

			// tandai analyzer dipakai
			found := false
			for _, name := range usedAnalyzers {
				if name == analyzer.Name() {
					found = true
					break
				}
			}
			if !found {
				usedAnalyzers = append(usedAnalyzers, analyzer.Name())
			}

			for _, fn := range functions {
				res := analyzer.Analyze(fn)
				allResults = append(allResults, res)
			}
		}
	}

	fmt.Print("\n\n")
	fmt.Println("✅ Scan complete")

	sort.Slice(allResults, func(i, j int) bool {
		return allResults[i].Complexity > allResults[j].Complexity
	})

	return allResults, usedAnalyzers, nil
}
