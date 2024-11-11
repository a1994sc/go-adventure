package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/a1994sc/axol/pkg/transform"
	"github.com/goccy/go-yaml"
	"github.com/santhosh-tekuri/jsonschema"
	"github.com/spf13/cobra"
)

var ListSchema fs.ReadFileFS

var schemaCmd = &cobra.Command{
	Use:     "schema [ file ]",
	Aliases: []string{"sch"},
	Short:   "Validate a schema",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		content, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatal(err)
		}

		docs, _ := transform.SplitYAML(content)

		for _, doc := range docs {
			var m interface{}
			err = yaml.Unmarshal([]byte(doc), &m)
			if err != nil {
				panic(err)
			}
			data, err := ListSchema.ReadFile("schema/list.schema.json")
			if err != nil {
				log.Fatal(err)
			}
			compiler := jsonschema.NewCompiler()
			if err := compiler.AddResource("schema/list.schema.json", strings.NewReader(string(data))); err != nil {
				panic(err)
			}
			schema, err := compiler.Compile("schema/list.schema.json")
			if err != nil {
				log.Fatal(err)
			}
			if err := schema.ValidateInterface(m); err != nil {
				log.Fatal(err)
			}
			fmt.Println(m)
		}
	},
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}
