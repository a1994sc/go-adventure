package main

import (
	"embed"

	"github.com/a1994sc/axol/cmd/stringer"
)

//go:embed schema/list.schema.json
var listSchema embed.FS

func main() {
	stringer.ListSchema = listSchema
	stringer.Execute()
}
