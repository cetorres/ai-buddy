package util

import (
	"fmt"
	"io"
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

func CopyFile(source string, dest string) error {
	sourcefile, err := os.Open(source)
	if err != nil {
			return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
			return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
			sourceinfo, err := os.Stat(source)
			if err != nil {
					err = os.Chmod(dest, sourceinfo.Mode())
			}
	}

	return nil
}

func CopyDir(source string, dest string) error {
	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
			return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
			return err
	}

	directory, _ := os.Open(source)
	objects, err := directory.Readdir(-1)

	for _, obj := range objects {
			sourcefilepointer := source + "/" + obj.Name()
			destinationfilepointer := dest + "/" + obj.Name()

			if obj.IsDir() {
					// create sub-directories - recursively
					err = CopyDir(sourcefilepointer, destinationfilepointer)
					if err != nil {
							fmt.Println(err)
					}
			} else {
					// perform copy
					err = CopyFile(sourcefilepointer, destinationfilepointer)
					if err != nil {
							fmt.Println(err)
					}
			}
	}

	return nil
}