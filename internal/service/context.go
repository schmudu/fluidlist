package service

import (
	"github.com/schmudu/fluidlist/internal/common"
)

const (
	ContextFile = "context"
)

func GetContext() (string, error) {
	workspaceDir := GetWorkspaceDir()
	content, error := common.GetFileContent(workspaceDir + "/" + ContextFile)
	if error != nil {
		return "", error
	}
	return string(content), nil
}
