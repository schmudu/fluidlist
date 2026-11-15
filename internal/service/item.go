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
	HashMinLength  = 4
)

func RemoveItems(hashes []string) {
	for _, hash := range hashes {
		deleteItem(hash)
	}
}

func deleteItem(hash string) {
	// STUB
	/*
		itemFolder, itemFile := getItemPath(hash)
		itemPath := itemFolder + "/" + itemFile
		if !common.FileExists(itemPath) {
			fmt.Printf("Item %s not found.\n", itemPath)
		} else {
			fmt.Printf("Removing item %s ...\n", itemPath)
		}
	*/
}

func AddItem(title string, date string) {
	checksum := common.GetChecksum([]byte(common.CurrentTimeEpoch()))
	folder, file := getItemPath(checksum)

	//titleBytes := []byte(strings.Join(title, " "))
	//date := common.CurrentTimeDateOnly()
	// parse time
	itemDate, error := common.ParseTime(date)
	if error != nil {
		fmt.Printf("Error parsing date %s\n", date)
		return
	}

	// completion
	completed := getDefaultCompletion()

	folderPath := GetItemsPath() + "/" + folder
	common.CreateFolder(folderPath)

	// write file
	common.WriteFileBytes(folderPath+"/"+file, append([]byte(itemDate+ItemsDelimiter+completed+ItemsDelimiter), []byte(title)...))

	// write checksum to list
	listPath := GetListsPath() + "/" + "default"
	common.AppendToFile(listPath, []byte(checksum))
}

func ListItems(date bool) {
	listName, error := GetContext()
	if error != nil {
		fmt.Printf("Error getting context file: %s\n", error.Error())
		return
	}
	printItems(listName, date)
}

func getItemHashes(listName string) ([]byte, error) {
	return common.GetFileContent(GetListsPath() + "/" + listName)
}

func printItems(listName string, date bool) {
	hashes, error := getItemHashes(listName)
	if error != nil {
		return
	}
	printItemObjects(strings.Split(string(hashes), "\n"), date)
}

func printItemObjects(hashes []string, date bool) {
	for _, hash := range slices.Backward(hashes) {
		if hash != "" {
			item, _ := getItemObject(hash)
			if date {
				timeLayout := "Mon Jan 2 2006"
				fmt.Printf("%s%s%s [%s] %s\n", common.Green, item.hash, common.Reset, item.date.Format(timeLayout), item.description)
			} else {
				fmt.Printf("%s %s\n", item.hash, item.description)
			}
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

func getDefaultCompletion() string {
	return "false"
}

func getItemPath(checksum string) (string, string) {
	if len(checksum) < HashMinLength {
		return "", ""
	}
	return checksum[0:HashMinLength], checksum[HashMinLength:]
}
