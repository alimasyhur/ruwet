package reporter

import (
	"testing"

	"github.com/alimasyhur/ruwet/internal/core"
)

func TestPrint_EmptyResults(t *testing.T) {
	// We can't easily capture fmt.Println output, but we can test the function doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Print panicked: %v", r)
		}
	}()

	Print([]core.Result{}, []string{})
}

func TestPrint_WithResults(t *testing.T) {
	results := []core.Result{
		{
			FunctionName: "testFunc",
			File:         "/tmp/test.go",
			Complexity:   5,
			Issues:       []string{},
			Suggestions:  []string{},
		},
		{
			FunctionName: "complexFunc",
			File:         "/tmp/complex.go",
			Complexity:   15,
			Issues:       []string{"High complexity"},
			Suggestions:  []string{"Refactor"},
		},
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Print panicked: %v", r)
		}
	}()

	Print(results, []string{"Go"})
}

func TestPrint_NoLanguages(t *testing.T) {
	results := []core.Result{
		{
			FunctionName: "testFunc",
			File:         "/tmp/test.go",
			Complexity:   3,
		},
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Print panicked: %v", r)
		}
	}()

	Print(results, []string{})
}
