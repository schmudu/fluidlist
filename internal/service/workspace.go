package service

import (
	"fmt"
	"os"

	"github.com/schmudu/fluidlist/internal/common"
)

const (
	DirLists = "lists"
	DirItems = "items"
)

func ValidateWorkspace() {
	validateDirLists(getWorkspaceDir())
	validateDirItems(getWorkspaceDir())
}

func InitializeWorkspace() {
	workspaceDir := getWorkspaceDir()
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
	currentTime := common.CurrentTimeEpoch()
	checksum := common.GetChecksum([]byte(currentTime))
	fmt.Printf("Checksum: %s time: %s epoch:%s\n", checksum, currentTime, common.CurrentTimeEpoch())

	// write file default list in lists folder
	common.WriteFileBytes(folderPath+"/default", nil)
	// write checksum in .context file in .fluidlists dir
}

func initializeDirs(workspaceDir string) {
	// create root dir
	initializeDir(workspaceDir)

	// create lists and items
	var dirCreated = initializeDir(listsPath(workspaceDir))
	createDefaultList(listsPath(workspaceDir))

	if dirCreated {
		initializeDir(itemsPath(workspaceDir))
	}

	if dirCreated {
		fmt.Printf("INFO: Workspace folders created and initialized: %s\n", workspaceDir)
	} else {
		fmt.Printf("INFO: Workspace folders already initialized at location: %s\n", workspaceDir)
	}
}

func validateDirLists(workspaceDir string) {
	ok := common.FolderExists(listsPath(workspaceDir))
	if !ok {
		fmt.Printf("ERROR: Lists directory '%s' does not exist.  Run '%s' and try again.\n", listsPath(workspaceDir), "fl init")
		os.Exit(1)
	}
}

func validateDirItems(workspaceDir string) {
	ok := common.FolderExists(itemsPath(workspaceDir))
	if !ok {
		fmt.Printf("ERROR: Items directory '%s' does not exist.  Run '%s' and try again.\n", itemsPath(workspaceDir), "fl init")
		os.Exit(1)
	}
}

func getWorkspaceDir() string {
	return os.Getenv("HOME") + "/.fluidlist"
}

func itemsPath(workspaceDir string) string {
	return workspaceDir + "/" + DirItems
}

func listsPath(workspaceDir string) string {
	return workspaceDir + "/" + DirLists
}
