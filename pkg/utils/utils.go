package utils

import (
	"fmt"
	"os"
)

// CheckError handles error checking and exit if error occurs
func CheckError(err error, message string) {
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
		if err != nil {
			fmt.Printf("Error encountered: %v", err)
		}
		os.Exit(1)
	}
}
