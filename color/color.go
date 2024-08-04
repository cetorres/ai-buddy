package color

import (
	"fmt"
	"os"
	"runtime"
)

type Color string

var COLOR_RESET Color   = "\033[0m"
var COLOR_RED Color     = "\033[31m"
var COLOR_GREEN Color   = "\033[32m"
var COLOR_YELLOW Color  = "\033[33m"
var COLOR_BLUE Color    = "\033[34m"
var COLOR_MAGENTA Color = "\033[35m"
var COLOR_CYAN Color    = "\033[36m"
var COLOR_GRAY Color    = "\033[37m"
var COLOR_WHITE Color   = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		COLOR_RESET  = ""
		COLOR_RED    = ""
		COLOR_GREEN  = ""
		COLOR_YELLOW = ""
		COLOR_BLUE   = ""
		COLOR_MAGENTA = ""
		COLOR_CYAN   = ""
		COLOR_GRAY   = ""
		COLOR_WHITE  = ""
	}
}

func PrintColor(color Color, text any) {
	fmt.Fprintf(os.Stdout, "%s%s%s\n", color, text, COLOR_RESET)
}