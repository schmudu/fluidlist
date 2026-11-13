/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/schmudu/fluidlist/internal/service"
	"github.com/spf13/cobra"
)

// itemsCmd represents the items command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List items in current context",
	Long:  `This command lists items in current context along with their hashes and dates.`,
	Run: func(cmd *cobra.Command, args []string) {
		date, err := cmd.Flags().GetBool("date")
		if err != nil {
			return
		}
		service.ListItems(date)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	lsCmd.PersistentFlags().BoolP("date", "d", false, "Show date")
}
