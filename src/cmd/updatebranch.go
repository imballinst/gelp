package cmd

import (
	"fmt"

	"github.com/fatih/color"
	helpers "github.com/imballinst/gelp/src/helpers"
	"github.com/spf13/cobra"
)

var updateBranchRemote string

// Root command.
var updateBranchCmd = &cobra.Command{
	Use:   "updatebranch",
	Short: "Updatebranch is used to update the current branch, or other branches.",
	Long: `When updating a branch using Git, we have some options, 2 of them are "fetch" and "pull".
When you want to update your current branch instead, perhaps you would want to use "git pull" instead
since it is shorter than "gelp updatebranch". However, when updating other branches, this command
can be handy.`,
	Example: fmt.Sprintf(`1) Update current branch (developer's advice: prefer "git pull" instead)
   %s

2) Update the "dev" branch to be updated from the remote "origin"
   %s
	 
3) Update the "dev" branch to be updated from the remote "upstream"
   %s`,
		color.BlueString("gelp updatebranch"),
		color.BlueString("gelp updatebranch dev"),
		color.BlueString("gelp updatebranch dev --remote upstream")),
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.UpdateBranch(updateBranchRemote, args[0])
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	updateBranchCmd.Flags().StringVarP(&updateBranchRemote, "remote", "r", "origin", "The base remote used for the update. Defaults to origin.")
}
