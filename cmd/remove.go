package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/MegeKaplan/gobox/internal/messages"
	"github.com/MegeKaplan/gobox/internal/storage"
	"github.com/MegeKaplan/gobox/internal/utils"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a saved package",
	Long:    "Removes a Go package from the gobox storage. The actual package is not uninstalled from your current project.",
	Example: "  gobox remove",
	Run:     runRemovePackage,
}

func runRemovePackage(cmd *cobra.Command, args []string) {
	packages, err := storage.LoadPackages()
	if err != nil {
		cmd.Println(color.RedString(messages.ErrLoadingPackagesFailed))
		return
	}

	if len(packages) == 0 {
		cmd.Println(color.YellowString(messages.StatusNoPackagesFound))
		return
	}

	utils.SortPackages(&packages, "last_used", true)

	options := []string{}
	for _, pkg := range packages {
		options = append(options, pkg.Name)
	}

	var selected string
	packageSelectPrompt := &survey.Select{
		Message: messages.PromptRemovePackage,
		Options: options,
	}

	var confirm bool
	confirmPrompt := &survey.Confirm{
		Message: messages.PromptConfirmPackageRemoval,
		Default: false,
	}

	if err := survey.AskOne(packageSelectPrompt, &selected); err != nil {
		cmd.Println(color.RedString(messages.ErrPromptFailed))
		return
	}

	if err := survey.AskOne(confirmPrompt, &confirm); err != nil {
		cmd.Println(color.RedString(messages.ErrPromptFailed))
		return
	}

	if !confirm {
		cmd.Println(color.YellowString(messages.StatusPackageRemovalCancelled))
		return
	}

	if err := storage.RemovePackage(selected); err != nil {
		cmd.Println(color.RedString(messages.ErrRemovePackageFailed, selected))
		return
	}

	cmd.Println(color.GreenString(messages.SuccessPackageRemoved, selected))
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
