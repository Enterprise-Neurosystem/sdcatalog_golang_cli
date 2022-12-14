/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package list

import (

	"github.com/spf13/cobra"
)

// catalogueCmd represents the catalogue command
var catalogueCmd = &cobra.Command{
	Use:   "catalogue",
	Short: "list the entries of the provided catalogue name",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	ListCmd.AddCommand(catalogueCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catalogueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catalogueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
