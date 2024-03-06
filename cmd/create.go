/*
Copyright © 2024 Christopher Stingl <cs@wlvs.io>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wolves/pyre/cmd/feature"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Sunbird project feature",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("cannot create: feature name argument required (ex: some-feature-name)")
		}

		if len(args) > 1 {
			return errors.New("cannot create: too many arguments. Only a single feature name is allowed (ex: some-feature-name)")
		}

		return nil
	},
	Long: `Create a Sunbird project feature

Generates the necessary Angular files for a default, bare-bones
project feature component. This includes the component files along with
specs, styles, and ngrx state management files.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c feature.Component

		cPath, err := cmd.Flags().GetString("path")
		cobra.CheckErr(err)

		if _, err := os.Stat(cPath); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Project directory %s does not exist.\n", cPath)
			os.Exit(1)
		}

		fmt.Printf("Path flag value: %s\n", cPath)

		c = feature.Component{
			ProjPath: cPath,
			Filename: args[0],
			Name:     kebabToTitle(args[0]),
		}

		noTests, err := cmd.Flags().GetBool("no-tests")
		cobra.CheckErr(err)

		err = c.Create(noTests)
		cobra.CheckErr(err)

		return nil
	},
}

func init() {
	// cobra.OnInitialize(initConfig)
	featureCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createCmd.Flags().StringP("path", "p", "", "Path to the directory where the feature component should be created (Required)")
	createCmd.Flags().BoolP("no-tests", "x", false, "Create files without test specs")

	createCmd.MarkFlagRequired("path")
}

func kebabToTitle(s string) string {
	words := strings.Split(s, "-")
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}
