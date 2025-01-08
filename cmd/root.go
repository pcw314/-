package cmd

import (
	"gitee.com/xygfm/authorization/cmd/start"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "authorization-api",
	Short: "授权系统API",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(start.Cmd)
}

func Execute() error {
	return rootCmd.Execute()
}
