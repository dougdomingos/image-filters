package main

import (
	"fmt"
	"os"
)

const (
	// ImageLoadingError declares the exit code for failures when loading an
	// image file.
	ImageLoadingError = 2

	// ImageSavingError declares the exit code for failures when saving the
	// processed image into a file.
	ImageSavingError = 3

	// FilterNotFoundError declares the exit code for failing to retrieve the
	// filter specified by the user.
	FilterNotFoundError = 4

	// InvalidExecutionModeError declares the exit code for providing an
	// unknown execution mode (i.e., not "serial" nor "concurrent").
	InvalidExecutionModeError = 5
	
	// OutputDirError declares the exit code for failures when checking the
	// existence of the specified output directory or creating it if needed.
	OutputDirError = 6
)

// terminateWithError prints the error and exits the program with the given code.
func terminateWithError(err error, code int) {
	fmt.Println(err)
	os.Exit(code)
}
