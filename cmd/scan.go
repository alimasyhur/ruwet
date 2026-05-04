package cmd

import (
	goanalyzer "github.com/alimasyhur/ruwet/internal/analyzer/go"
	jsanalyzer "github.com/alimasyhur/ruwet/internal/analyzer/js"
	"github.com/alimasyhur/ruwet/internal/core"
	"github.com/alimasyhur/ruwet/internal/reporter"
	"github.com/alimasyhur/ruwet/internal/runner"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [path]",
	Short: "Scan project",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		core.Register(goanalyzer.Analyzer{})
		core.Register(jsanalyzer.Analyzer{})

		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		results, lang, _ := runner.Run(path)

		reporter.Print(results, lang)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
