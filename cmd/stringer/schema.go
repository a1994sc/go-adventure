package stringer

import (
	"fmt"
	"io/fs"
	"log"
	"strings"

	yaml "github.com/goccy/go-yaml"
	"github.com/santhosh-tekuri/jsonschema"
	"github.com/spf13/cobra"
)

var yamlText = `
- one
- two
`

var schemaText = `
{
	"type": "array",
	"items": {
		"type": "string"
	}
}
`

var ListSchema fs.ReadFileFS

var schemaCmd = &cobra.Command{
	Use:     "schema",
	Aliases: []string{"sch"},
	Short:   "Reverses a string",
	Run: func(cmd *cobra.Command, args []string) {
		var m interface{}
		err := yaml.Unmarshal([]byte(yamlText), &m)

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
		fmt.Println(schemaText)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}
