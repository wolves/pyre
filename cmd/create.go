/*
Copyright © 2024 Christopher Stingl <cs@wlvs.io>
*/
package cmd

import (
	"fmt"
	"strings"

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

		projects := viper.Get("projects")

		fmt.Println("Projs", projects)

		c := feature.Component{
			SunbirdDir: getSunbirdDir(),
			Filename:   args[0],
			Name:       kebabToTitle(args[0]),
		}

		noTests, err := cmd.Flags().GetBool("no-tests")
		cobra.CheckErr(err)

		err = c.Create(noTests)
		cobra.CheckErr(err)
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
	createCmd.Flags().BoolP("no-tests", "x", false, "Create files without test specs")
}

func kebabToTitle(s string) string {
	words := strings.Split(s, "-")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}
