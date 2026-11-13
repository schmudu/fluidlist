package service

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/schmudu/fluidlist/internal/common"
)

type Item struct {
	hash        string
	completed   bool
	date        time.Time
	description string
}

const (
	ItemsDelimiter = ","
)

func ListItems() {
	listName, error := GetContext()
	if error != nil {
		fmt.Printf("Error getting context file: %s\n", error.Error())
		return
	}
	printItems(listName)
}

func getItemHashes(listName string) ([]byte, error) {
	return common.GetFileContent(GetListsPath() + "/" + listName)
}

func printItems(listName string) {
	hashes, error := getItemHashes(listName)
	if error != nil {
		return
	}
	printItemObjects(strings.Split(string(hashes), "\n"))
}

func printItemObjects(hashes []string) {
	for _, hash := range slices.Backward(hashes) {
		if hash != "" {
			item, _ := getItemObject(hash)
			fmt.Printf("[%s] %s\n", item.hash[0:9], item.description)
		}
	}
}

func getItemObject(hash string) (Item, error) {
	itemPathFolder, itemPathFile := getItemPath(hash)
	path := GetItemsPath() + "/" + itemPathFolder + "/" + itemPathFile
	content, error := common.GetFileContent(path)
	if error != nil {
		return Item{}, error
	}
	item := buildItem(hash, content)
	return item, nil
}

func buildItem(hash string, content []byte) Item {
	data := strings.Split(string(content), ",")
	date, _ := time.Parse(time.DateOnly, data[0])
	completed := bool(data[1] == "true")
	return Item{hash: hash, completed: completed, date: date, description: data[2]}
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
