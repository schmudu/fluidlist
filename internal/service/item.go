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

	// parse time
	itemDate, insertionRequired, error := common.ParseTime(date)
	if error != nil {
		fmt.Printf("ERROR: Unable to insert item due to error: %s\n", error)
		return
	}

	// completion
	completed := getDefaultCompletion()

	folderPath := GetItemsPath() + "/" + folder
	common.CreateFolder(folderPath)

	// write file
	common.WriteFileBytes(folderPath+"/"+file, append([]byte(itemDate.Format(time.DateOnly)+ItemsDelimiter+completed+ItemsDelimiter), []byte(title)...))

	// get path to list
	listPath := GetListsPath() + "/" + "default"
	if insertionRequired {
		common.InsertIntoFile(listPath, []byte(checksum), itemDate)
	} else {
		common.AppendToFile(listPath, []byte(checksum))
	}
}

func ListItems(date bool) {
	listName, error := GetContext()
	if error != nil {
		fmt.Printf("Error getting context file: %s\n", error.Error())
		return
	}
	printItems(listName, date)
}

func getItemHashes(listName string) ([]string, error) {
	return common.ReadFile(GetListsPath() + "/" + listName)
}

func printItems(listName string, date bool) {
	hashes, error := getItemHashes(listName)
	if error != nil {
		return
	}
	printItemObjects(hashes, date)
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
	content, error := common.ReadFile(path)
	if error != nil {
		return Item{}, error
	}
	item := buildItem(hash, content[0])
	return item, nil
}

func buildItem(hash string, content string) Item {
	data := strings.Split(content, ",")
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
