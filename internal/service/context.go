package service

import (
	"github.com/schmudu/fluidlist/internal/common"
)

const (
	ContextFile = "context"
)

func GetContext() (string, error) {
	workspaceDir := GetWorkspaceDir()
	content, error := common.ReadFile(workspaceDir + "/" + ContextFile)
	if error != nil {
		return "", error
	}
	return content[0], nil
}
