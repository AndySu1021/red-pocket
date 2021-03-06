package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"red-packet/cmd"
	"runtime"
)

var rootCmd = &cobra.Command{Use: "server migrate"}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	rootCmd.AddCommand(cmd.ServerCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
