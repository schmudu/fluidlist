package service

import (
	"strings"

	"github.com/schmudu/fluidlist/internal/common"
)

func AddItem(title []string) {
	checksum := common.GetChecksum([]byte(common.CurrentTimeEpoch()))
	folder, file := getItemPath(checksum)

	folderPath := GetItemsPath() + "/" + folder
	common.CreateFolder(folderPath)

	byteSlice := []byte(strings.Join(title, " "))
	common.WriteFileBytes(folderPath+"/"+file, byteSlice)

	// write checksum to list
}

func getItemPath(checksum string) (string, string) {
	return checksum[0:5], checksum[5:]
}
