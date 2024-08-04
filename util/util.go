package util

import (
	"fmt"
	"io"
	"os"

	"github.com/cetorres/ai-buddy/color"
)

func PrintError(err any) {
	color.PrintColor(color.COLOR_RED, fmt.Sprintf("ERROR: %s", err))
}

func PrintWarning(text any) {
	color.PrintColor(color.COLOR_YELLOW, fmt.Sprintf("WARNING: %s", text))
}

func PrintSuccess(text any) {
	color.PrintColor(color.COLOR_GREEN, text)
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