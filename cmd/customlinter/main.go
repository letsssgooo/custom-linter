package main

import (
	"log"

	"github.com/letsssgooo/custom-linter/internal/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	log.Println("ваиваиваиваи")
	singlechecker.Main(analyzer.Analyzer)
}
