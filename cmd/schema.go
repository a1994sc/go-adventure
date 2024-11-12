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

type Document struct {
	data interface{}
}

func (d *Document) addField(key, value string) {
	v, _ := d.data.(map[string]interface{})
	v[key] = value
	d.data = v
}

func (d *Document) containsField(key string) bool {
	v, _ := d.data.(map[string]interface{})
	return v[key] != nil
}

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

		d := &Document{}

		docs, _ := transform.SplitYAML(content)

		for _, doc := range docs {
			err = yaml.Unmarshal([]byte(doc), &d.data)
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

			if !d.containsField("x-tra") {
				d.addField("x-tra", "blah")
			}

			fmt.Println(d.data)

			schema, err := compiler.Compile("schema/list.schema.json")
			if err != nil {
				log.Fatal(err)
			}

			if err := schema.ValidateInterface(d.data); err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}
