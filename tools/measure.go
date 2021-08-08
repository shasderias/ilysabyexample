package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
)

const maxLen = 120

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readLines(path string) []string {
	srcBytes, err := ioutil.ReadFile(path)
	check(err)
	return strings.Split(string(srcBytes), "\n")
}

var commentPat = regexp.MustCompile("\\s*\\/\\/")

func main() {
	sourcePaths, err := filepath.Glob("./examples/*/*")
	check(err)
	foundLongFile := false
	for _, sourcePath := range sourcePaths {
		foundLongLine := false
		lines := readLines(sourcePath)
		for i, line := range lines {
			// Convert tabs to spaces before measuring, so we get an accurate measure
			// of how long the output will end up being.
			line := strings.Replace(line, "\t", "    ", -1)
			if !foundLongLine && !commentPat.MatchString(line) && (utf8.RuneCountInString(line) > maxLen) {
				fmt.Printf("measure: %s:%d\n", sourcePath, i+1)
				foundLongLine = true
				foundLongFile = true
			}
		}
	}
	if foundLongFile {
		os.Exit(1)
	}
}
