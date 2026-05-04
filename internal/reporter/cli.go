package reporter

import (
	"fmt"
	"strings"

	"github.com/fatih/color"

	"github.com/alimasyhur/ruwet/internal/core"
)

var (
	red    = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
	bold   = color.New(color.Bold).SprintFunc()
)

func Print(results []core.Result, langs []string) {

	fmt.Println(bold("🚀 RUWET ANALYSIS RESULT"))
	fmt.Println(strings.Repeat("=", 40))

	// language
	fmt.Printf("✔ Detected: %s\n\n", cyan(strings.Join(langs, ", ")))

	fmt.Printf("📊 Total Functions: %s\n\n", bold(len(results)))

	fmt.Println("Legend:")
	fmt.Println(red("■ High (>15)"), yellow("■ Medium (>8)"), green("■ Low (≤8)"))
	fmt.Println("")

	if len(results) == 0 {
		fmt.Println("⚠ No functions found or analyzed")
		return
	}

	fmt.Println(bold("🔥 Top messy functions:\n"))

	for i, r := range results {
		if i >= 10 {
			break
		}

		colorFn := getColor(r.Complexity)

		fmt.Printf("%s %s\n",
			colorFn(fmt.Sprintf("%2d.", i+1)),
			bold(r.FunctionName),
		)

		fmt.Printf("   📁 %s\n", r.File)
		fmt.Printf("   %s Complexity: %s\n", severityIcon(r.Complexity), colorFn(r.Complexity))

		for _, issue := range r.Issues {
			fmt.Printf("   ⚠ %s\n", colorFn(issue))
		}

		fmt.Println("")
	}
}

func getColor(complexity int) func(a ...interface{}) string {
	if complexity > 15 {
		return red
	} else if complexity > 8 {
		return yellow
	}
	return green
}

func severityIcon(c int) string {
	if c > 15 {
		return "🔥"
	} else if c > 8 {
		return "⚠"
	}
	return "✅"
}
