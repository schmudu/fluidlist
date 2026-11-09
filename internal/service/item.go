package service

import (
	"fmt"

	"github.com/schmudu/fluidlist/internal/common"
)

func AddItem(title string) {
	fmt.Printf("Adding item: %s...\n", title)
	common.WriteFileBytes()
}
