package pattern

import (
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/cetorres/ai-buddy/config"
	"github.com/cetorres/ai-buddy/util"
)

func GetPatternsDir() string {
	configDir := config.GetConfigDirectory()

	if configDir != "" {
		dirPath := configDir + "patterns"
		return dirPath
	}
	return "./patterns"
}

func IsExistPatternDir() bool {
	dir := GetPatternsDir()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func CopyPatternsDirToConfigDir() error {
	oldDir := "./patterns"
	newDir := GetPatternsDir()
	err := util.CopyDir(oldDir, newDir)
	return err
}

func GetPatternList() ([]string, error) {
	var patterns []string
	patternsDir := GetPatternsDir()

	err := filepath.Walk(patternsDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
					return err
			}

			pathParts := strings.Split(path, "/")
			dir := pathParts[max(len(pathParts) - 1,1)]
			if info.IsDir() && path != patternsDir && !slices.Contains(patterns, dir) {
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