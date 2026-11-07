/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/schmudu/fluidlist/internal/service"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add item",
	Long:  `Add item to current list.`,
	Run: func(cmd *cobra.Command, args []string) {
		service.ValidateWorkspace()
		title, _ := cmd.Flags().GetString("title")
		service.AddItem(title)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//addCmd.Flags().BoolP("title", "t", false, "Item title")
	addCmd.Flags().StringP("title", "t", "", "Title of item")
}
