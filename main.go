package main

import (
	"fmt"

	yaml "github.com/goccy/go-yaml"
)

func main() {
	var v struct {
		A int
		B string
	}
	v.A = 1
	v.B = "hello"
	bytes, err := yaml.Marshal(v)
	if err != nil {
		//...
	}
	fmt.Println(string(bytes)) // "a: 1\nb: hello\n"
}
