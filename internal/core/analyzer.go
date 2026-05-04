package core

type Analyzer interface {
	Name() string
	Detect(path string) (bool, error)
	Parse(filePath string) ([]Function, error)
	Analyze(fn Function) Result
}
