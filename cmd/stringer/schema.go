package stringer

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	yaml "github.com/goccy/go-yaml"
	"github.com/santhosh-tekuri/jsonschema"
	"github.com/spf13/cobra"
)

var ListSchema fs.ReadFileFS

var schemaCmd = &cobra.Command{
	Use:     "schema",
	Aliases: []string{"sch"},
	Short:   "Reverses a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		content, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatal(err)
		}

		var m interface{}
		err = yaml.Unmarshal([]byte(content), &m)

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
		fmt.Println("Yaml is valid")
	},
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}
