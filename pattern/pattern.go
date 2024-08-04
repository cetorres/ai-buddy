package pattern

import (
	"os"
	"slices"

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

	files, err := os.ReadDir(patternsDir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			patterns = append(patterns, f.Name())
		}
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