package main

import (
	"github.com/Guilherme-De-Marchi/dynamic-plugin/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
