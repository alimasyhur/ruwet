package core

var analyzers []Analyzer

func Register(a Analyzer) {
	analyzers = append(analyzers, a)
}

func GetAnalyzers() []Analyzer {
	return analyzers
}
