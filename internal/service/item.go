package service

import (
	"fmt"
	"strings"

	"github.com/schmudu/fluidlist/internal/common"
)

const (
	ItemsDelimiter = ","
)

func ListItems() {
	fmt.Println("ListItems")
	listName, error := GetContext()
	if error != nil {
		fmt.Printf("Error getting context file: %s\n", error.Error())
		return
	}
	var items []string
	items, error = getItems(listName)
	fmt.Printf("ListName: %s item:%s error:%s\n", listName, items, error)
}

func getItems(listName string) ([]string, error) {
	content, error := common.GetFileContent(GetListsPath() + "/" + listName)
	if error != nil {
		return nil, error
	}
	fmt.Println("content:", string(content))
	fmt.Println("==========")
	for _, item := range strings.Split(string(content), "\n") {
		if item != "" {
			fmt.Println(item)
		}
	}
	fmt.Println("==========")
	return strings.Split(string(content), "\n"), nil
}

func AddItem(title []string) {
	checksum := common.GetChecksum([]byte(common.CurrentTimeEpoch()))
	folder, file := getItemPath(checksum)

	folderPath := GetItemsPath() + "/" + folder
	common.CreateFolder(folderPath)

	titleBytes := []byte(strings.Join(title, " "))
	date := common.CurrentTimeDateOnly()
	completed := getDefaultCompletion()
	common.WriteFileBytes(folderPath+"/"+file, append([]byte(date+ItemsDelimiter+completed+ItemsDelimiter), titleBytes...))

	// write checksum to list
	listPath := GetListsPath() + "/" + "default"
	common.AppendToFile(listPath, []byte(checksum))
}

func getDefaultCompletion() string {
	return "false"
}

func getItemPath(checksum string) (string, string) {
	return checksum[0:4], checksum[4:]
}
