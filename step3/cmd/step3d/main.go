package main

import (
	"os"

	"github.com/amaurymartiny/step3/cmd/step3d/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := cmd.Execute(rootCmd); err != nil {
		os.Exit(1)
	}
}
