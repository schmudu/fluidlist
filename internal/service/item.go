package service

import (
	"strings"

	"github.com/schmudu/fluidlist/internal/common"
)

const (
	ItemsDelimiter = ","
)

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
