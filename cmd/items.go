/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/schmudu/fluidlist/internal/service"
	"github.com/spf13/cobra"
)

// itemsCmd represents the items command
var itemsCmd = &cobra.Command{
	Use:   "items",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		date, err := cmd.Flags().GetBool("date")
		if err != nil {
			return
		}
		service.ListItems(date)
	},
}

func init() {
	rootCmd.AddCommand(itemsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// itemsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// itemsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	itemsCmd.PersistentFlags().BoolP("date", "d", false, "Show date")
}
