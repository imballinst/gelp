package cmd

import (
	"fmt"
	"os"

	helpers "github.com/imballinst/gelp/src/helpers"
	"github.com/spf13/cobra"
)

var ShowVersion bool

// Root command.
var rootCmd = &cobra.Command{
	Use:   "gelp",
	Short: "Gelp is a Git CLI helper",
	Long: `Committed changes in the wrong branch? Want to clean-up branch post-merge? 
Then, gelp might be a good tool for these use-cases!`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}

		// Allow 0 arguments.
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Hello!")
	},
}

func init() {
	// Root flags.
	rootCmd.Flags().BoolVarP(&ShowVersion, "version", "v", false, "Show current gelp version")
	rootCmd.Version = helpers.GetVersion()

	// Add other commands.
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(squashToCmd)
	rootCmd.AddCommand(postMergeCmd)
	rootCmd.AddCommand(freshCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
