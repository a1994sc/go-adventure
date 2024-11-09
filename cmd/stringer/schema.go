package stringer

import (
	"fmt"
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
		compiler := jsonschema.NewCompiler()
		if err := compiler.AddResource("schema.json", strings.NewReader(schemaText)); err != nil {
			panic(err)
		}
		schema, err := compiler.Compile("schema.json")
		if err != nil {
			panic(err)
		}
		if err := schema.ValidateInterface(m); err != nil {
			panic(err)
		}
		fmt.Println(schemaText)
	},
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}
