package stringer

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "dev"
var commit = "HEAD"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Axol",
	Long:  `All software has versions. This is Axol's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Axol Static Site Generator v%s -- %s\n", version, commit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
