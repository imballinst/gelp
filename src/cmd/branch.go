package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bndr/gotabulate"
	"github.com/fatih/color"
	helpers "github.com/imballinst/gelp/src/helpers"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// Root command.
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Lists the local barnches, as well as their last commit and last commit datetime.",
	Long: `Just got back from vacation? Took a day of time-off in the middle of the week? When we get back to work,
perhaps we can forget what we were working on previously. The normal "git branch" don't always help, especially
if we changed the branch before we took our leave.

What this command does is, it lists all of your branches (like normal "git branch", then it will also list
the latest commit and the latest commit's date. This enables you to recognize which branch did you work on last time.`,
	Example: fmt.Sprintf(`1) gelp branch
   %s`,
		color.CyanString("gelp branch")),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("`gelp branch` doesn't accept an argument. Please remove it")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		gitLogInfos, err := helpers.BranchList()
		if err != nil {
			panic(err)
		}

		var table [][]string
		for _, info := range gitLogInfos {
			table = append(table, []string{(info.Branch), info.Title, info.Date})
		}

		t := gotabulate.Create(table)
		t.SetAlign("left")
		t.SetHeaders([]string{"Branch", "Last Commit Title", "Last Update Date"})
		stringified := strings.ReplaceAll(t.Render("plain"), "\n", `\n`)
		array := strings.Split(stringified, "\\n")

		var selectOptions []string
		for idx, row := range array {
			if idx%2 == 1 && idx > 1 {
				trimmed := strings.Trim(row, " ")
				if len(trimmed) > 0 {
					selectOptions = append(selectOptions, row)
				}
			}
		}
		fmt.Println(len(selectOptions))
		commitFormat := "{{ . }}"
		templates := &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "\U000027A1 " + commitFormat,
			Inactive: "  " + commitFormat,
			Selected: "\U000027A1 " + commitFormat,
		}

		headerLabel := color.New(color.Bold, color.FgCyan)
		prompt := promptui.Select{
			Label:     headerLabel.Sprintf("    %s", array[1]),
			Items:     selectOptions,
			Templates: templates,
			Size:      5,
		}

		pickedIdx, _, err := prompt.Run()
		if err != nil {
			panic(err)
		}

		pickedCommit := gitLogInfos[pickedIdx]
		err = helpers.CheckoutBranch("branch", pickedCommit.Branch)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
}
