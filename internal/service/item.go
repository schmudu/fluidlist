package service

import (
	"fmt"
	"os"
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
		insertIntoFile(listPath, checksum, itemDate)
	} else {
		common.AppendToFile(listPath, checksum)
	}
}

func insertIntoFile(path string, checksum string, itemDate time.Time) error {
	fmt.Printf("Need to insert %s into file\n", itemDate)
	checksumList, error := common.ReadFile(path)
	if error != nil {
		return error
	}

	for index, item := range checksumList {
		fmt.Printf("===i:%d hash:%s\n", index, item)
	}

	insertIndex, error := finditemIndex(itemDate, checksumList, 0, len(checksumList)-1)
	fmt.Printf("===insert index:%d newLine:%s\n", insertIndex, checksum)
	insertLine(path, insertIndex, checksum)

	return nil
}

// insertLine inserts text at a 0-based line index.
func insertLine(filePath string, lineIndex int, newLine string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")

	if lineIndex < 0 || lineIndex > len(lines) {
		return fmt.Errorf("line index %d out of range (0-%d)", lineIndex, len(lines))
	}

	// Insert the new line at the target index
	lines = slices.Insert(lines, lineIndex, newLine)

	output := strings.Join(lines, "\n")
	return os.WriteFile(filePath, []byte(output), 0644)
}

func finditemIndex(itemDate time.Time, checksumList []string, startIndex int, endIndex int) (int, error) {
	// uses binary insertion algorithm to find index to insert item
	// the insertion index is based on the item creation date
	// if checksumList is empty, then return first spot
	if len(checksumList) == 0 {
		return 0, nil
	}

	nextPotentialIndexMatch := (startIndex + endIndex) / 2

	nextPotentialDate := getItemDateFromHash(checksumList[nextPotentialIndexMatch])
	fmt.Printf("====potential index: %d date:%s\n", nextPotentialIndexMatch, nextPotentialDate)
	if nextPotentialIndexMatch == startIndex {
		if itemDate.After(nextPotentialDate) {
			return startIndex + 1, nil
		} else {
			return startIndex, nil
		}
	} else if itemDate.After(nextPotentialDate) {
		// if potential index matches start index, then return
		return finditemIndex(itemDate, checksumList, nextPotentialIndexMatch, endIndex)
	} else {
		return finditemIndex(itemDate, checksumList, startIndex, nextPotentialIndexMatch)
	}
}

func getItemDateFromHash(hash string) time.Time {
	item, _ := getItemObject(hash)
	return item.date
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
