/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// isActiveCmd represents the isActive command
var isActiveCmd = &cobra.Command{
	Use:   "isActive",
	Short: "Checks is a listing is active",
	Long: `Take a URL, parases the HTML to assess is the listing is still active.jA longer description that spans multiple lines and likely contains examples
	it will err on the side of causion, and mark a listing a unavalible if it cannot assess if it is live.
	I.e. there can be false negatives.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("isActive called")
	},
}

func init() {
	rootCmd.AddCommand(isActiveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// isActiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// isActiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
