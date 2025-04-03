package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// result
	// cmd运行
	rootCmd = &cobra.Command{
		Use:   "run",
		Short: "Submit links to search engines",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {

			return
		},
	}
)

func Execute() {

	rootCmd.AddCommand(cmdRun)
	rootCmd.AddCommand(httpRun)
	rootCmd.Execute()

}
