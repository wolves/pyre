/*
Copyright © 2024 Christopher Stingl <cs@wlvs.io>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wolves/pyre/cmd/feature"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Sunbird project feature",
	Long: `Create a Sunbird project feature

Generates the necessary Angular files for a default, bare-bones
project feature component. This includes the component files along with
specs, styles, and ngrx state management files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called", args)

		sbDir := viper.GetString("sunbird_dir")
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		f := feature.Component{
			ProjectPath: filepath.Join(home, sbDir),
			Name:        args[0],
		}

		fmt.Printf("Component %+v", f)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	featureCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
