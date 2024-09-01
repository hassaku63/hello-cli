/*
Copyright © 2024 hassaku63
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name string
)

// subcmd1Cmd represents the subcmd1 command
var subcmd1Cmd = &cobra.Command{
	Use:   "subcmd1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("subcmd1 called")
	},
}

func init() {
	rootCmd.AddCommand(subcmd1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subcmd1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	name := *subcmd1Cmd.Flags().StringP("name", "n", "a", "Name of the user")

	if name == "" {
		panic("name is required")
	} else {
		fmt.Println("Hello, " + name)
	}
}
