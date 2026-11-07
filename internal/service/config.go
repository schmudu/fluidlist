package service

import (
	"fmt"
	"github.com/schmudu/fluidlist/internal/common"
	"os"
)

const (
	DIR_LISTS = "lists"
	DIR_ITEMS = "items"
)

func ValidateWorkspace() {
	fmt.Println("INFO: Validating workspace")
	workspaceDir := getWorkspaceDir()
	validateDirLists(workspaceDir)
}

func validateDirLists(workspaceDir string) {
	listsPath := workspaceDir + "/" + DIR_LISTS
	exists := common.FolderExists(listsPath)
	if !exists {
		fmt.Printf("ERROR: Lists directory '%s' does not exist.  Run '%s' and try again.\n", listsPath, "fl init")
		os.Exit(1)
	}

	dirPath := workspaceDir + "/" + DIR_ITEMS
	exists = common.FolderExists(dirPath)
	if !exists {
		fmt.Printf("ERROR: Items directory '%s' does not exist.  Run '%s' and try again.\n", dirPath, "fl init")
		os.Exit(1)
	}
}

func getWorkspaceDir() string {
	return os.Getenv("HOME")
}
