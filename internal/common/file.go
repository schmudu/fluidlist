package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func WriteFile() {
	data := User{Name: "Alice Smith", Email: "alice@example.com", Age: 30}

	// 1. Create or open the file
	file, err := os.Create("user.json")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}
	}(file)

	// 2. Initialize the encoder and write directly to the file
	encoder := json.NewEncoder(file)

	// Optional: Pretty-print the JSON output
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		return
	}

	fmt.Println("JSON file written successfully using NewEncoder!")
}

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return !info.IsDir()
}

func CreateFolder(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating folder: %v\n", err)
		return err
	}
	return nil
}

func FolderExists(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		// The path exists, now check if it's actually a directory
		if info.IsDir() {
			fmt.Println("Directory exists.")
			return true
		}
	}
	return false
}
