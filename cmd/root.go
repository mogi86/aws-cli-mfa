package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "mogi-awscli",
		Short: "\"mogi-awscli\" can set AWS ENV",
		Long: `"mogi-awscli" can set AWS ENV.
                This command line tool cover the ENV below now.
                  - aws_access_key_id
                  - aws_secret_access_key
                  - aws_session_token
                These ENV is used for "aws cli" command by using MFA.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("this is root command.")
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}