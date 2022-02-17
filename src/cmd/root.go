package gelp

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gelp",
	Short: "Gelp is a Git CLI helper",
	Long: `Committed changes in the wrong branch? Want to clean-up branch post-merge? 
Then, gelp might be a good tool for these use-cases!`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Run(versionCmd, args)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
