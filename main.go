package main

import (
	"fmt"

	"github.com/livanjimenez/task-tracker-cli/cmd"
)

func main() {
    rootCmd := cmd.RootCmd()

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}