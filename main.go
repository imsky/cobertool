package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Line struct {
	Hit    bool
	Number uint64
}

type SourceFile struct {
	Lines []Line
}

func Run(reports []string) {
	//todo: rework this, it's really a map of source file names to line sets
	sourceFiles := make(map[string][]SourceFile)

	for _, reportFileName := range reports {
		file, _ := os.Open(reportFileName)
		decoder := xml.NewDecoder(file)

		var sourceFilename string
		var lines []Line = make([]Line, 0, 100)

		for {
			t, _ := decoder.Token()
			if t == nil {
				break
			}

			switch se := t.(type) {
			case xml.StartElement:
				switch se.Name.Local {
				case "class":
					for _, attr := range se.Attr {
						if attr.Name.Local == "filename" {
							sourceFilename = attr.Value
							if _, exists := sourceFiles[sourceFilename]; !exists {
								sourceFiles[sourceFilename] = make([]SourceFile, 0, len(reports))
							}
						}
					}
				case "line":
					//todo: check that class file name is set
					var line Line

					for _, attr := range se.Attr {
						if attr.Name.Local == "number" {
							line.Number, _ = strconv.ParseUint(attr.Value, 10, 64)
						} else if attr.Name.Local == "hits" {
							line.Hit = attr.Value != "0"
						}
					}

					lines = append(lines, line)
				}
			case xml.EndElement:
				switch se.Name.Local {
				case "class":
					if len(lines) > 0 {
						class := SourceFile{Lines: lines}
						sourceFiles[sourceFilename] = append(sourceFiles[sourceFilename], class)
					}

					//todo: use a function to reset lines
					lines = make([]Line, 0, 100)
				}

			}

		}

		file.Close()
	}

	var totalCoverage float64

	//todo: validate that each source file list of lines is same length
	for _, fileSets := range sourceFiles {
		fileLineHitMap := make(map[uint64]bool)

		for _, sourceFile := range fileSets {
			for _, line := range sourceFile.Lines {
				fileLineHitMap[line.Number] = fileLineHitMap[line.Number] || line.Hit
			}
		}

		totalLines := len(fileLineHitMap)
		hitLines := 0

		for _, hit := range fileLineHitMap {
			if hit {
				hitLines++
			}
		}

		if totalLines > 0 {
			totalCoverage += float64(hitLines) / float64(totalLines)
		}
	}

	fmt.Println(totalCoverage / float64(len(sourceFiles)))
}

//todo: add flag for path-rewriting when writing source

func main() {
	flag.Parse()
	Run(flag.Args())
}
