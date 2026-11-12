package service

import (
	"fmt"
	"os"

	"github.com/schmudu/fluidlist/internal/common"
)

const (
	DirLists    = "lists"
	DirItems    = "items"
	DefaultList = "default"
)

func GetWorkspaceDir() string {
	return os.Getenv("HOME") + "/.fluidlist"
}

func GetItemsPath() string {
	return GetWorkspaceDir() + "/" + DirItems
}

func GetListsPath() string {
	return GetWorkspaceDir() + "/" + DirLists
}

func ValidateWorkspace() {
	validateDirLists()
	validateDirItems()
}

func InitializeWorkspace() {
	workspaceDir := GetWorkspaceDir()
	initializeDirs(workspaceDir)
}

func initializeDir(folderPath string) bool {
	var dirCreated = false
	errors := common.CreateFolder(folderPath)
	if errors == nil {
		dirCreated = true
	}
	return dirCreated
}

func createDefaultList(folderPath string) {
	// write file default list in lists folder
	common.WriteFileBytes(folderPath+"/"+DefaultList, nil)
}

func initializeDirs(workspaceDir string) {
	// create root dir
	initializeDir(workspaceDir)

	// create context file - tracks current list
	common.WriteFileBytes(workspaceDir+"/"+ContextFile, []byte(DefaultList))

	// create lists dir
	var dirCreated = initializeDir(GetListsPath())
	createDefaultList(GetListsPath())

	if dirCreated {
		// create items dir
		initializeDir(GetItemsPath())
	}

	if dirCreated {
		fmt.Printf("INFO: Workspace folders created and initialized: %s\n", workspaceDir)
	} else {
		fmt.Printf("INFO: Workspace folders already initialized at location: %s\n", workspaceDir)
	}
}

func validateDirLists() {
	ok := common.FolderExists(GetListsPath())
	if !ok {
		fmt.Printf("ERROR: Lists directory '%s' does not exist.  Run '%s' and try again.\n", GetListsPath(), "fl init")
		os.Exit(1)
	}
}

func validateDirItems() {
	ok := common.FolderExists(GetItemsPath())
	if !ok {
		fmt.Printf("ERROR: Items directory '%s' does not exist.  Run '%s' and try again.\n", GetItemsPath(), "fl init")
		os.Exit(1)
	}
}
