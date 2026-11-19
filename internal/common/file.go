package common

import (
	"bufio"
	"errors"
	"log"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func GetChecksumFilepath(checksum string) {

}

func WriteFileBytes(path string, content []byte) error {
	err := os.WriteFile(path, content, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %s", err)
	}
	return err
}

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return !info.IsDir()
}

func AppendToFile(path string, content string) error {
	// Open the file in append mode, creating it if it doesn't exist
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Append text to the file
	if _, err := f.WriteString(content + "\n"); err != nil {
		log.Fatal(err)
	}
	return nil
}

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func CreateFolder(path string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func FolderExists(path string) bool {
	info, err := os.Stat(path)
	if err == nil {
		// The path exists, now check if it's actually a directory
		if info.IsDir() {
			return true
		}
	}
	return false
}
