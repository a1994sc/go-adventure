package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
)

var findCmd = &cobra.Command{
	Use:     "find",
	Aliases: []string{"find"},
	Short:   "",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var path = setBaseDirectory(args)
		pattern := `\.ya?ml$`

		fileRegEx, e := regexp.Compile(pattern)
		if e != nil {
			log.Fatal(e)
		}

		err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
			if err == nil && fileRegEx.MatchString(d.Name()) {
				println(path)
			}

			return nil
		})

		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(findCmd)
}
