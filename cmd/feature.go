/*
Copyright © 2024 Christopher Stingl <cs@wlvs.io>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// featureCmd represents the feature command
var featureCmd = &cobra.Command{
	Use:   "feature <command>",
	Short: "Manage Sunbird features",
	Long:  `Work with Sunbird project features`,
}

func init() {
	rootCmd.AddCommand(featureCmd)
}
