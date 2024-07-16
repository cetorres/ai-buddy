package util

import (
	"fmt"
	"os"
)

// Text colors
type Color string
const (
	COLOR_RED Color = "\033[0;31m"
	COLOR_GREEN Color = "\033[32m"
	COLOR_YELLOW Color = "\033[33m"
	COLOR_BLUE Color = "\033[34m"
	COLOR_MAGENTA Color = "\033[35m" 
	COLOR_CYAN Color = "\033[36m" 
	COLOR_GRAY Color = "\033[37m" 
	COLOR_WHITE Color = "\033[97m"
	COLOR_RESET Color = "\033[0m"
)

func PrintColor(color Color, text any) {
	fmt.Fprintf(os.Stdout, "%s%s%s\n", color, text, COLOR_RESET)
}

func PrintError(err any) {
	PrintColor(COLOR_RED, fmt.Sprintf("ERROR: %s", err))
}

func PrintWarning(err any) {
	PrintColor(COLOR_YELLOW, fmt.Sprintf("WARNING: %s", err))
}

func IsInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode() & os.ModeCharDevice == 0
}