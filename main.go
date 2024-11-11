package main

import (
	"embed"

	"github.com/a1994sc/axol/cmd"
)

//go:embed schema/list.schema.json
var listSchema embed.FS

func main() {
	cmd.ListSchema = listSchema
	cmd.Execute()
}
