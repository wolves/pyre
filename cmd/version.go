/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/spf13/cobra"
)

var PyreVersion string

func getPyreVersion() string {
	noVersionAvail := "No version info available for this build, run 'pyre help version' for additional info"

	if len(PyreVersion) != 0 {
		return PyreVersion
	}

	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return noVersionAvail
	}

	if len(buildInfo.Main.Version) != 0 {
		if buildInfo.Main.Version != "(devel)" {
			return buildInfo.Main.Version
		}
	}

	var vcsRevision string
	var vcsTime time.Time
	for _, setting := range buildInfo.Settings {
		switch setting.Key {
		case "vcs.revision":
			vcsRevision = setting.Value
		case "vcs.time":
			vcsTime, _ = time.Parse(time.RFC3339, setting.Value)
		}
	}

	if vcsRevision != "" {
		return fmt.Sprintf("%s, (%s)", vcsRevision, vcsTime)
	}

	return noVersionAvail
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display application version information",
	Long: `
The version command provides information about the application's version.

Pyre requires version information to be embedded at compile time.
For detailed version information, Pyre needs to be built as specified in the README installation instructions.
If Pyre is built within a version control repository and other version info isn't available,
the revision hash will be used instead.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		version := getPyreVersion()
		fmt.Printf("Pyre CLI version: %v\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
