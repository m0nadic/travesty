package main

import (
	"fmt"
	"os"
	"travesty/cmd/commands"
)

func main() {
	err := commands.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
