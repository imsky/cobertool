package main

import (
	"testing"
)

func TestComputeTotalCoverage(t *testing.T) {
	sourceFiles := make(map[string][]SourceFile)
	a := SourceFile{Lines: []Line{{Number: 1, Hit: true}, {Number: 2, Hit: true}, {Number: 3, Hit: false}}}
	b := SourceFile{Lines: []Line{{Number: 1, Hit: false}, {Number: 2, Hit: false}, {Number: 3, Hit: true}}}

	sourceFiles["a"] = []SourceFile{a}
	coverage := ComputeTotalCoverage(sourceFiles)
	if int(coverage) != 66 {
		t.Errorf("Incorrect single source file coverage calculation: %v", coverage)
	}

	sourceFiles["b"] = []SourceFile{b}
	coverage = ComputeTotalCoverage(sourceFiles)
	if int(coverage) != 50 {
		t.Errorf("Incorrect multiple source file coverage calculation: %v", coverage)
	}

	sourceFiles["a"] = []SourceFile{a, b}
	coverage = ComputeTotalCoverage(sourceFiles)
	if int(coverage) != 66 {
		t.Errorf("Incorrect merged source file coverage calculation: %v", coverage)
	}
}

func TestCobertool(t *testing.T) {
	main()
	Run([]string{"test-report.xml"})
}
