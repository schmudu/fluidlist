package service

import (
	"errors"
	"fmt"
	"os"
)

func ValidateConfig() {
	fmt.Println("INFO: Validating config")
	validateDirLists()
}

func validateDirLists() {
	dirPath := "./patrick"

	info, err := os.Stat(dirPath)
	if err == nil {
		// The path exists, now check if it's actually a directory
		if info.IsDir() {
			fmt.Println("Directory exists.")
		} else {
			fmt.Println("Path exists, but it's a file, not a directory.")
		}
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Directory does not exist.")
	} else {
		// The directory might exist, but we can't read it (e.g., permission denied)
		fmt.Printf("Error checking directory: %v\n", err)
	}
}
