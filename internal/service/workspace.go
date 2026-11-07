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

func InitializeWorkspace() {
	workspaceDir := getWorkspaceDir()
	fmt.Printf("INFO: Initializing workspace: %s\n", workspaceDir)
	initializeDirs(workspaceDir)
}

func initializeDirs(workspaceDir string) {
	// create lists dir
	errors := common.CreateFolder(listsPath(workspaceDir))
	if errors != nil {
		fmt.Printf("ERROR: Unabled to create directory: %s\n", listsPath(workspaceDir))
	}
	errors = common.CreateFolder(itemsPath(workspaceDir))
	if errors != nil {
		fmt.Printf("ERROR: Unabled to create directory: %s\n", listsPath(workspaceDir))
	}
	fmt.Printf("INFO: Workspace folders created and initialized: %s\n", workspaceDir)
}

func validateDirLists(workspaceDir string) {
	ok := common.FolderExists(listsPath(workspaceDir))
	if !ok {
		fmt.Printf("ERROR: Lists directory '%s' does not exist.  Run '%s' and try again.\n", listsPath(workspaceDir), "fl init")
		os.Exit(1)
	}

	ok = common.FolderExists(itemsPath(workspaceDir))
	if !ok {
		fmt.Printf("ERROR: Items directory '%s' does not exist.  Run '%s' and try again.\n", itemsPath(workspaceDir), "fl init")
		os.Exit(1)
	}
}

func getWorkspaceDir() string {
	return os.Getenv("HOME") + "/.fluidlist"
}

func itemsPath(workspaceDir string) string {
	return workspaceDir + "/" + DIR_ITEMS
}

func listsPath(workspaceDir string) string {
	return workspaceDir + "/" + DIR_LISTS
}
