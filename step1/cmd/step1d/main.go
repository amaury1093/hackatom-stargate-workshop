package main

import (
	"os"

	"github.com/amaurymartiny/step1/cmd/step1d/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
