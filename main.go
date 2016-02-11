package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/pdxjohnny/s-db/commands"
	"github.com/pdxjohnny/s-db/db/variables"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: variables.ServiceName}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
