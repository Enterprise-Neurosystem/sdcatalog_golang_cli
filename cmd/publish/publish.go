/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package publish

import (
	"github.com/spf13/cobra"
)

// PublishCmd represents the publish command
var PublishCmd = &cobra.Command{
	Use:   "publish",
	Short: "publish is a palette that contains publishing based commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
