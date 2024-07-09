package main

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const patternsDir = "./patterns/"

func getPatternList() []string {
	var patterns []string
	err := filepath.Walk(patternsDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
					printError(err)
					return err
			}

			dir := strings.Split(path, "/")[1]

			if info.IsDir() && path != patternsDir && !slices.Contains(patterns, dir) {
				patterns = append(patterns, dir)
			}

			return nil
	})
	if err != nil {
		printError(err)
		return nil
	}
	return patterns
}

func getPatternPrompt(pattern string) string {
	patterns := getPatternList()
	if !slices.Contains(patterns, pattern) {
		return ""
	}

	content, err := os.ReadFile(patternsDir + "/" + pattern + "/system.md")

	if err != nil {
		printError("Could not obtain pattern text.")
		return ""
	}

	return string(content)
}