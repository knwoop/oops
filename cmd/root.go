package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oops",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
    examples and usage of using your application. For example:
    Cobra is a CLI library for Go that empowers applications.
    This application is a tool to generate the needed files
    to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello, 世界!!")
	},
}

func init() {
	cobra.EnableCommandSorting = false

	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true
}

func Execute(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}
