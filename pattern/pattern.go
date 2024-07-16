package pattern

import (
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/cetorres/ai-buddy/constants"
)

func GetPatternsDir() string {
	if os.Getenv(constants.PATTERNS_DIR_ENV) != "" {
		dir := os.Getenv(constants.PATTERNS_DIR_ENV)
		dir = strings.TrimSuffix(dir, "/")
		if !strings.Contains(dir, "/patterns") {
			dir = dir + "/patterns"
		}
		return dir
	}
	return "./patterns" 
}

func GetPatternList() ([]string, error) {
	var patterns []string

	err := filepath.Walk(GetPatternsDir(), func(path string, info os.FileInfo, err error) error {
			if err != nil {
					return err
			}

			pathParts := strings.Split(path, "/")
			dir := pathParts[max(len(pathParts) - 1,1)]
			if info.IsDir() && path != GetPatternsDir() && !slices.Contains(patterns, dir) {
				patterns = append(patterns, dir)
			}

			return nil
	})

	if err != nil {
		return nil, err
	}

	return patterns, nil
}

func GetPatternPrompt(pattern string) string {
	patterns, _ := GetPatternList()

	if patterns == nil || !slices.Contains(patterns, pattern) {
		return ""
	}

	content, err := os.ReadFile(GetPatternsDir() + "/" + pattern + "/system.md")

	if err != nil {
		return ""
	}

	return string(content)
}