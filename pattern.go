package main

import (
	"os"
	"path/filepath"
	"slices"
	"strings"
)

const PATTERNS_DIR_ENV = "AI_BUDDY_PATTERNS"

func getPatternsDir() string {
	if os.Getenv(PATTERNS_DIR_ENV) != "" {
		dir := os.Getenv(PATTERNS_DIR_ENV)
		dir = strings.TrimSuffix(dir, "/")
		if !strings.Contains(dir, "/patterns") {
			dir = dir + "/patterns"
		}
		return dir
	}
	return "./patterns" 
}

func getPatternList() ([]string, error) {
	var patterns []string

	err := filepath.Walk(getPatternsDir(), func(path string, info os.FileInfo, err error) error {
			if err != nil {
					return err
			}

			pathParts := strings.Split(path, "/")
			dir := pathParts[max(len(pathParts) - 1,1)]
			if info.IsDir() && path != getPatternsDir() && !slices.Contains(patterns, dir) {
				patterns = append(patterns, dir)
			}

			return nil
	})

	if err != nil {
		return nil, err
	}

	return patterns, nil
}

func getPatternPrompt(pattern string) string {
	patterns, _ := getPatternList()

	if patterns == nil || !slices.Contains(patterns, pattern) {
		return ""
	}

	content, err := os.ReadFile(getPatternsDir() + "/" + pattern + "/system.md")

	if err != nil {
		return ""
	}

	return string(content)
}