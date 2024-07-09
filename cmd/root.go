package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "simple CLI tool",
	Long:  "myapp is a simple CLI tool to manage your application",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

}
